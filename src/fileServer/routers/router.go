package routers

import (
	"fileServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
    //beego.Router("/res/*", &controllers.ResController{})
	beego.Router("/supercat2/action/", &controllers.SuperCat2ActionController{})

	beego.Router("supercat/action/", &controllers.SuperCatActionController{})

	beego.Router("/upload", &controllers.SupercatActionUpload{})
}
