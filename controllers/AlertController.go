package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"metis-v1.0/helpers"
	"metis-v1.0/models"
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
			fmt.Println("Send Failed")
			c.Data["json"] = JSONS{"error", "Send Failed"}
		} else {
			code := helpers.GetRandomString(6)
			c.SetSession("email", email.ToEmail)
			r := helpers.Get()
			defer r.Close()
			_, erre := r.Do("SET", email.ToEmail, code)
			if erre != nil {
				fmt.Println(erre)
				c.Data["json"] = JSONS{"error", "Send Failed"}
			} else {
				code = helpers.RegisterEmail(code)
				//if num := models.NewDac().GetDacByEmail(email.ToEmail); num != 1 {
				//	PublicKey := models.NewEth().GetPublicKey(email.ToEmail)
				//	if PublicKey == "nil" {
				//		c.Data["json"] = JSONS{"error", "发送失败"}
				//	} else {
				//		code = code + " " + PublicKey
				//	}
				//}
				err := SendMail(email.ToEmail, "Metis Team", code)
				if err != nil {
					fmt.Println("发送失败")
					c.Data["json"] = JSONS{"error", "Send Failed"}
				} else {
					c.Data["json"] = JSONS{"ok", "Send Success"}
				}
				fmt.Println("发送成功")
			}
		}
		c.ServeJSON()
	}
}
func (c *AlertController) EmailLogin() {
	c.TplName = "alert/register_success.html"
}
func (c *AlertController) EmailCode() {
	email := c.GetSession("email")
	c.Data["email"] = email
	c.TplName = "alert/email-code.html"
	if c.Ctx.Input.IsPost() {
		code := helpers.Code{}
		if err := c.ParseForm(&code); err != nil {
			fmt.Println("Code is Wrong")
			c.Data["json"] = JSONS{"error", "Wrong Code"}
		} else {
			publicKey := ""
			if email != nil {
				publicKey = models.NewEth().GetPublicKeyByStatus(email.(string))
			}
			c.SetSession("public_key", publicKey)
			r := helpers.Get()
			defer r.Close()
			if emailcode, _ := redis.String(r.Do("GET", email)); emailcode == code.EmailCode {
				_, newMember := models.NewEth().SetEthEmail(email.(string))
				if newMember == 1 {
					successCode := helpers.RegisterSuccessEmail(publicKey)
					_ = SendMail(email.(string), "Metis Team", successCode)
				}
				c.Data["json"] = JSONS{"ok", publicKey}
			} else {
				fmt.Println(emailcode)
				fmt.Println(code.EmailCode)
				c.Data["json"] = JSONS{"error", "Wrong Code"}
			}
		}
		c.ServeJSON()
	}
}
