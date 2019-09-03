package web

import (
	"HelloWorld/test/router/controllers"

	"github.com/kataras/iris"
)

// 路由相关文档：https://studyiris.com/example/iris.html

func InitUrlSuffix(app *iris.Application) {
	v1 := app.Party("/v1")
	{
		student := v1.Party("/student")
		{
			student.Get("/student", func(ctx iris.Context) { ctx.WriteString("Welcome to student system ~") })
			student.Post("/set_student_info", controllers.SetStudentInfo)
			student.Any("/get_student_info", controllers.GetStudentInfo) // Any方法支持Post或Get请求.
		}
	}

}
