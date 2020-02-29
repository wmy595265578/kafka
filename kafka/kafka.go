package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"github.com/siddontang/go/log"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		log.Error("Init producer failed err :", err)
		return
	}

	//defer client.Close()
	logs.Debug("init kafka successful")

	return
}

func SendToKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed,err:%s, data:%s,topic:%s", err, data, topic)
		return
	}

	logs.Debug("send successful ,pid:%v  offset:%s ,topic:%v\n", pid, offset, topic)
	return
}
