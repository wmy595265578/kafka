package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"kafka/tailf"
)

type Config struct {
	logLevel string
	logPath  string

	etcdAddr string
	etcdKey  string

	chanSize    int
	kafka      string
	collectConf []tailf.CollectConf
}

var (
	appConfig *Config
)

func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	//cc = &CollectConf{}
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid  collenct::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid  collenct::topic")
		return
	}
	appConfig.collectConf = append(appConfig.collectConf, cc)
	return

}
func loadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("New config failed ,err:", err)
		return
	}
	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs/"
	}

	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}

	appConfig.kafka = conf.String("kafka::server_addr")
	if len(appConfig.kafka) == 0 {
		err = fmt.Errorf("kafka invalid addr")
		return
	}
    appConfig.etcdAddr = conf.String("etcd::etcd_addr")
	if len(appConfig.etcdAddr) == 0 {
		err = fmt.Errorf("invalid etcd addr")
		return
	}

	appConfig.etcdKey = conf.String("etcd::etcd_Key")
	if len(appConfig.etcdKey) == 0 {
		err = fmt.Errorf("invalid etcd key")
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed ,err:%v\n", err)
		return
	}
	return
}
