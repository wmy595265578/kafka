package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)
func convertLogLevel (level string) int {
	//var logLevel int
	switch (level){
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo

	}
	return logs.LevelDebug
}

func initLogger() (err error)  {

	 config :=make(map[string]interface{})
	 config["filename"]= appConfig.logPath
	 config["level"]=  convertLogLevel(appConfig.logLevel)

	configStr,err := json.Marshal(config)

	if err != nil {
		fmt.Println(" initLogger failed,Marshal error failed",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr))

	return
}
