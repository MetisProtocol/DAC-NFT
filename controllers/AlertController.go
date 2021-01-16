package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"metisV1.0/helpers"
	"metisV1.0/models"
)

type AlertController struct {
	beego.Controller
}

func (c *AlertController) Wallet() {
	c.TplName = "alert/wallet.html"
}
func (c *AlertController) EmailRegister() {
	c.TplName = "alert/email-register.html"
	if c.Ctx.Input.IsPost() {
		email := helpers.Email{}
		if err := c.ParseForm(&email); err != nil {
			fmt.Println("发送失败")
			c.Data["json"] = JSONS{"error", "发送失败"}
		} else {
			code := helpers.GetRandomString(6)
			c.SetSession("email", email.ToEmail)
			r := helpers.Get()
			defer r.Close()
			_, erre := r.Do("SET", email.ToEmail, code)
			if erre != nil {
				c.Data["json"] = JSONS{"error", "发送失败"}
			} else {
				if num := models.NewDac().GetDacByEmail(email.ToEmail); num != 1 {
					PublicKey := models.NewEth().GetPublicKey(email.ToEmail)
					if PublicKey == "nil" {
						c.Data["json"] = JSONS{"error", "发送失败"}
					} else {
						code = code + " " + PublicKey
					}
				}
				err := SendMail(email.ToEmail, "你好", code)
				if err != nil {
					fmt.Println("发送失败")
					c.Data["json"] = JSONS{"error", "发送失败"}
				} else {
					c.Data["json"] = JSONS{"ok", "发送成功"}
				}
				fmt.Println("发送成功")
			}
		}
		c.ServeJSON()
	}
}
func (c *AlertController) EmailLogin() {
	c.TplName = "alert/email-login.html"
}
func (c *AlertController) EmailCode() {
	email := c.GetSession("email")
	c.Data["email"] = email
	c.TplName = "alert/email-code.html"
	if c.Ctx.Input.IsPost() {
		code := helpers.Code{}
		if err := c.ParseForm(&code); err != nil {
			fmt.Println("验证码错误")
			c.Data["json"] = JSONS{"error", "验证码输入有误"}
		} else {
			publicKey := ""
			if email != nil {
				publicKey = models.NewEth().GetPublicKeyByStatus(email.(string))
			}
			c.SetSession("public_key", publicKey)
			fmt.Println("==========")
			fmt.Println(publicKey)
			fmt.Println("==========")
			fmt.Println(code.EmailCode)
			r := helpers.Get()
			defer r.Close()
			if emailcode, _ := redis.String(r.Do("GET", email)); emailcode == code.EmailCode {
				models.NewEth().SetEthEmail(email.(string))
				c.Data["json"] = JSONS{"ok", publicKey}
			} else {
				fmt.Println(emailcode)
				fmt.Println(code.EmailCode)
				c.Data["json"] = JSONS{"error", "验证码输入错误"}
			}
		}
		c.ServeJSON()
	}
}
