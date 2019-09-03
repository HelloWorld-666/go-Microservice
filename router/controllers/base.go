package controllers

import (
	"HelloWorld/test/base/proto/student_info"

	"github.com/micro/go-micro/client"
)

// micro框架中的client，具体用法未知 ？

const STUDENT_SERVICE = "student-base-server" // 定义客户端调用的服务名

var (
	studentService student_info.StudentService
)

// controller创建客户端实例，调用具体方法：
func CreateClientObj(client client.Client) {
	studentService = student_info.NewStudentService(STUDENT_SERVICE, client)
}
