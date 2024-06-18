package main

import (
	"blogServe/business/global"
	"blogServe/business/initialize"
	initConfig "blogServe/business/initialize/config"
	initMysql "blogServe/business/initialize/db"
	initLog "blogServe/business/initialize/log"
	initRedis "blogServe/business/initialize/redis"
	initRoute "blogServe/business/initialize/route"
	"fmt"
	"go.uber.org/zap"
)

func main() {

	global.Config = initConfig.InitConfig("", "")
	global.DefaultDb = initMysql.InitMysql()
	global.LOG = initLog.InitZap()
	global.RedisClient = initRedis.InitRedis()
	route := initRoute.Route()
	defer func() {
		err := global.RedisClient.Close()
		if err != nil {
			global.LOG.Error("redis close occur error"+err.Error(), zap.Any("err", err))
		}
		if db, err := global.DefaultDb.DB(); err == nil {
			err = db.Close()
			if err != nil {
				global.LOG.Error("database close occur error"+err.Error(), zap.Any("err", err))
			}
		}

	}()
	port := fmt.Sprintf(":%d", global.Config.Serve.Port)
	s := initialize.InitAndStartServeUseGin(port, route)
	fmt.Println("server started at port ", port)
	err := s.ListenAndServe().Error()

	fmt.Println("listen error ", err)
}
