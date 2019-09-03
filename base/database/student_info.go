package database

import (
	"github.com/go-xorm/xorm"
	// "log"
)

type StudentInfo struct {
	Id        int64  `xorm:"notnull pk autoincr BIGINT(21)" json:"id"`
	Name      string `xorm:"notnull comment('学生姓名') VARCHAR(20)" json:"name"`
	Age       string `xorm:"notnull comment('年龄') VARCHAR(6)" json:"age"`
	Sex       string `xorm:"notnull comment('性别') VARCHAR(20)" json:"sex"`
	MathScore int    `xorm:"notnull comment('数学成绩') INT(11)" json:"mathscore"`
}

/*
	如果结构体拥有TableName() string的成员方法，那么此方法的返回值即是该结构体对应的数据库表名。
*/

func (m *StudentInfo) TableName() string {
	return "tbl_student" // 表名
}

func (m *StudentInfo) Insert(session *xorm.Session) (bool, error) { // 此处必须为指针 ？返回insert是否成功的状态值
	if session == nil {
		// log.Printf("OrmEngine: %v", OrmEngine)
		session = OrmEngine.NewSession()
		defer session.Close()
	}

	affected, err := session.Insert(m)
	if err != nil {
		return false, err
	} else if affected == 0 { // 新插入0条数据
		return false, nil
	}

	return true, nil
}

func (m *StudentInfo) Update() {

}

func (m *StudentInfo) Delete() {

}

func (m *StudentInfo) FindPage() {

}
