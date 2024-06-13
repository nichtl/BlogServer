package main

import (
	"blogServe/business/global"
	"blogServe/business/initialize"
	initConfig "blogServe/business/initialize/config"
	initMysql "blogServe/business/initialize/db"
	initRedis "blogServe/business/initialize/redis"
	initRoute "blogServe/business/initialize/route"
	"fmt"
)

func main() {

	global.Config = initConfig.InitConfig("", "")
	global.DefaultDb = initMysql.InitMysql()
	global.RedisClient = initRedis.InitRedis()
	route := initRoute.Route()
	port := fmt.Sprintf(":%d", global.Config.Serve.Port)
	s := initialize.InitAndStartServeUseGin(port, route)

	fmt.Println("server started at port ", port)
	err := s.ListenAndServe().Error()

	fmt.Println("listen error ", err)
}
