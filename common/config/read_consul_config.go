package config

import (
	"HelloWorld/test/common/logger"
	"os"

	log "github.com/cihub/seelog"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/consul"
)

// type Config struct {
// 	RateLimit struct {
// 		Qps int `json:"qps" yaml:"qps"`
// 		Cap int `json:"cap" yaml:"cap"`
// 	} `json:"rateLimit" yaml:"rateLimit"`
// 	Database struct {
// 		DataSourceName string `json:"dataSourceName" yaml:"dataSourceName"`
// 		Timezone       string `json:"timezone" yaml:"timezone"`
// 		MaxIdleConns   int    `json:"maxIdleConns" yaml:"maxIdleConns"`
// 		MaxOpenConns   int    `json:"maxOpenConns" yaml:"maxOpenConns"`
// 		SyncTable      bool   `json:"syncTable" yaml:"syncTable"`
// 	} `json:"database" yaml:"database"`
// 	GoodaMchtOpenServer struct {
// 		CheckLoginUrl         string `json:"checkLoginUrl" yaml:"checkLoginUrl"`
// 		QueryMchtByUserIdUrl  string `json:"queryMchtByUserIdUrl" yaml:"queryMchtByUserIdUrl"`
// 		GetMchtUrl            string `json:"getMchtUrl" yaml:"getMchtUrl"`
// 		QueryStoreByMchtNoUrl string `json:"queryStoreByMchtNoUrl" yaml:"queryStoreByMchtNoUrl"`
// 		GetMchtUserUrl        string `json:"getMchtUserUrl" yaml:"getMchtUserUrl"`
// 		GetStoreUrl           string `json:"getStoreUrl" yaml:"getStoreUrl"`
// 		QueryDeskSnUrl        string `json:"queryDeskSnUrl" yaml:"queryDeskSnUrl"`
// 		BindDeskSnUrl         string `json:"bindDeskSnUrl" yaml:"bindDeskSnUrl"`
// 	} `json:"goodaMchtOpenServer" yaml:"goodaMchtOpenServer"`
// }

const (
	CONSUL_CONFIG_ADDRESS string = "CONSUL_CONFIG_ADDRESS" // 启动脚本时，用于匹配设置的环境变量->"CONSUL_CONFIG_ADDRESS"
)

// var (
// 	AppConfig *Config = new(Config)
// )

func init() {
	consulConfigAddress := os.Getenv(CONSUL_CONFIG_ADDRESS)
	if len(consulConfigAddress) == 0 {
		panic("请设置读取consul配置的 IP+端口 环境变量 ！")
	}
	log.Info("consul服务地址:", consulConfigAddress)
	consulSource := consul.NewSource(
		consul.WithAddress(consulConfigAddress),
		consul.WithPrefix("/config"),
		consul.StripPrefix(true),
	)
	conf := config.NewConfig()
	err := conf.Load(consulSource)
	if err != nil {
		log.Error(err)
		return
	}

	// loadAppConfig(conf)
	logger.LoadLogCfg(conf)
}

// func loadAppConfig(conf config.Config) {
// 	err := conf.Get("config", common.GOODACDINING_BASE_SERVER).Scan(AppConfig)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}

// 	watcher, err := conf.Watch("config", common.GOODACDINING_BASE_SERVER)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	go func() {
// 		for {
// 			log.Debug("start watcher.Next()")
// 			v, err := watcher.Next()
// 			if err != nil {
// 				panic(err)
// 			}
// 			log.Debug("watcher:", string(v.Bytes()))
// 			if len(v.Bytes()) > 0 {
// 				v.Scan(AppConfig)
// 			}
// 		}
// 	}()
// }
