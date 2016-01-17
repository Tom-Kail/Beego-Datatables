package routers

import (
	"tatu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testAjax", &controllers.MainController{}, "*:TestAjax")
}
