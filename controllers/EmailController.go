package controllers

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo string, subject string, body string) error {

	mailConn := map[string]string{
		"user": "mydesmond@163.com",
		"pass": "Misswang0406",
		"host": "smtp.163.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "Metis")) //这种方式可以添加别名，即“XX官方”
	m.SetHeader("To", mailTo)                                       //发送给多个用户
	m.SetHeader("Subject", subject)                                 //设置邮件主题
	m.SetBody("text/html", body)                                    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
