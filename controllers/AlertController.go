package controllers

import (
	"github.com/gomodule/redigo/redis"
	"metis-v1.0/helpers"
	"metis-v1.0/models"
)

type AlertController struct {
	BaseController
}

func (c *AlertController) Prepare() {
	c.BaseController.Prepare()
}

func (c *AlertController) Wallet() {
	c.TplName = "alert/wallet.html"
}
func (c *AlertController) EmailRegister() {
	c.TplName = "alert/email-register.html"
	if c.Ctx.Input.IsPost() {
		email := helpers.Email{}
		if err := c.ParseForm(&email); err != nil {
			c.JsonResult("error", "error", "Send Failed")
		} else {
			code := helpers.GetRandomString(6)
			c.SetSession("email", email.ToEmail)
			r := helpers.Get()
			defer r.Close()
			_, erre := r.Do("SET", email.ToEmail, code)
			if erre != nil {
				c.JsonResult("error", "Send Failed")
			} else {
				code = helpers.RegisterEmail(code)
				err := SendMail(email.ToEmail, "Metis Team", code)
				if err != nil {
					c.JsonResult("error", "Send Failed")
				} else {
					c.JsonResult("ok", "Send Success")
				}
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
			c.JsonResult("error", "Wrong Code")
		} else {
			publicKey := ""
			if email != nil {
				publicKey = models.NewEth().GetPublicKeyByStatus(email.(string))
			}
			c.SetSession("public_key", publicKey)
			r := helpers.Get()
			defer r.Close()
			if emailCode, _ := redis.String(r.Do("GET", email)); emailCode == code.EmailCode {
				_, newMember := models.NewEth().SetEthEmail(email.(string))
				if newMember == 1 {
					successCode := helpers.RegisterSuccessEmail(email.(string), publicKey)
					_ = SendMail(email.(string), "Metis Team", successCode)
				}
				c.JsonResult("ok", publicKey)
			} else {
				c.JsonResult("error", "Wrong Code")
			}
		}
		c.ServeJSON()
	}
}
