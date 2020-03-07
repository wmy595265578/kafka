package routers

import (
	"github.com/astaxie/beego"
	"kafka/web_admin/controllers/AppController"
)

func init() {
	beego.Router("/", &AppController.AppController{}, "*:AppList")
	beego.Router("/index", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/list", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/apply", &AppController.AppController{}, "*:AppApply")
	beego.Router("/app/create", &AppController.AppController{}, "*:AppCreate")

	beego.Router("/log/list", &AppController.LogController{}, "*:LogList")
	beego.Router("/log/apply", &AppController.LogController{}, "*:LogApply")
	beego.Router("/log/create", &AppController.LogController{}, "*:LogCreate")
}
