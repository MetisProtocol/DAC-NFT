package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pochard/commons/randstr"
	"math/rand"
	"metis-v1.0/helpers"
	"metis-v1.0/models"
	"net/url"
	"path"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	publicKey := c.GetSession("public_key")
	var dataValue string
	if publicKey != nil {
		dataValue = publicKey.(string)
	}
	c.Data["dac_list"] = models.NewDac().GetDacList()
	num := models.NewDac().CountDac()
	dacNum := map[string]int{
		"dac_num":          num,
		"dac_locked":       2000 - num,
		"dac_num_total":    num * 10,
		"dac_locked_total": 20000 - num*10,
	}
	c.Data["dac_num"] = dacNum
	c.Data["public_key"] = dataValue
	c.TplName = "index.html"

	if c.Ctx.Input.IsPost() {
		dac := helpers.Dac{}
		if err := c.ParseForm(&dac); err != nil {
			c.Data["json"] = JSONS{"error", "Register Failed"}
		} else {
			Uid := models.NewDac().RegisterDac(dac, dataValue)
			c.Data["json"] = JSONS{"ok", Uid}
		}
		c.ServeJSON()
	}
}

func (c *MainController) GetDac() {
	dacUid := c.Ctx.Input.Param(":dac_uid")
	Dac := models.NewDac().GetDac(dacUid)
	shareAddr := c.GetSession("share_addr")
	if c.Ctx.Input.IsPost() {
		fmt.Println("111")
		if Dac.Status == true {
			c.Data["json"] = JSONS{"error", "Already Register Success"}
			fmt.Println("222")
		} else {
			dacMd5 := models.NewDac().SubmitDac(dacUid)
			fmt.Println("333")
			if dacMd5 == "nil" {
				c.Data["json"] = JSONS{"error", "Register Failed"}
				fmt.Println("444")
			} else {
				fmt.Println("555")
				if shareAddr != nil {
					fmt.Println("6666")
					models.NewTokenStaking().InsertNew(shareAddr.(string), Dac.Uid)
					c.DelSession("share_addr")
				}
				c.Data["json"] = JSONS{"ok", dacMd5}
			}
		}
		c.ServeJSON()
	}
	Token := models.NewTokenStaking().GetExist(dacUid, "register")
	num := models.NewDac().CountDac()
	publicKey := c.GetSession("public_key")
	var dataValue string
	if publicKey != nil {
		dataValue = publicKey.(string)
	}
	c.Data["token"] = Token
	c.Data["dac_locked_total"] = 20000 - num*10
	c.Data["dac"] = Dac
	c.Data["public_key"] = dataValue
	c.TplName = "get_dac.html"
}

func (c *MainController) GetGrade() {
	dacUid := c.Ctx.Input.Param(":dac_uid")
	num := models.NewDac().SetGrade(dacUid)
	if num == 0 {
		c.Data["json"] = JSONS{"error", "Get Failed"}
	} else {
		c.Data["json"] = JSONS{"ok", "Get Success"}
	}
	c.ServeJSON()
}

func (c *MainController) GetToken() {
	dacUid := c.Ctx.Input.Param(":dac_uid")
	AccountAddr := c.GetSession("public_key")
	if AccountAddr != nil {
		owner := AccountAddr.(string)
		num := models.NewTokenStaking().GetToken(owner, dacUid)
		if num == 0 {
			c.Data["json"] = JSONS{"error", "Get Failed, DAC Is Not Enough"}
		} else {
			c.Data["json"] = JSONS{"ok", "Get Success"}
		}
	} else {
		c.Data["json"] = JSONS{"error", "Please Login First"}
	}
	c.ServeJSON()
}

func (c *MainController) DacShare() {
	dacMd5 := c.Ctx.Input.Param(":dac_md5")
	Dac := models.NewDac().GetDacByMd5(dacMd5)
	publicKey := c.GetSession("public_key")
	var dataValue string
	if publicKey != nil {
		dataValue = publicKey.(string)
	}
	c.Data["dac"] = Dac
	c.Data["public_key"] = dataValue
	c.Data["dac_public_key"] = dacMd5
	c.TplName = "alert/register_success.html"
}

func (c *MainController) ShareRegister() {
	dacMd5 := c.Ctx.Input.Param(":dac_md5")
	Dac := models.NewDac().GetDacByMd5(dacMd5)
	c.SetSession("share_addr", Dac.DacOwner)
	fmt.Println(c.GetSession("share_addr"))
	publicKey := c.GetSession("public_key")
	var dataValue string
	if publicKey != nil {
		dataValue = publicKey.(string)
	}
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
	c.Data["public_key"] = dataValue
	c.TplName = "share_register.html"
}

func (c *MainController) GetDacName() {
	dacName := c.Ctx.Input.Param(":dac_name")
	dacName, _ = url.QueryUnescape(dacName)
	dacList := [2]string{}
	i := 0
	if GetDacNameFromModel(dacName) {
		myReturn := JSONS{"ok", dacName}
		c.Data["json"] = &myReturn
	} else {
		for {
			dacNameRand := dacName + randstr.RandomAlphanumeric(3)
			if GetDacNameFromModel(dacNameRand) {
				dacList[i] = dacNameRand
				i++
			}
			if i == 2 {
				jsonName, _ := json.Marshal(dacList)
				myReturn := JSONS{"error", string(jsonName)}
				c.Data["json"] = &myReturn
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
	c.SetSession("public_key", publicKey)
	myReturn := JSONS{"ok", publicKey}
	c.Data["json"] = &myReturn
	c.ServeJSON()
}

func (c *MainController) Profile() {
	publicKey := c.GetSession("public_key")
	var dataValue string
	if publicKey != nil {
		dataValue = publicKey.(string)
	}
	c.Data["dac_list"] = models.NewDac().GetAllDacByAccount(dataValue)
	stakingListOwner := models.NewTokenStaking().GetOwnerStaking(dataValue)
	stakingListOther := models.NewTokenStaking().GetOtherStaking(dataValue)
	c.Data["staking_list"] = stakingListOwner
	c.Data["staking_list_other"] = stakingListOther
	c.Data["sum_staking"] = len(stakingListOwner) * 10
	c.Data["sum_staking_other"] = len(stakingListOther) * 10
	c.Data["total_staking"] = len(stakingListOwner)*10 + len(stakingListOther)*10
	c.Data["public_key"] = dataValue
	c.Data["my_dac"] = models.NewDac().GetDacByAccount(dataValue)
	c.Data["get_dac_name"] = c.GetDacNameByUid
	c.TplName = "profile.html"
}

func (c *MainController) File() {
	f, h, _ := c.GetFile("Filedata")
	ext := path.Ext(h.Filename)
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
	fileName := fmt.Sprintf("%x", hashName) + ext
	defer f.Close()
	path := "./static/uploads/" + fileName
	fmt.Println(path)
	a := c.SaveToFile("Filedata", path)
	fmt.Println(a)
	myReturn := JSONS{"ok", fileName}
	c.Data["json"] = &myReturn
	c.ServeJSON()
}

func (c *MainController) GetPrivateKey() {
	publicKey := c.GetSession("public_key")
	if publicKey != nil {
		if email, privateKey := models.NewEth().GetEthEmailKeyByPublicKey(publicKey.(string)); email != "nil" && privateKey != "nil" {
			code := helpers.RequestPrivateKey(privateKey)
			err := SendMail(email, "Metis Team", code)
			if err != nil {
				fmt.Println("发送失败")
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
