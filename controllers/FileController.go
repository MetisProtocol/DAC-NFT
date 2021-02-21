package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"path"
	"time"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) File() {
	f, h, _ := c.GetFile("Filedata")
	ext := path.Ext(h.Filename)
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
	fileName := fmt.Sprintf("%x", hashName) + ext
	defer f.Close()
	filePath := "./static/uploads/" + fileName
	_ = c.SaveToFile("Filedata", filePath)
	myReturn := JSONS{"ok", fileName}
	c.Data["json"] = &myReturn
	c.ServeJSON()
}
