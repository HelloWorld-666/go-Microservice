package handlers

import (
	"context"
	"strconv"

	"HelloWorld/test/base/database"
	"HelloWorld/test/base/proto/student_info"

	log "github.com/cihub/seelog"
)

type StudentInfoHandler struct{}

// 具体业务流程操作：
func (m *StudentInfoHandler) SetStudentInfo(ctx context.Context, req *student_info.SetStudentInfoReq, rsp *student_info.SetStudentInfoRsp) error {
	log.Debug("[grpc][SetStudentInfo][base]")

	var err error
	var dbStu database.StudentInfo
	dbStu.Name = req.Name
	dbStu.Age = req.Age
	dbStu.Sex = req.Sex
	dbStu.MathScore, err = strconv.Atoi(req.MathScore)
	if err != nil {
		log.Error(err)
	}

	insertResult, err := dbStu.Insert(nil)
	if err != nil {
		log.Error("插入数据失败 ! ")
	} else if !insertResult {
		log.Info("要插入的数据不存在 ! ")
	}

	rsp.Name = dbStu.Name
	rsp.Age = dbStu.Age
	rsp.Sex = dbStu.Sex
	rsp.MathScore = strconv.Itoa(dbStu.MathScore)

	return nil
}

func (m *StudentInfoHandler) GetStudentInfo(ctx context.Context, req *student_info.GetStudentInfoReq, rsp *student_info.GetStudentInfoRsp) error {
	var dbStu database.StudentInfo
	rsp.Name = dbStu.Name
	rsp.Age = dbStu.Age
	rsp.Sex = dbStu.Sex
	rsp.MathScore = strconv.Itoa(dbStu.MathScore)

	return nil
}
