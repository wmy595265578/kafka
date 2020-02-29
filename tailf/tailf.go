package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
}

var (
	tailObjMgr *TailObjMgr
)

func InitTail(conf []CollectConf, chanSize int) (err error) {

	if len(conf) == 0 {
		err = fmt.Errorf("invalid config for log collect ,conf:%v", conf)
		return

	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}
	for _, v := range conf {
		obj := &TailObj{
			conf: v,
		}

		tails, tailerr := tail.TailFile(v.LogPath, tail.Config{
			Location: &tail.SeekInfo{
				Offset: 0,
				Whence: 2,
			},
			ReOpen:      true,
			MustExist:   false,
			Poll:        true,
			Pipe:        false,
			RateLimiter: nil,
			Follow:      true,
			MaxLineSize: 0,
			Logger:      nil,
		})
         err = tailerr
		if err != nil {
			fmt.Println("tail file err", err)
			return
		}

		obj.tail = tails
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

		go readFromTail(obj)
	}

	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		lines, ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen,filename:%s\n", tailObj.tail.Filename)
			return
		}

		textMsg := &TextMsg{
			Msg:   lines.Text,
			Topic: tailObj.conf.Topic,
		}

		tailObjMgr.msgChan <- textMsg
	}

}
