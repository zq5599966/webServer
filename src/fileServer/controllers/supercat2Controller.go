package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	_ "fileServer/models"

	"fileServer/models"
	"encoding/base64"
	"io/ioutil"
	"log"
	"crypto/md5"
)


type SuperCat2ActionController struct {
	beego.Controller
}

type SuperCatActionController struct {
	beego.Controller
}

type SupercatActionUpload struct {
	beego.Controller
}

func returnErr(c *SuperCat2ActionController) {
	c.Ctx.WriteString("error")
}

func decodeDataStr(str string) string{
	decodeByte, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("decode base64 error=====%v\n", err)
		return "error"
	}

	fmt.Println("decode str====", string(decodeByte))

	return string(decodeByte)
}

const (
	RESPONSE_UNKNOWERROR = iota
	RESPONSE_SUCCESS
)


func (c *SuperCat2ActionController) Get() {

	actionId := c.GetString("id", "unknow")
	data := c.GetString("data", "null")

	fmt.Println("actionid=====", actionId)
	fmt.Printf("data====%v\n", data)

	if actionId == "request" {
		data = decodeDataStr(data)
		if data == "error" {
			returnErr(c)
			return
		}
		c.Ctx.WriteString(models.NewFlagEvent(data))

	} else if actionId == "response" {
		data = decodeDataStr(data)
		if data == "error" {
			returnErr(c)
			return
		}
		c.Ctx.WriteString(models.GetFlagsEvent(data))

	} else if actionId == "share" {
		//http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, "phantomcat2://veewo.com?data=123qww", http.StatusFound)
		fmt.Println("fuck share=============")
		c.Data["scheme"] = data
		c.TplName = "scheme.html"

	} else {
		//unknow
		fmt.Println("send error")
		returnErr(c)
		return
	}
}

func (c *SuperCatActionController) Get() {

	actionId := c.GetString("id", "unknow")
	data := c.GetString("data", "null")

	fmt.Println("actionid=====", actionId)
	fmt.Printf("data====%v\n", data)

	if actionId == "multiPlay" {
		c.Data["scheme"] = data
		c.TplName = "supercat_scheme.html"
		//c.Ctx.WriteString("success")
		return;
	}

	c.Ctx.WriteString("success")
}


func (c *SupercatActionUpload) Post(){
	file, h, _ := c.GetFile("file")
	path := "static/res/" + h.Filename
	println("path====" + path)
	println("file name====" + h.Filename)
	file.Close()

	c.SaveToFile("file", path)

	if createManifest(path) {
		c.Ctx.WriteString("success")
	}else{
		c.Ctx.WriteString("error")
	}




}

func createManifest(path string) bool {
	data, err := ioutil.ReadFile(path)
	if err != nil{
		println("create manifest error")
		log.Fatal(err)
		return false
	}

	fmt.Printf("Md5: %x\n\n", md5.Sum(data))


	return true
}