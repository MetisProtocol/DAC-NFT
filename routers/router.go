package routers

import (
	"github.com/astaxie/beego"
	"metis-v1.0/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/set_session/:public_key", &controllers.MainController{}, "get:SetUserSession")
	beego.Router("/file", &controllers.FileController{}, "post:File")
	beego.Router("/logout", &controllers.MainController{}, "get:Logout")
	beego.Router("/get_dac/:dac_uid", &controllers.MainController{}, "*:GetDac")
	beego.Router("/set_grade/:dac_uid", &controllers.MainController{}, "get:GetGrade")
	beego.Router("/get_token/:dac_uid", &controllers.MainController{}, "get:GetToken")
	beego.Router("/dac_share/:dac_md5", &controllers.MainController{}, "get:DacShare")
	beego.Router("/profile", &controllers.MainController{}, "get:Profile")
	beego.Router("/share_register/:dac_md5", &controllers.MainController{}, "get:ShareRegister")
	beego.Router("/check_name/:dac_name", &controllers.MainController{}, "*:GetDacName")
	beego.Router("/alert/wallet", &controllers.AlertController{}, "get:Wallet")
	beego.Router("/alert/email-register", &controllers.AlertController{}, "*:EmailRegister")
	beego.Router("/request_private", &controllers.MainController{}, "get:GetPrivateKey")
	beego.Router("/alert/email-code", &controllers.AlertController{}, "*:EmailCode")
	beego.Router("/1", &controllers.MainController{}, "get:Test1")
	beego.Router("/2", &controllers.MainController{}, "get:Test2")
	beego.Router("/3", &controllers.MainController{}, "get:Test3")
	beego.Router("/4", &controllers.MainController{}, "get:Test4")
	beego.Router("/5", &controllers.MainController{}, "get:Test5")
	beego.Router("/6", &controllers.MainController{}, "get:Test6")
	beego.Router("/7", &controllers.MainController{}, "get:Test7")
	beego.Router("/8", &controllers.MainController{}, "get:Test8")
	beego.Router("/register", &controllers.AccountController{}, "*:Register")
}
