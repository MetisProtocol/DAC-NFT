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
	"path"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	publicKey := c.GetSession("public_key")
	fmt.Println("====")
	fmt.Println(c.GetSession("share_addr"))
	//c.DelSession("share_addr")
	var dataValue string
	var owner = "nil"
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
			fmt.Println("注册失败")
			c.Data["json"] = JSONS{"error", "注册失败"}
		} else {
			fmt.Println("注册成功")
			Uid := models.NewDac().RegisterDac(dac, owner)
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
		if Dac.Status == true {
			c.Data["json"] = JSONS{"error", "您已经注册成功了"}
		} else {
			publicKey := models.NewDac().SubmitDac(dacUid)
			if publicKey == "nil" {
				c.Data["json"] = JSONS{"error", "注册失败"}
			} else {
				if shareAddr != nil {
					models.NewTokenStaking().InsertNew(shareAddr.(string), Dac.Uid)
					c.DelSession("share_addr")
				}
				c.Data["json"] = JSONS{"ok", publicKey}
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
		c.Data["json"] = JSONS{"error", "获取失败"}
	} else {
		c.Data["json"] = JSONS{"ok", "获取成功"}
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
			c.Data["json"] = JSONS{"error", "获取失败,Token已经分配完毕"}
		} else {
			c.Data["json"] = JSONS{"ok", "获取成功"}
		}
	} else {
		c.Data["json"] = JSONS{"error", "获取失败,请先登录"}
	}
	c.ServeJSON()
}

func (c *MainController) DacShare() {
	dacAddr := c.Ctx.Input.Param(":dac_addr")
	Dac := models.NewDac().GetDacByAddr(dacAddr)
	publicKey := c.GetSession("public_key")
	var dataValue string
	if publicKey != nil {
		dataValue = publicKey.(string)
	}
	c.Data["dac"] = Dac
	c.Data["public_key"] = dataValue
	c.Data["dac_public_key"] = dacAddr
	c.TplName = "dac_share.html"
}

func (c *MainController) ShareRegister() {
	dacAddr := c.Ctx.Input.Param(":dac_addr")
	Dac := models.NewDac().GetDacByAddr(dacAddr)
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
	c.Data["dac_num"] = dacNum
	c.Data["dac"] = Dac
	c.Data["public_key"] = dataValue
	c.TplName = "share_register.html"
}

func (c *MainController) GetDacName() {
	dacName := c.Ctx.Input.Param(":dac_name")
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
	stakingList := models.NewTokenStaking().GetOwnerStaking(dataValue)
	c.Data["staking_list"] = stakingList
	c.Data["sum_staking"] = len(stakingList) * 10
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

func (c *MainController) GetDacNameByUid(dacUid string) string {
	dac := models.NewDac().GetDac(dacUid)
	return dac.DacName
}

func (c *MainController) Logout() {
	c.DelSession("public_key")
	c.DelSession("email")
	c.DelSession("share_addr")
	c.Data["json"] = JSONS{"ok", "退出成功"}
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
