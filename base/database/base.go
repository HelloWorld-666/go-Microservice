package database

import (
	_ "github.com/go-sql-driver/mysql"

	log "github.com/cihub/seelog"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

/*
	1.上面的mysql驱动相关包不能缺少，否则连接数据库时会报错如下：
	  sql: unknown driver "mysql" (forgotten import?)
*/

var OrmEngine *xorm.Engine = nil

func init() { // init首字母小写
	var err error
	//OrmEngine, err = xorm.NewEngine("mysql", "root:x2014.mc@tcp(129.204.12.208:3306)/test?charset=utf8&parseTime=true&loc=Local") // "数据库登录用户:数据库登录密码@tcp(数据库IP:数据库端口)/数据库名称?charset=utf8")
	OrmEngine, err = xorm.NewEngine("mysql", "root:x2014.mc@tcp(129.204.12.208:3306)/test?charset=utf8") // 这里不能再重新分配OrmEngine，例如：OrmEngine, err := .......否则，在database包中的该变量就不能全局共享一个变量值了，会被重新赋初值
	if err != nil {
		log.Critical("数据库初始化失败...   err:%v", err)
	}

	err = OrmEngine.Ping()
	if err != nil {
		log.Info("连接数据库失败! ")
	} else {
		log.Info("连接数据库成功! ")
	}

	OrmEngine.SetColumnMapper(core.SnakeMapper{}) // 数据库表结构名称映射规则： MathScore <-----> math_score（默认），此外还有SameMapper、GonicMapper等映射规则...

	err = OrmEngine.Sync2(
		new(StudentInfo), // 逗号! 同步StudentInfo结构体所在的表结构.
	)
	if err != nil {
		log.Critical(err)
	} else {
		log.Info("数据库表结构同步成功！")
	}
}

/*
	parseTime=true&loc=Local说明会解析时间，时区是机器的local时区。机器之间的时区可能不一致会设置有问题，这导致从相同库的不同实例查询出来的结果可能解析以后就不一样。
	因此推荐将loc统一设置为一个时区，如parseTime=true&loc=America%2FChicago
*/
