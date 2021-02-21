package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pochard/commons/randstr"
	"metis-v1.0/helpers"
	"metis-v1.0/models"
	"net/url"
	"time"
)

type MainController struct {
	BaseController
}

func (c *MainController) Prepare() {
	c.BaseController.Prepare()
}

func (c *MainController) Index() {
	c.Data["dac_list"] = models.NewDac().GetDacList()
	num := models.NewDac().CountDac()
	dacNum := map[string]int{
		"dac_num":          num,
		"dac_locked":       2000 - num,
		"dac_num_total":    num * 10,
		"dac_locked_total": 20000 - num*10,
	}
	c.Data["dac_num"] = dacNum
	c.Data["public_key"] = c.publicKey
	c.TplName = "index.html"

	if c.Ctx.Input.IsPost() {
		dac := helpers.Dac{}
		if err := c.ParseForm(&dac); err != nil {
			c.JsonResult("error", "Register Failed")
		} else {
			Uid := models.NewDac().RegisterDac(dac, c.publicKey)
			c.JsonResult("ok", Uid)
		}
	}
}

func (c *MainController) GetDac() {
	dacUid := c.Ctx.Input.Param(":dac_uid")
	Dac := models.NewDac().GetDac(dacUid)
	shareAddr := c.GetSession("share_addr")
	if c.Ctx.Input.IsPost() {
		if Dac.Status == true {
			c.JsonResult("error", "Already Register Success")
		} else {
			dacMd5 := models.NewDac().SubmitDac(dacUid)
			if dacMd5 == "nil" {
				c.JsonResult("error", "Register Failed")
			} else {
				if shareAddr != nil {
					models.NewTokenStaking().InsertNew(shareAddr.(string), Dac.Uid)
					c.DelSession("share_addr")
				}
				c.JsonResult("ok", dacMd5)
			}
		}
	}
	Token := models.NewTokenStaking().GetExist(dacUid, "register")
	num := models.NewDac().CountDac()
	c.Data["token"] = Token
	c.Data["dac_locked_total"] = 20000 - num*10
	c.Data["dac"] = Dac
	c.Data["public_key"] = c.publicKey
	c.TplName = "get_dac.html"
}

func (c *MainController) GetGrade() {
	dacUid := c.Ctx.Input.Param(":dac_uid")
	num := models.NewDac().SetGrade(dacUid)
	if num == 0 {
		c.JsonResult("error", "Get Failed")
	} else {
		c.JsonResult("ok", "Get Success")
	}
}

func (c *MainController) GetToken() {
	dacUid := c.Ctx.Input.Param(":dac_uid")
	if c.publicKey != "nil" {
		owner := c.publicKey
		num := models.NewTokenStaking().GetToken(owner, dacUid)
		if num == 0 {
			c.JsonResult("error", "Get Failed, DAC Is Not Enough")
		} else {
			c.JsonResult("ok", "Get Success")
		}
	} else {
		c.JsonResult("error", "Please Login First")
	}
}

func (c *MainController) DacShare() {
	dacMd5 := c.Ctx.Input.Param(":dac_md5")
	Dac := models.NewDac().GetDacByMd5(dacMd5)
	c.Data["dac"] = Dac
	c.Data["public_key"] = c.publicKey
	c.Data["dac_public_key"] = dacMd5
	c.TplName = "alert/register_success.html"
}

func (c *MainController) ShareRegister() {
	dacMd5 := c.Ctx.Input.Param(":dac_md5")
	Dac := models.NewDac().GetDacByMd5(dacMd5)
	c.SetSession("share_addr", Dac.DacOwner)
	num := models.NewDac().CountDac()
	dacNum := map[string]int{
		"dac_num":          num,
		"dac_locked":       2000 - num,
		"dac_num_total":    num * 10,
		"dac_locked_total": 20000 - num*10,
	}
	c.Data["dac_list"] = models.NewDac().GetDacListLimit()
	c.Data["dac_num"] = dacNum
	c.Data["dac"] = Dac
	c.Data["public_key"] = c.publicKey
	c.TplName = "share_register.html"
}

func (c *MainController) GetDacName() {
	dacName := c.Ctx.Input.Param(":dac_name")
	dacName, _ = url.QueryUnescape(dacName)
	dacList := [2]string{}
	i := 0
	if GetDacNameFromModel(dacName) {
		c.JsonResult("ok", dacName)
	} else {
		for {
			dacNameRand := dacName + randstr.RandomAlphanumeric(3)
			if GetDacNameFromModel(dacNameRand) {
				dacList[i] = dacNameRand
				i++
			}
			if i == 2 {
				jsonName, _ := json.Marshal(dacList)
				c.JsonResult("error", string(jsonName))
				break
			}
		}
	}
	c.ServeJSON()
}

func GetDacNameFromModel(dacName string) bool {
	return models.NewDac().GetDacByName(dacName)
}

type JSONS struct {
	Code    string
	Message string
}

func (c *MainController) SetUserSession() {
	publicKey := c.Ctx.Input.Param(":public_key")
	var remember CookieRemember
	c.SetSession("public_key", publicKey)
	remember.publicKey = publicKey
	remember.Time = time.Now()
	v, err := helpers.Encode(remember)
	if err == nil {
		c.SetSecureCookie(beego.AppConfig.String("key"), "publicKey", v, 24*3600*365)
	}
	c.JsonResult("ok", publicKey)
}

func (c *MainController) Profile() {
	c.Data["dac_list"] = models.NewDac().GetAllDacByAccount(c.publicKey)
	stakingListOwner := models.NewTokenStaking().GetOwnerStaking(c.publicKey)
	stakingListOther := models.NewTokenStaking().GetOtherStaking(c.publicKey)
	c.Data["staking_list"] = stakingListOwner
	c.Data["staking_list_other"] = stakingListOther
	c.Data["sum_staking"] = len(stakingListOwner) * 10
	c.Data["sum_staking_other"] = len(stakingListOther) * 10
	c.Data["total_staking"] = len(stakingListOwner)*10 + len(stakingListOther)*10
	c.Data["public_key"] = c.publicKey
	c.Data["my_dac"] = models.NewDac().GetDacByAccount(c.publicKey)
	c.Data["get_dac_name"] = c.GetDacNameByUid
	c.TplName = "profile.html"
}

func (c *MainController) GetPrivateKey() {
	if c.publicKey != "nil" {
		if email, privateKey := models.NewEth().GetEthEmailKeyByPublicKey(c.publicKey); email != "nil" && privateKey != "nil" {
			code := helpers.RequestPrivateKey(privateKey)
			err := SendMail(email, "Metis Team", code)
			if err != nil {
				c.Data["json"] = JSONS{"error", "Send Failed"}
			} else {
				c.Data["json"] = JSONS{"ok", "The private key will be sent to you via your email address. Please save it somewhere safe and secret!"}
			}
			fmt.Println("发送成功")
		}
	} else {
		c.Data["json"] = JSONS{"error", "Login First"}
	}
	c.ServeJSON()
}

func (c *MainController) GetDacNameByUid(dacUid string) string {
	dac := models.NewDac().GetDac(dacUid)
	return dac.DacName
}

func (c *MainController) Logout() {
	c.DelSession("public_key")
	c.DelSession("email")
	c.DelSession("share_addr")
	c.Data["json"] = JSONS{"ok", "Logout"}
	c.ServeJSON()
}

func (c *MainController) Test1() {
	models.NewEth().InsertEth()
	models.NewDacAddr().InsertDacAddr()
	c.TplName = "test/test1.html"
}
func (c *MainController) Test2() {
	models.NewBlackList().InsertBlackList()
	c.TplName = "test/test2.html"
}
func (c *MainController) Test3() {
	c.TplName = "test/test3.html"
}
func (c *MainController) Test4() {
	c.TplName = "test/test4.html"
}
func (c *MainController) Test5() {
	c.TplName = "test/test5.html"
}
func (c *MainController) Test6() {
	c.TplName = "test/test6.html"
}
func (c *MainController) Test7() {
	c.TplName = "test/test7.html"
}
func (c *MainController) Test8() {
	c.TplName = "test/test8.html"
}
