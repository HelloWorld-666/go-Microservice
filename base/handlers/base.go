package handlers

import (
	"HelloWorld/test/base/proto/student_info"

	"github.com/micro/go-micro/server"
)

// handler注册具体方法：
func RegisterService(server server.Server) {
	// student_info.RegisterStudentServiceHandler(server, new(StudentInfoHandler)) // 两种方式选其一
	student_info.RegisterStudentServiceHandler(server, &StudentInfoHandler{})
}
