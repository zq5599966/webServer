package routers

import (
	"fileServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/res/*", &controllers.ResController{})
}
