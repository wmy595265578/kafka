package main

import (
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
	"kafka/tailf"
	"strings"
	"time"
)

type EtcdClient struct {
	client *clientv3.Client
	kyes []string
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr ,key string) (collectConf []tailf.CollectConf,err error) {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:            []string{addr},
		AutoSyncInterval:     0,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		LogConfig:            nil,
		Context:              nil,
		PermitWithoutStream:  false,
	})

	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	etcdClient = &EtcdClient{
		client: cli,
		kyes:   nil,
	}

	if strings.HasPrefix(key,"/") == false {
		key = key + "/"
	}



	return
}
