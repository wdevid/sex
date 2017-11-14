package routers

import (
	"sex/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginContreller{})
    beego.Router("/regester", &controllers.RegesterContreller{})
}
