/*
 * @FileName:   findSrv_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/11 上午9:51
 * @Description:
 */

package consul

import (
	"context"
	"msFileStore/proto"
	utils "msFileStore/utils/port"
	"msFileStore/webServer/global"
	"msFileStore/webServer/initialize"
	"msFileStore/webServer/model"

	"os"
	"testing"
)

const configPath = "/media/xjj/common/work/project/goProject/msFileStore/config-debug.yaml"

func TestInitSrvconn(t *testing.T) {

	//服务端配置初始化--------------------------------------------
	initialize.InitConfig(configPath)
	initialize.GormMysql()
	if global.DB != nil {
		model.RegisterTable()
		initialize.MysqlTables(global.DB)
		//db, _ := global.GVA_DB.DB()
		//defer db.Close()
	}
	//如果是本地开发环境端口号固定，线上环境启动获取端口号----------------
	debug := os.Getenv("FILESTORE_DEBUG")
	if debug == "" {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.GrpcWeb.Port = port
		}
	}
	//测试，连接conful查找GRPC服务端，获取其客户端对象-----------------------
	client, err := InitSrvConn("192.168.32.90", 8500, "fileStore_srv")
	if err != nil {
		t.Error(err)
	}
	resdata, err := client.PingPong(context.Background(), &proto.Ping{Stroke: 1111})
	if err != nil {
		t.Error(err)
	}
	t.Log(resdata)

}
