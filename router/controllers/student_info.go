package controllers

import (
	"HelloWorld/test/base/proto/student_info"
	"HelloWorld/test/router/models"
	"context"

	log "github.com/cihub/seelog"
	"github.com/kataras/iris"
)

func SetStudentInfo(ctx iris.Context) {
	log.Debug("[grpc][SetStudentInfo][router]")

	var stu models.StudentInfo
	err := ctx.ReadJSON(&stu) // 这里一定要用指针么 ？
	if err != nil {
		log.Error("err: ReadJson error ! ")
	}

	var req student_info.SetStudentInfoReq
	req.Name = stu.Name
	req.Age = stu.Age
	// if err != nil {	// 到时候可以定义为stu.Age可以定义为int类型.*****************
	// 	log.Error(err)
	// }
	req.Sex = stu.Sex
	req.MathScore = stu.MathScore
	log.Info("[req]:", req)

	rsp, err := studentService.SetStudentInfo(context.Background(), &req) // context.Background()的作用 ？调用base的grpc服务
	if err != nil {
		log.Error("err:", err)
	}
	log.Info("[rsp]:", rsp)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(rsp)
	//studentService.SetStudentInfo(ctx, &req) // 错误演示！不是ctx，而是context.Background
}

func GetStudentInfo(ctx iris.Context) {
	var req student_info.GetStudentInfoReq

	rsp, err := studentService.GetStudentInfo(context.Background(), &req)
	if err != nil {
		log.Error("err:", err)
	}
	// rsp.Name = "test name"

	log.Info("[rsp]:", rsp)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(rsp)
}
