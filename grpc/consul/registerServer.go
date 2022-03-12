/*
 * @FileName:   register.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/8 下午1:36
 * @Description:
 */

package consul

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
)

/**
@Function:ConsulRegister
@Description: 将grpc服务端注册到consul
@Param consulAddr string:
@Return:
@author:JunjieXu
*/
func NewConsulRegisterSrv(consulAddr string, consulPort int, serverHost string, serverPort int, serverName string, serverTags []string, serverId string) (consulClient *api.Client, err error) {

	//服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulAddr, consulPort)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = serverName
	registration.ID = serverId
	registration.Port = serverPort
	registration.Tags = serverTags
	registration.Address = serverHost

	//生成对应的检查对象，填写自己的地址
	check := &api.AgentServiceCheck{
		GRPC: fmt.Sprintf("%s:%d", serverHost, serverPort),
		//HTTP:                           fmt.Sprintf("%v:%v", serverHost, serverPort), //todo fmt.Sprintf("%v:%v/health", address, port),
		Interval:                       "5s",  //间隔5
		Timeout:                        "5s",  //5秒超时
		DeregisterCriticalServiceAfter: "15s", //严重超时取消注册服务时间
	}
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("consul register err: %v", err))
	}

	return client, nil
}
