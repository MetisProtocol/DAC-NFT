package models

import (
	"github.com/astaxie/beego/orm"
	"metis-v1.0/conf"
	"time"
)

type DacAddr struct {
	Id         int       `orm:"pk;auto;"`
	PublicKey  string    `orm:"size(50);unique"`
	PrivateKey string    `orm:"size(100);unique;"`
	Status     bool      `default:"0"`
	DacId      int       `orm:"null"`
	Ctime      time.Time `orm:"auto_now_add;type(datetime)"`
	Utime      time.Time `orm:"auto_now;type(datetime)"`
}

func (m *DacAddr) TableName() string {
	return "dac_addr"
}

func (m *DacAddr) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewDacAddr() *DacAddr {
	return &DacAddr{}
}

func (m *DacAddr) InsertDacAddr() {
	o := orm.NewOrm()
	keys := map[string]string{
		"0x982a8CbE734cb8c29A6a7E02a3B0e4512148F6F9": "0xd353907ab062133759f149a3afcb951f0f746a65a60f351ba05a3ebf26b67f5c",
		"0xCDC1E53Bdc74bBf5b5F715D6327Dca5785e228B4": "0x971c58af72fd8a158d4e654cfbe98f5de024d28547005909684f58c9c46a25c4",
		"0xf5d1EAF516eF3b0582609622A221656872B82F78": "0x85d168288e7fcf84b1841e447fc7945b1e27bfe9a3776367079a6427405eac66",
		"0xf8eA26C3800D074a11bf814dB9a0735886C90197": "0xf3da3ac70552606ed09d16dd2808c924826094f0c5cbfcb4f2e0e1cfc70ff8dd",
		"0x2647116f9304abb9F0B7aC29aBC0D9aD540506C8": "0xbf20e9c05d70ce59a6b125eab3b4122eb75044a33749c4c5a77e3b0b86fa091e",
		"0x80a32A0E5cA81b5a236168C21532B32e3cBC95e2": "0x5d4af11a54d4a5196b0073ba26a1114cb113e1339d9354c8165b8e181c89cad9",
		"0x47f55A2ACe3b84B0F03717224DBB7D0Df4351658": "0xef78746d079c9d72d2e9a3c10447d1d4aaae6a51541d0296da4fc9ec7e060aff",
		"0xC817898296b27589230B891f144dd71A892b0C18": "0xc95286117cd74213417aeca52118ccd03ec240582f0a9a3e4ef7b434523179f3",
	}
	for k, v := range keys {
		var dacAddr DacAddr
		dacAddr.PublicKey = k
		dacAddr.PrivateKey = v
		_, _ = o.Insert(&dacAddr)
	}
}
