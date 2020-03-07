package main

import (
	"context"
	//"encoding/json"

	//"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"kafka/tailf"
	"time"
)

const (
	EtcdKey = "./my.log/192.168.1.212"
)

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.1.212:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

		var logConfArr []tailf.CollectConf
		logConfArr = append(
			logConfArr,
			tailf.CollectConf{
				LogPath: "./my.log",
				Topic:   "mykafka",
			},
		)
		//logConfArr = append(
		//	logConfArr,
		//	tailf.CollectConf{
		//		LogPath: "./my2.log",
		//		Topic:   "mykafka",
		//	},
		//)

		//data, err := json.Marshal(logConfArr)
		//if err != nil {
		//	fmt.Println("json failed, ", err)
		//	return
		//}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		//cli.Delete(ctx, EtcdKey)
		//return
		//_, err = cli.Put(ctx, EtcdKey, string(data))
		cancel()
		//if err != nil {
		//	fmt.Println("put failed, err:", err)
		//	return
		//}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	//cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}
