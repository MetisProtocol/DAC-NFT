package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"metis-v1.0/conf"
	"metis-v1.0/helpers"
	"strconv"
	"time"
)

type Dac struct {
	Id         int       `orm:"pk;auto;"`
	Uid        string    `orm:"size(16);unique"`
	DacName    string    `orm:"size(50);unique;"`
	DacAddr    string    `orm:"size(50);null;"`
	DacProduce string    `orm:"size(300);"`
	DacLogo    string    `orm:"size(100);"`
	DacGrade   string    `orm:"size(30);null;"`
	DacChain   int       `orm:"null"`
	DacMd5     string    `orm:"size(32);null;"`
	DacOwner   string    `orm:"size(50);"`
	OwnerEmail string    `orm:"size(50);null"`
	Status     bool      `default:"0"`
	Ctime      time.Time `orm:"auto_now_add;type(datetime)"`
	Utime      time.Time `orm:"auto_now;type(datetime)"`
}

func (m *Dac) TableName() string {
	return "dac"
}

func (m *Dac) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewDac() *Dac {
	return &Dac{}
}

func (m *Dac) RegisterDac(dac helpers.Dac, dacOwner string) string {
	o := orm.NewOrm()
	m.Uid = m.GetLastDac()
	m.DacName = dac.DacName
	m.DacLogo = dac.DacLogo
	m.DacProduce = dac.DacProduce
	m.DacOwner = dacOwner
	if email := NewEth().GetEthEmailByPublicKey(dacOwner); email != "nil" {
		m.OwnerEmail = email
	}
	_, _ = o.Insert(m)
	return m.Uid
}

func (m *Dac) GetDac(dacUid string) Dac {
	var dac Dac
	o := orm.NewOrm()
	_ = o.QueryTable("metis_dac").Filter("Uid", dacUid).One(&dac)
	return dac
}

func (m *Dac) GetDacByEmail(Email string) int {
	num := 1
	var dac Dac
	o := orm.NewOrm()
	err := o.QueryTable("metis_dac").Filter("OwnerEmail", Email).One(&dac)
	if err == orm.ErrNoRows {
		num = 0
	}
	return num
}

func (m *Dac) GetDacByAddr(dacAddr string) Dac {
	var dac Dac
	o := orm.NewOrm()
	_ = o.QueryTable("metis_dac").Filter("DacAddr", dacAddr).One(&dac)
	return dac
}

func (m *Dac) GetDacByMd5(dacMd5 string) Dac {
	var dac Dac
	o := orm.NewOrm()
	_ = o.QueryTable("metis_dac").Filter("DacMd5", dacMd5).One(&dac)
	return dac
}

func (m *Dac) GetDacByAccount(accountAddr string) Dac {
	var dac Dac
	o := orm.NewOrm()
	_ = o.QueryTable("metis_dac").Filter("DacOwner", accountAddr).One(&dac)
	return dac
}

func (m *Dac) GetAllDacByAccount(accountAddr string) []*Dac {
	var dac []*Dac
	o := orm.NewOrm()
	_, _ = o.QueryTable("metis_dac").Filter("DacOwner", accountAddr).OrderBy("-id").All(&dac)
	return dac
}

func (m *Dac) SetGrade(dacUid string) int {
	num := 0
	var dac Dac
	o := orm.NewOrm()
	err := o.QueryTable("metis_dac").Filter("Uid", dacUid).One(&dac)
	if err == nil {
		dac.DacGrade = "GENESIS"
		if _num, err := o.Update(&dac); err == nil {
			num = int(_num)
		}
	}
	return num
}

func (m *Dac) SubmitDac(dacUid string) string {
	DacMd5 := "nil"
	var dac Dac
	var dacAddr DacAddr
	o := orm.NewOrm()
	err := o.QueryTable("metis_dac").Filter("Uid", dacUid).One(&dac)
	err1 := o.QueryTable("metis_dac_addr").Filter("Status", false).One(&dacAddr)
	if err == nil && err1 == nil {
		dac.Status = true
		dac.DacAddr = dacAddr.PublicKey
		dac.DacMd5 = helpers.Md5(dac.DacAddr)
		dacAddr.Status = true
		dacAddr.DacId = dac.Id
		_, err2 := o.Update(&dac)
		_, err3 := o.Update(&dacAddr)
		if err2 == nil && err3 == nil {
			DacMd5 = dac.DacMd5
		}
	}
	return DacMd5
}

func (m *Dac) GetLastDacStatus() int {
	var dac Dac
	var dacChain int
	o := orm.NewOrm()
	err := o.QueryTable("metis_dac").Filter("Status", true).OrderBy("-id").One(&dac)
	if err == orm.ErrNoRows {
		dacChain = 0
	} else {
		dacChain = dac.DacChain
	}
	return dacChain
}

func (m *Dac) GetDacList() []*Dac {
	var dac []*Dac
	o := orm.NewOrm()
	o.QueryTable("metis_dac").Filter("Status", true).OrderBy("-id").All(&dac)
	return dac
}

func (m *Dac) GetDacListLimit() []*Dac {
	var dac []*Dac
	o := orm.NewOrm()
	o.QueryTable("metis_dac").Filter("Status", true).OrderBy("-id").Limit(10).All(&dac)
	return dac
}

func (m *Dac) CountDac() int {
	o := orm.NewOrm()
	num, _ := o.QueryTable("metis_dac").Filter("Status", true).Count()
	return int(num)
}

func (m *Dac) GetLastDac() string {
	var dac Dac
	var dacUid string
	o := orm.NewOrm()
	err := o.QueryTable("metis_dac").OrderBy("-id").One(&dac)
	if err == orm.ErrNoRows {
		dacUid = "D00001"
	} else {
		i, _ := strconv.Atoi(dac.Uid[1:])
		i++
		dacUid = fmt.Sprintf("D%05d", i)
	}
	return dacUid
}

func (m *Dac) GetDacByName(dacName string) bool {
	var dac Dac
	var registered bool
	o := orm.NewOrm()
	err := o.QueryTable("metis_dac").Filter("DacName__iexact", dacName).One(&dac)
	if err == orm.ErrNoRows {
		registered = true
	} else {
		registered = false
	}
	return registered
}
