package main

import (
	"github.com/astaxie/beego/logs"
	"kafka/kafka"
	"kafka/tailf"
	"time"
)

func serverRun() (err error) {

	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
   		 logs.Error("send msg to kafka failed,err:%s\n",err)
   		 time.Sleep(time.Second)
   		 continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error){

	//fmt.Printf("msg:%s,topic:%s\n",msg.Msg,msg.Topic)

	kafka.SendToKafka(msg.Msg,msg.Topic)

	return
}
