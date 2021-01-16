package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"metis-v1.0/conf"
	"time"
)

type TokenStaking struct {
	Id          int    `orm:"pk;auto;"`
	AccountAddr string `orm:"size(50)"`
	DacUid      string `orm:"size(20)"`
	Price       int
	StakingType string    `orm:"size(50)"`
	Status      bool      `default:"0"`
	Ctime       time.Time `orm:"auto_now_add;type(datetime)"`
	Utime       time.Time `orm:"auto_now;type(datetime)"`
}

func (m *TokenStaking) TableName() string {
	return "token_staking"
}

func (m *TokenStaking) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewTokenStaking() *TokenStaking {
	return &TokenStaking{}
}

func (m *TokenStaking) GetToken(AccountAddr string, DacUid string) int {
	num := 0
	o := orm.NewOrm()
	m.AccountAddr = AccountAddr
	m.Price = 10
	m.DacUid = DacUid
	m.StakingType = "register"
	id, err := o.Insert(m)
	if err == nil {
		num = int(id)
	}
	return num
}

func (m *TokenStaking) GetExist(DacUid string, StakingType string) int {
	num := 1
	o := orm.NewOrm()
	var tokenStaking TokenStaking
	err := o.QueryTable("metis_token_staking").Filter("DacUid", DacUid).Filter("StakingType", StakingType).One(&tokenStaking)
	if err == orm.ErrNoRows {
		num = 0
	}
	return num
}

func (m *TokenStaking) GetOwnerStaking(publicKey string) []*TokenStaking {
	var tokenStaking []*TokenStaking
	o := orm.NewOrm()
	_, _ = o.QueryTable("metis_token_staking").Filter("AccountAddr", publicKey).Filter("StakingType", "share").OrderBy("-id").All(&tokenStaking)
	return tokenStaking
}

func (m *TokenStaking) InsertNew(shareAddr string, DacUid string) int {
	num := 1
	o := orm.NewOrm()
	fmt.Println(shareAddr)
	var tokenStaking TokenStaking
	if err := o.QueryTable("metis_token_staking").Filter("DacUid", DacUid).Filter("AccountAddr", shareAddr).Filter("StakingType", "share").One(&tokenStaking); err != orm.ErrNoRows {
		num = 0
	} else {
		m.AccountAddr = shareAddr
		m.DacUid = DacUid
		m.StakingType = "share"
		m.Price = 10
		m.Status = true
		_, _ = o.Insert(m)
	}
	return num
}
