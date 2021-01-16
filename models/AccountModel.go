package models

import (
	"github.com/astaxie/beego/orm"
	"metis-v1.0/conf"
	"metis-v1.0/helpers"
)

type Users struct {
	Id       int    `orm:"pk;auto;"`
	Uid      string `orm:"size(16)"`
	UserName string `orm:"size(16);unique;"`
	Password string `orm:"size(32);"`
	Email    string `orm:"size(32);unique"`
	Phone    string `orm:"size(32);unique"`
	Salt     string `orm:"size(16);"`
}

func (m *Users) TableName() string {
	return "users"
}

func (m *Users) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewUser() *Users {
	return &Users{}
}

func (m *Users) Register(user helpers.User) error {
	o := orm.NewOrm()
	m.Uid = "A00002"
	m.UserName = user.UserName
	m.Password = user.Password
	m.Email = user.Email
	m.Phone = "13111111111"
	m.Salt = "adbs"
	o.Insert(m)
	return nil
}
