package main

import (
	_ "HelloWorld/test/common/config"
	"HelloWorld/test/common/logger"
	"HelloWorld/test/router/controllers"
	_ "HelloWorld/test/router/web"

	log "github.com/cihub/seelog"
	micro "github.com/micro/go-micro"

	"fmt"
)

/*
	1.上面不要漏掉micro包
	2.上面的下划线 "_" 自动执行web包中的init()函数
*/
func main() {
	fmt.Println("我是main包的输出函数")
	logger.InitLogger()

	service := micro.NewService(
		micro.Name("student-router-server"), // 服务端定义服务名
		micro.Version("1.0.0"),
	)
	service.Init()

	controllers.CreateClientObj(service.Client()) // 不要忘记调用该函数,否则grpc调用studentService.SetStudentInfo时,studentService指针指向的内存出现错误.

	err := service.Run()
	if err != nil {
		log.Criticalf("router server.Run() failed ! err: %v", err)
	}
}
