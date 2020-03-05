package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"kafka/kafka"
	"kafka/tailf"

)

func main() {
	filename := "./conf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed ,err %sv\n", err)
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger conf failed ,err %sv\n", err)
		return
	}

	logs.Debug("initialize successful")
	logs.Debug("load conf successful:%v,", appConfig)


	collectConf,err := initEtcd(appConfig.etcdAddr,appConfig.etcdKey)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}

	logs.Debug("initialize etcd succ")

	err = tailf.InitTail(collectConf,appConfig.chanSize)
	if err != nil {
		logs.Debug("init Tail error,err:%v", err)
		return
	}
	logs.Debug("initialize  all successful")

	err = kafka.InitKafka(appConfig.kafka)
	if err != nil {
		logs.Debug("init kafka error,err:%v", err)
		return
	}
	logs.Debug("initialize  all successful")

	err = serverRun()
	if err != nil {
		logs.Debug("serverRun() error,err:%v", err)
		return
	}
	logs.Info("process exited")
}
