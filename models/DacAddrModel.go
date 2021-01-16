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
		"0x64E078A8Aa15A41B85890265648e965De686bAE6": "0x0874049f95d55fb76916262dc70571701b5c4cc5900c0691af75f1a8a52c8268",
		"0x2F560290FEF1B3Ada194b6aA9c40aa71f8e95598": "0x21d7212f3b4e5332fd465877b64926e3532653e2798a11255a46f533852dfe46",
		"0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1": "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
		"0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0": "0x6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1",
		"0x22d491Bde2303f2f43325b2108D26f1eAbA1e32b": "0x6370fd033278c143179d81c5526140625662b8daa446c22ee2d73db3707e620c",
		"0xE11BA2b4D45Eaed5996Cd0823791E0C93114882d": "0x646f1ce2fdad0e6deeeb5c7e8e5543bdde65e86029e2fd9fc169899c440a7913",
		"0xd03ea8624C8C5987235048901fB614fDcA89b117": "0xadd53f9a7e588d003326d1cbf9e4a43c061aadd9bc938c843a79e7b4fd2ad743",
		"0x95cED938F7991cd0dFcb48F0a06a40FA1aF46EBC": "0x395df67f0c2d2d9fe1ad08d1bc8b6627011959b79c53d7dd6a3536a33ab8a4fd",
		"0x73EB6d82CFB20bA669e9c178b718d770C49BB52f": "0xd09ba371c359f10f22ccda12fd26c598c7921bda3220c9942174562bc6a36fe8",
		"0x9D8E5fAc117b15DaCED7C326Ae009dFE857621f1": "0x2d2719c6a828911ed0c50d5a6c637b63353e77cf57ea80b8e90e630c4687e9c5",
		"0x0D38e653eC28bdea5A2296fD5940aaB2D0B8875c": "0x21118f9a6de181061a2abd549511105adb4877cf9026f271092e6813b7cf58ab",
		"0x1B569e8f1246907518Ff3386D523dcF373e769B6": "0x1166189cdf129cdcb011f2ad0e5be24f967f7b7026d162d7c36073b12020b61c",
	}
	for k, v := range keys {
		var dacAddr DacAddr
		dacAddr.PublicKey = k
		dacAddr.PrivateKey = v
		_, _ = o.Insert(&dacAddr)
	}
}
