package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.etcd.io/etcd/clientv3"
	"kafka/web_admin/models"
	_ "kafka/web_admin/routers"
	"time"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/logadmin")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	models.InitDb(database)
	return
}

func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.1.212:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	logs.Info("connect succ")
	models.InitEtcd(cli)

	return
}
func main() {


	err := initDb()
	if err != nil {
		logs.Warn("initDb failed, err:%v", err)
		return
	}


	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err:%v", err)
		return
	}
	beego.Run()
}
