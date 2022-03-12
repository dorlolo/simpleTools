/*
 * @FileName:   findServer.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/10 下午9:01
 * @Description:
 */

package consul

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"msFileStore/proto"
	"msFileStore/webServer/global"
)

//
//  FilterFileStoreServices
//  @Description: consul原生的基于HTTP协议的服务发现方法。支持自定义负载均衡。
//				  注意这个函数中服务测试的方法需要每次都需要更改。
//  @param consulHost : consul 的host地址
//  @param consulPort : consul 的 端口
//  @param serverHost : 全局变量中服务端host的全局变量指针地址
//  @param serverPort : 全局变量中服务端端口的全局变量指针地址
//  @return proto.FileStoreGreeterClient
//
func FilterFileStoreServices(consulHost string, consulPort int, srvName string) (proto.FileStoreGreeterClient, error) {
	cfg := api.DefaultConfig()
	//consulInfo := global.ServerConfig.Consul
	cfg.Address = fmt.Sprintf("%s:%d", consulHost, consulPort)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//根据服务名或id查询grpc服务端地址
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service=="%v"`, srvName))
	if err != nil {
		panic(err)
	}

	for _, value := range data {
		// 跳过负载均衡策略，拿到直接退出
		//把连接的gprc服务地址和端口复制到全局变量
		global.ServerConfig.GrpcSrv.Host = value.Address
		global.ServerConfig.GrpcSrv.Port = value.Port
		break
	}

	address := fmt.Sprintf("%v:%v", global.ServerConfig.GrpcSrv.Host, global.ServerConfig.GrpcSrv.Port)
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		panic(err)
	}
	//defer conn.Close()
	//获取grpc客户端对象
	FileStoreClient := proto.NewFileStoreGreeterClient(conn)
	//测试通信是否正常
	ping := proto.Ping{Stroke: 1111}
	pong, err := FileStoreClient.PingPong(context.Background(), &ping)
	if err != nil {
		return nil, err
	} else if pong.GetStroke() != 1111 {
		return nil, errors.New(fmt.Sprintf("ping data is %d,but pong data is %d", ping.GetStroke(), pong.GetStroke()))
	}
	return FileStoreClient, nil
}

//
//  InitSrvConn
//  @Description: 基于GRPC协议的服务发现方法。负载均衡规则是由consul服务决定的，
// 				  ！！！本方法会导致每次向GRPC服务端的请求都要经过consul服务转发。！！
//  @param consulHost : consul 的host地址
//  @param consulPort : consul 的 端口
//  @param serverHost : 全局变量中服务端host的全局变量指针地址
//  @param serverPort : 全局变量中服务端端口的全局变量指针地址
//  @return proto.FileStoreGreeterClient
//
func InitSrvConn(consulHost string, consulPort int, srvName string) (proto.FileStoreGreeterClient, error) {
	address := fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulHost, consulPort, srvName)
	//var bulider = NewBuilder()
	userConn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//grpc.WithResolvers(bulider),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		fmt.Println("[InitSrvConn] 连接 【用户服务失败】", err)
	}

	FileStoreClient := proto.NewFileStoreGreeterClient(userConn)
	if FileStoreClient == nil {
		return FileStoreClient, errors.New("未找到匹配的服务端")
	}
	//测试通信是否正常
	ping := proto.Ping{Stroke: 1111}
	pong, err := FileStoreClient.PingPong(context.Background(), &ping)
	if err != nil {
		return nil, err
	} else if pong.GetStroke() != 1111 {
		return nil, errors.New(fmt.Sprintf("ping data is %d,but pong data is %d", ping.GetStroke(), pong.GetStroke()))
	}
	return FileStoreClient, nil
}
