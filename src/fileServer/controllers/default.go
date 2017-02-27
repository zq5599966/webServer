package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type ResController struct {
	beego.Controller
}

func (c *ResController) Get() {

	fmt.Println("url path======", c.Ctx.Request.URL.Path)
	fmt.Println("url Fragment======", c.Ctx.Request.URL.Fragment)
	fmt.Println("url Host======", c.Ctx.Request.URL.Host)
	fmt.Println("url Opaque======", c.Ctx.Request.URL.Opaque)
	fmt.Println("url RawPath======", c.Ctx.Request.URL.RawPath)
	fmt.Println("url RawQuery======", c.Ctx.Request.URL.RawQuery)
	fmt.Println("url Scheme======", c.Ctx.Request.URL.Scheme)



	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}