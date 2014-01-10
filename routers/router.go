package routers

import (
	"github.com/niuplus/adminkiss/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
