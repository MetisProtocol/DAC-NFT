package controllers

import (
	"compress/gzip"
	"encoding/json"
	"github.com/astaxie/beego"
	"io"
	"metis-v1.0/helpers"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	publicKey string
}

type CookieRemember struct {
	publicKey string
	Time      time.Time
}

func (this *BaseController) Prepare() {
	publicKey := this.GetSession("public_key")
	if publicKey != nil {
		this.publicKey = publicKey.(string)
	} else {
		this.publicKey = ""
		if cookie, ok := this.GetSecureCookie(beego.AppConfig.String("key"), "publicKey"); ok {
			var remember CookieRemember
			err := helpers.Decode(cookie, &remember)
			if err == nil {
				this.publicKey = remember.publicKey
			}
		}
	}
}

func (this *BaseController) JsonResult(Code string, Message string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["Code"] = Code
	jsonData["Message"] = Message

	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	//this.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")//解决回退出现json的问题
	//使用gzip原始，json数据会只有原本数据的10分之一左右
	if strings.Contains(strings.ToLower(this.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
		this.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		//gzip压缩
		w := gzip.NewWriter(this.Ctx.ResponseWriter)
		defer w.Close()
		w.Write(returnJSON)
		w.Flush()
	} else {
		io.WriteString(this.Ctx.ResponseWriter, string(returnJSON))
	}
	this.StopRun()
}
