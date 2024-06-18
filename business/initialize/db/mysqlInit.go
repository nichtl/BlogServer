package db

import (
	"blogServe/business/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMysql() *gorm.DB {
	m := global.Config.Mysql
	if m.Host == "" || m.Database == "" || m.Port == 0 || m.Username == "" || m.Password == "" {
		panic(fmt.Errorf("mysql initialize error  invalid config"))
	}
	dsn, err := m.GetDSN()
	if err != nil {
		return nil
	}
	mysqlConnectConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置}
	}
	gormConf := &gorm.Config{}
	gormConf.Logger = global.Config.GormConfig.GetLogLevel()

	db, err := gorm.Open(mysql.New(mysqlConnectConfig), gormConf)
	if err != nil {
		panic(err)
	}
	db.InstanceSet("gorm:table_options", "Engine=Innodb")
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(m.MinIdelConn)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTimeMillis) * time.Millisecond)
	return db
}
