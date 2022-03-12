/*
 * @FileName:   initDB.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/7 下午2:09
 * @Description:
 */

package simpleDB

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitMysql() {

}

//初始化mysql
func InitGormMysql(m Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	//dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil singulartable
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//自动迁移数据表
func AutoMigrate(db *gorm.DB, tableList []interface{}) {
	err := db.AutoMigrate(tableList...)
	if err != nil {
		fmt.Println("register table failed:", err)
		os.Exit(0)
	}
	fmt.Println("register table success")
}
