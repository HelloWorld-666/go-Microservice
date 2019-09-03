package main

import (
	_ "HelloWorld/test/base/database"
	"HelloWorld/test/base/handlers"
	_ "HelloWorld/test/common/config"
	"HelloWorld/test/common/logger"

	log "github.com/cihub/seelog"
	micro "github.com/micro/go-micro"
)

/*
	1.上面不要漏掉micro
	2. _下划线自动调用上面的初始化init()函数.否则在insert操作中，OrmEngine.NewSession()会报错；原因：OrmEngine未成功连接mysql.
*/

func main() {
	logger.InitLogger()

	service := micro.NewService(
		micro.Name("student-base-server"), // 服务端定义服务名
		micro.Version("1.0.0"),
	)
	service.Init()

	handlers.RegisterService(service.Server())

	err := service.Run()
	if err != nil {
		log.Criticalf("base server.Run() failed ! err: %v", err)
	}
}
