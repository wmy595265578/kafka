package models

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type CollectConf struct {
	LogPath string `json:"logpath"`
	Topic   string `json:"topic"`
}

type LogInfo struct {
	LogId      int    `db:"log_id"`
	AppName    string `db:"app_name"`
	AppID      int    `db:"app_id"`
	LogPath    string `db:"log_path"`
	CreateTime string `db:"create_time"`
	Topic      string `db:"topic"`
	status     int    `db:"topic"`
}

var (
	etcdClient *clientv3.Client
)

func InitEtcd(client *clientv3.Client) {
	etcdClient = client
}
func GetAllLogInfo() (logList []LogInfo, err error) {
	err = Db.Select(&logList,
		"select log.log_id, app.app_name, log.create_time, log.log_path ,log.topic from tbl_app_info as app,tbl_log_info as log where app.app_id=log.app_id")
	if err != nil {
		logs.Warn("Get All App Info failed, err:%v", err)
		return
	}
	return
}

func CreateLog(info *LogInfo) (err error) {
	conn, err := Db.Begin()
	if err != nil {
		logs.Warn("CreateLog failed, Db.Begin error:%v", err)
		return
	}
	defer func() {
		if err != nil {
			conn.Rollback()
			return
		}
		conn.Commit()

	}()
	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", info.AppName)
	if err != nil || len(appId) == 0 {
		logs.Warn("select app_id failed ,err:%v", err)
		return
	}
	info.AppID = appId[0]
	r, err := conn.Exec("insert into tbl_log_info(app_name, app_id, log_path,topic)values(?,?,?,?)",
		info.AppName, info.AppID, info.LogPath, info.Topic)
	if err != nil {
		logs.Warn("CreateLog failed, Db.Exec error:%v", err)
		return
	}

	_, err = r.LastInsertId()
	if err != nil {
		logs.Warn("CreateLog failed, Db.LastInsertId error:%v", err)
		return
	}
	return

}

const (
	EtcdKey = "./my.log/192.168.1.212"
)

func SetLogConfToEtcd(etcdKey string, info *LogInfo) (err error) {

	var logConfArr []CollectConf
	logConfArr = append(
		logConfArr,
		CollectConf{
			LogPath: info.LogPath,
			Topic:   info.Topic,
		},
	)
	//logConfArr = append(
	//	logConfArr,
	//	tailf.CollectConf{
	//		LogPath: "./my2.log",
	//		Topic:   "mykafka",
	//	},
	//)

	data, err := json.Marshal(logConfArr)
	if err != nil {
		logs.Error("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//cli.Delete(ctx, EtcdKey)
	//return
	_, err = etcdClient.Put(ctx, EtcdKey, string(data))
	logs.Info(string(data))
	cancel()
	return
	//if err != nil {
	//	fmt.Println("put failed, err:", err)
	//	return
	//}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, EtcdKey)
	//cancel()
	//if err != nil {
	//	fmt.Println("get failed, err:", err)
	//	return
	//}
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	//}
}
