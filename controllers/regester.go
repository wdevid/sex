package controllers

import "github.com/astaxie/beego"

type RegesterContreller struct {
	beego.Controller
}

func (c *RegesterContreller) Get()  {
	c.TplName = "regester.html"
}
