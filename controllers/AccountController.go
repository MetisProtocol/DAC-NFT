package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"metis-v1.0/helpers"
	"metis-v1.0/models"
)

type AccountController struct {
	beego.Controller
}

func (this *AccountController) Register() {
	this.TplName = "account/login_register.html"

	if this.Ctx.Input.IsPost() {
		u := helpers.User{}
		if err := this.ParseForm(&u); err != nil {
			fmt.Println("注册失败")
			this.Data["json"] = map[string]string{"code": "error", "message": "注册失败"}
		} else {
			fmt.Println("注册成功")
			err = models.NewUser().Register(u)
			this.Data["json"] = map[string]string{"code": "ok", "message": "注册成功"}
		}
		this.ServeJSON()
	}
}

func (this *AccountController) Test() {
	this.TplName = "test/test.html"
}
