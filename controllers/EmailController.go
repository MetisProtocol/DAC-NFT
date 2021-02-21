package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo string, subject string, body string) error {

	mailConn := map[string]string{
		"user": beego.AppConfig.String("email_user"),
		"pass": beego.AppConfig.String("email_pass"),
		"host": beego.AppConfig.String("email_host"),
		"port": beego.AppConfig.String("db_database"),
	}

	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "Metis"))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
