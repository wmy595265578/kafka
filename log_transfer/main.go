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
	 /*
	 err = initKafka()
	 if err !=nil {
	 	logs.Error("init kafka  failed err:%s",err)
		 return
	 }

	 err = initES()
	 if err !=nil {
		 logs.Error("init ES failed err:%s",err)
		 return
	 }

	 err = run()
	if err !=nil {
		logs.Error("init Run failed err:%s",err)
		return
	}
	logs.Warn("warning,log_transfer is exited")
*/
}

