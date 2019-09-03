package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func init() {
	app := iris.New()
	app.Use(recover.New()) // 从任何与http相关的panic中恢复
	app.Use(logger.New())  // 将请求记录到终端
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello, welocome~")
	})

	InitUrlSuffix(app) // 路由子域名跳转

	go app.Run(iris.Addr(":8080")) // 启动服务，需要在程序最后执行	// 这里要在线程里跑 ？
}
