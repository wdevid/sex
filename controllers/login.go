package controllers

import "github.com/astaxie/beego"

type LoginContreller struct {
	beego.Controller
}

func (c *LoginContreller) Get()  {
	c.TplName = "login.html"
}
