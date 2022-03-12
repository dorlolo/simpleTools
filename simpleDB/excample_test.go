/*
 * @FileName:   use.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/7 下午2:18
 * @Description:
 */

package simpleDB

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

//全局参数
var (
	DB        *gorm.DB
	TableList []interface{}
)

//添加数据库连接参数
var dbConfig = Mysql{
	Path:     "127.0.0.1:3306", // 服务器地址:端口
	Dbname:   "msfilestore",    // 数据库名
	Username: "xjj",            // 数据库用户名
	Password: "666666",         // 数据库密码
}

//建立一张测试表
type TestTable struct {
	gorm.Model
	Name string `json:"name" gorm:"comment;名字"`
}

func (*TestTable) TableName() string {
	return "test_Table"
}

//初始化数据库方法
func InitDB() {
	TableList = append(TableList, TestTable{})
	DB = InitGormMysql(dbConfig)
	if DB != nil {
		AutoMigrate(DB, TableList)
	}
}

func TestInitDB(t *testing.T) {
	InitDB()
	fmt.Println("初始化结束")
}
