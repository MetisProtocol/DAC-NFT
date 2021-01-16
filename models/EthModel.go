package models

import (
	"github.com/astaxie/beego/orm"
	"metis-v1.0/conf"
	"time"
)

type Eth struct {
	Id         int       `orm:"pk;auto;"`
	PublicKey  string    `orm:"size(50);unique"`
	PrivateKey string    `orm:"size(100);unique;"`
	Status     bool      `default:"0"`
	Email      string    `orm:"size(50);null;"`
	Ctime      time.Time `orm:"auto_now_add;type(datetime)"`
	Utime      time.Time `orm:"auto_now;type(datetime)"`
}

func (m *Eth) TableName() string {
	return "eth"
}

func (m *Eth) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewEth() *Eth {
	return &Eth{}
}

func (m *Eth) GetPublicKey(Email string) string {
	o := orm.NewOrm()
	PublicKey := "nil"
	var eth Eth
	err := o.QueryTable("metis_eth").Filter("Email", Email).Filter("Status", false).One(&eth)
	if err == orm.ErrNoRows {
		PublicKey = m.GetLastEth(Email)
	} else {
		PublicKey = eth.PublicKey
	}
	return PublicKey
}

func (m *Eth) GetPublicKeyByStatus(Email string) string {
	o := orm.NewOrm()
	PublicKey := "nil"
	var eth Eth
	err := o.QueryTable("metis_eth").Filter("Email", Email).One(&eth)
	if err == orm.ErrNoRows {
		PublicKey = m.GetLastEth(Email)
	} else {
		PublicKey = eth.PublicKey
	}
	return PublicKey
}

func (m *Eth) GetLastEth(Email string) string {
	PublicKey := "nil"
	o := orm.NewOrm()
	var eth Eth
	err := o.QueryTable("metis_eth").Filter("Status", false).One(&eth)
	if err != orm.ErrNoRows {
		PublicKey = eth.PublicKey
		eth.Email = Email
		_, _ = o.Update(&eth)
	}
	return PublicKey
}

func (m *Eth) SetEthEmail(Email string) string {
	PublicKey := "nil"
	o := orm.NewOrm()
	var eth Eth
	err := o.QueryTable("metis_eth").Filter("Email", Email).One(&eth)
	if err != orm.ErrNoRows {
		PublicKey = eth.PublicKey
		eth.Status = true
		_, _ = o.Update(&eth)
	}
	return PublicKey
}

func (m *Eth) GetEthEmailByPublicKey(publicKey string) string {
	email := "nil"
	o := orm.NewOrm()
	var eth Eth
	err := o.QueryTable("metis_eth").Filter("PublicKey", publicKey).Filter("Status", true).One(&eth)
	if err != orm.ErrNoRows {
		email = eth.Email
	}
	return email
}

func (m *Eth) InsertEth() {
	o := orm.NewOrm()
	keys := map[string]string{
		"0xC7739909e08A9a0F303A010d46658Bdb4d5a6786": "0x86117111fcb34df8d0e58505969021b9308513c6e94d16172f0c8789a7130a43",
		"0x99cce66d3A39C2c2b83AfCefF04c5EC56E9B2A58": "0xdcb8686c211c231be763f0a95cc02227a707643fd2631bda99fcdbd03cd9ca3d",
		"0x4b930E7b3E491e37EaB48eCC8a667c59e307ef20": "0xb74ffec4abd7e93889196054d5e6ed8ea9c1c3314e77a74c00f851c47f5268fd",
		"0x02233B22860f810E32fB0751f368fE4ef21A1C05": "0xba30972105ec13423116d2e5c11a8d282805ac3654bb4c1c2f5fa63f4da42dad",
		"0x89c1D413758F8339Ade263E6e6bC072F1d429f32": "0x87ad1798a2d32434f72598575237528a435416da1bdc900025c415903647957e",
		"0x61bBB5135b43F03C96570616d6d3f607b7103111": "0x5d4af11a54d4a5196b0073ba26a1114cb113e1339d9354c8165b8e181c89cad9",
		"0x8C4cE7a10A4e38EE96feD47C628Be1FfA57Ab96e": "0xa03bf2b145b0154c2e788a1d4642d235f6ff1c8aceeb41d0d7232525da8bdb77",
		"0x73EB6d82CFB20bA669e9c178b718d770C49BB52f": "0xd09ba371c359f10f22ccda12fd26c598c7921bda3220c9942174562bc6a36fe8",
		"0x9D8E5fAc117b15DaCED7C326Ae009dFE857621f1": "0x2d2719c6a828911ed0c50d5a6c637b63353e77cf57ea80b8e90e630c4687e9c5",
	}
	for k, v := range keys {
		var eth Eth
		eth.PublicKey = k
		eth.PrivateKey = v
		_, _ = o.Insert(&eth)
	}
}
