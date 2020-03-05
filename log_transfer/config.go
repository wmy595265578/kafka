package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	kafkaAddr string
	ESAddr    string
	LogPath   string
	LogLevel  string
}

var (
	logConfig *LogConfig
)

func initConfig(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed ,err:", err)
		return
	}

	logConfig = &LogConfig{}

	logConfig.LogLevel = conf.String("log::LogLevel")
	if len(logConfig.LogLevel) == 0 {
		logConfig.LogLevel = "debug"
	}

	logConfig.LogPath = conf.String("log::LogPath")
	if len(logConfig.LogPath) == 0 {
		logConfig.LogPath = "./log"
	}

	logConfig.kafkaAddr = conf.String("kafka::kafkaAddr")
	if len(logConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafkaAddr err")
		return
	}

	logConfig.ESAddr = conf.String("es::ESAddr")
	if len(logConfig.ESAddr) == 0 {
		err = fmt.Errorf("invalid ESAddr err")
		return
	}
	return
}
