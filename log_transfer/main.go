package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main()  {
     err := initConfig("ini","../conf/logransfer.conf")
     if err != nil {
     	panic(err)
		 return
	 }
	 fmt.Println(logConfig)

	 err = initLogger(logConfig.LogPath,logConfig.LogLevel)
	 if err !=nil {
	 	panic(err)
		 return
	 }
	 logs.Info("init log success")

	 err = initKafka(logConfig.kafkaAddr,logConfig.Topic)
	 if err !=nil {
	 	logs.Error("init kafka  failed err:%s",err)
		 return
	 }
	 logs.Debug("init kafka successful")

	 err = initES(logConfig.ESAddr)
	 if err !=nil {
		 logs.Error("init ES failed err:%s",err)
		 return
	 }
	logs.Debug("init ES successful")

         kafkaClient.wg.Add(1)
		 err = run()
		if err !=nil {
			logs.Error("init Run failed err:%s",err)
			return
		}
		kafkaClient.wg.Done()
		logs.Debug("func run is running")
		 kafkaClient.wg.Wait()
		logs.Warn("warning,log_transfer is exited")

}

