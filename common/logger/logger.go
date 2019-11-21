package logger

import (
	"goodacdining-common"

	log "github.com/cihub/seelog"
	"github.com/micro/go-config"

	"fmt"
)

type Config struct {
	LogConfig struct {
		IsDevelopmentEnv       string `json:"isDevelopmentEnv" yaml:"isDevelopmentEnv"`
		Format                 string `json:"format" yaml:"format"`
		InfoFileName           string `json:"infoFileName" yaml:"infoFileName"`
		DebugFileName          string `json:"debugFileName" yaml:"debugFileName"`
		ErrorFileName          string `json:"errorFileName" yaml:"errorFileName"`
		CriticalFileName       string `json:"criticalFileName" yaml:"criticalFileName"`
		SaveOldInfoLogPath     string `json:"saveOldInfoLogPath" yaml:"saveOldInfoLogPath"`
		SaveOldDebugLogPath    string `json:"saveOldDebugLogPath" yaml:"saveOldDebugLogPath"`
		SaveOldErrorLogPath    string `json:"saveOldErrorLogPath" yaml:"saveOldErrorLogPath"`
		SaveOldCriticalLogPath string `json:"saveOldCriticalLogPath" yaml:"saveOldCriticalLogPath"`
		TimePattern            string `json:"timePattern" yaml:"timePattern"`
		MaxRolls               int    `json:"maxRolls" yaml:"maxRolls"`
	} `json:"log" yaml:"log"`
}

var (
	logConfig *Config = new(Config)
)

func InitLogger() {
	logConf := logConfig.LogConfig // 从consul中读取配置文件
	log.Debug("读取consul日志模块配置:", logConf)

	consoleWriter, err := log.NewConsoleWriter() // 标准输出 参数需传入log.NewSplitDisPatcher()的第二个参数中
	if err != nil {
		log.Error("日志模块初始化失败:log.NewConsoleWriter().")
	}
	infoFileWriterTime, err := log.NewRollingFileWriterTime(logConf.InfoFileName, 1, logConf.SaveOldInfoLogPath, logConf.MaxRolls, logConf.TimePattern, 0, true, true) // 滚动日志文件
	if err != nil {
		log.Error("日志模块初始化失败:infoFileWriterTime error.")
	}
	errorFileWriterTime, err := log.NewRollingFileWriterTime(logConf.ErrorFileName, 1, logConf.SaveOldErrorLogPath, logConf.MaxRolls, logConf.TimePattern, 0, true, true) // 滚动日志文件
	if err != nil {
		log.Error("日志模块初始化失败:errorFileWriterTime error.")
	}
	criticalFileWriterTime, err := log.NewRollingFileWriterTime(logConf.CriticalFileName, 1, logConf.SaveOldCriticalLogPath, logConf.MaxRolls, logConf.TimePattern, 0, true, true) // 滚动日志文件
	if err != nil {
		log.Error("日志模块初始化失败:criticalFileWriterTime error.")
	}

	format, err := log.NewFormatter(logConf.Format) // 等价于配置中的<format>
	if err != nil {
		log.Error("日志模块初始化失败:log.NewFormatter().")
	}

	infoFilter, err := log.NewFilterDispatcher(format, []interface{}{infoFileWriterTime}, log.InfoLvl)
	if err != nil {
		log.Error("日志模块初始化失败:infoFilter error.")
	}
	errorFilter, err := log.NewFilterDispatcher(format, []interface{}{errorFileWriterTime}, log.ErrorLvl)
	if err != nil {
		log.Error("日志模块初始化失败:errorFilter error.")
	}
	criticalFilter, err := log.NewFilterDispatcher(format, []interface{}{criticalFileWriterTime}, log.CriticalLvl)
	if err != nil {
		log.Error("日志模块初始化失败:criticalFilter error.")
	}

	//selectOutputLevel, _ := log.NewListConstraints([]log.LogLevel{log.DebugLvl})    		  // 可选择输出任意日志级别
	selectOutputLevel, err := log.NewMinMaxConstraints(log.TraceLvl, log.CriticalLvl) // 输出 Trace ~ Critical 级别之间的日志
	if err != nil {
		log.Error("日志模块初始化失败:log.NewMinMaxConstraints().")
	}

	// 测试、生产环境下不打印debug日志 		 	0:测试环境、生产环境   1:开发环境
	if logConf.IsDevelopmentEnv == "0" {
		root, err := log.NewSplitDispatcher(format, []interface{}{infoFilter, errorFilter, criticalFilter, consoleWriter}) // 等价于配置中的<output>
		if err != nil {
			log.Error("日志模块初始化失败:log.NewSplitDispatcher().")
		}
		logger := log.NewAsyncLoopLogger(log.NewLoggerConfig(selectOutputLevel, nil, root)) // 生成一个完整的seelog配置 nil
		log.ReplaceLogger(logger)
	} else {
		debugFileWriterTime, err := log.NewRollingFileWriterTime(logConf.DebugFileName, 1, logConf.SaveOldDebugLogPath, logConf.MaxRolls, logConf.TimePattern, 0, true, true) // 滚动日志文件
		if err != nil {
			log.Error("日志模块初始化失败:debugFileWriterTime error.")
		}
		debugFilter, err := log.NewFilterDispatcher(format, []interface{}{debugFileWriterTime}, log.DebugLvl)
		if err != nil {
			log.Error("日志模块初始化失败:debugFilter error.")
		}

		root, err := log.NewSplitDispatcher(format, []interface{}{infoFilter, debugFilter, errorFilter, criticalFilter, consoleWriter}) // 等价于配置中的<output>
		if err != nil {
			log.Error("日志模块初始化失败:log.NewSplitDispatcher().")
		}
		logger := log.NewAsyncLoopLogger(log.NewLoggerConfig(selectOutputLevel, nil, root)) // 生成一个完整的seelog配置 nil
		log.ReplaceLogger(logger)
	}
}

// xml文件配置
// func InitLogger() {
// 	mainlog, err := log.LoggerFromConfigAsFile("../common/config.xml")
// 	if err != nil {
// 		log.Critical("err parsing config log file", err)
// 		return
// 	}
// 	//mainlog.Flush()
// 	log.ReplaceLogger(mainlog) //替换mainlog的配置文件
// }

// 加载位于consul上的配置文件到结构体对象中
func LoadLogCfg(conf config.Config) {
	err := conf.Get("config", common.GOODACDINING_LOG_CONFIGURATION).Scan(logConfig) // consul->key_value目录：config/log_configuration
	if err != nil {
		log.Debug(err)
		return
	}

	// // 动态读取consul上的配置文件，暂时有问题，无法热加载
	// watcher, err := conf.Watch("config", "log_configuration")
	// if err != nil {
	// 	log.Error(err)
	// 	return
	// }
	// go func() {
	// 	for {
	// 		log.Debug("start watcher.Next()")
	// 		v, err := watcher.Next()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		log.Debug("watcher:", string(v.Bytes()))
	// 		if len(v.Bytes()) > 0 {
	// 			v.Scan(AppConfig)
	// 		}
	// 	}
	// }()
}

func init() {
	fmt.Println("我是logger包中的init()函数！！！")
}
