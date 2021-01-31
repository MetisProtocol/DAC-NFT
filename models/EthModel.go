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

func (m *Eth) SetEthEmail(Email string) (string, int) {
	PublicKey := "nil"
	newMember := 0
	o := orm.NewOrm()
	var eth Eth
	err := o.QueryTable("metis_eth").Filter("Email", Email).Filter("Status", false).One(&eth)
	if err != orm.ErrNoRows {
		PublicKey = eth.PublicKey
		eth.Status = true
		_, _ = o.Update(&eth)
		newMember = 1
	}
	return PublicKey, newMember
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

func (m *Eth) GetEthEmailKeyByPublicKey(publicKey string) (string, string) {
	email := "nil"
	privateKey := "nil"
	o := orm.NewOrm()
	var eth Eth
	err := o.QueryTable("metis_eth").Filter("PublicKey", publicKey).Filter("Status", true).One(&eth)
	if err != orm.ErrNoRows {
		email = eth.Email
		privateKey = eth.PrivateKey
	}
	return email, privateKey
}

func (m *Eth) InsertEth() {
	o := orm.NewOrm()
	keys := map[string]string{
		//"0xC7739909e08A9a0F303A010d46658Bdb4d5a6786": "0x86117111fcb34df8d0e58505969021b9308513c6e94d16172f0c8789a7130a43",
		//"0x99cce66d3A39C2c2b83AfCefF04c5EC56E9B2A58": "0xdcb8686c211c231be763f0a95cc02227a707643fd2631bda99fcdbd03cd9ca3d",
		//"0x4b930E7b3E491e37EaB48eCC8a667c59e307ef20": "0xb74ffec4abd7e93889196054d5e6ed8ea9c1c3314e77a74c00f851c47f5268fd",
		//"0x02233B22860f810E32fB0751f368fE4ef21A1C05": "0xba30972105ec13423116d2e5c11a8d282805ac3654bb4c1c2f5fa63f4da42dad",
		//"0x89c1D413758F8339Ade263E6e6bC072F1d429f32": "0x87ad1798a2d32434f72598575237528a435416da1bdc900025c415903647957e",
		//"0x61bBB5135b43F03C96570616d6d3f607b7103111": "0x5d4af11a54d4a5196b0073ba26a1114cb113e1339d9354c8165b8e181c89cad9",
		//"0x8C4cE7a10A4e38EE96feD47C628Be1FfA57Ab96e": "0xa03bf2b145b0154c2e788a1d4642d235f6ff1c8aceeb41d0d7232525da8bdb77",
		//"0x73EB6d82CFB20bA669e9c178b718d770C49BB52f": "0xd09ba371c359f10f22ccda12fd26c598c7921bda3220c9942174562bc6a36fe8",
		//"0x9D8E5fAc117b15DaCED7C326Ae009dFE857621f1": "0x2d2719c6a828911ed0c50d5a6c637b63353e77cf57ea80b8e90e630c4687e9c5",

		"0x3F8d1eC00Cbb5CaEffa4c3cf620cA7aBa9D0a7ED": "0x98aed43383320bcd881ebb6008f52cab58fddf3bc7a08f9c92c4339872f9f99c",
		"0xffF91ec916d9eFFF8119A0ADE3358764f1F87e4E": "0x068c31f042082d826d6ecedb7c92beb5d0dacc1086f0a65a0a4ccf50d22e85a0",
		"0xD786E63F72c5270Ce7858021939aae0172054A1B": "0x5cd12959dae958f2a746f1fb26dd557fda39584c8590c6c3057438abcf33519d",
		"0x13E33800F84a1e8e52E305532f51f8095a34AC35": "0xa214f33247fdae837db1c4ecea82756997fd5046cf1da54c1f4efda1056c4a24",
		"0x3BB05828F4220F89e38754CC5C04994F61759078": "0xfa65a521da07edca2698b5e7e85410530b919e415f53e7d0ec9e17d1b173dc03",
		"0x1385F3CC51BFa743e2f8C79C8768a8a9CB7BEb05": "0x22484488b8159f072aa64673d605e4a149254d7cc48ceaac204f19ee9506075c",
		"0x3B84B0fC56A99734fc507159DF9669441C8e8065": "0xeffe12dcf8d0c84f841f90e4ec9824d7d791ce19ea0c38cc1a27d8439f6fc2fd",
		"0x53324d5e589A0facb6f7508AA4F3baA8B08a69Fb": "0xadab3f406ad1bc49ae79bb39ad3dff0c1c82a29e2792c926b49dfaef7ae4deb4",
		"0xD78125c0E1D0227f97D4e8DDcDeAFe920e705D67": "0x3aa8a1a2cbba5b7fd1726e433f60785944f55ab25b89defb96bf66e75ae056c8",
		"0x7366C479D8C1bEA452B67327255FEDB43B3ECc8E": "0xfec45ed1ff8a26bab000010d27baac2001c5e1d8fa27569c244bd4a7a1875bc7",

		"0x5dDEE435078ab3145f90954aE48267EA21554dDC": "0xfb25aac47fd501142dd5de16cabf2edbf91bd3e5ce4075760965dfe7b7a3efa0",
		"0x832EB4A25C22D8787db9fFbb5617b5A525Eb1Cb3": "0x0128a3ad251c704180d895dd9a8a32e2728b3a9ae9cd57f734db2459815e7b6c",
		"0x3a7a90aE229A17014C3d68cd7Eddf95582084D74": "0x72dce90bbe718017d6b26ac779556adbf569ea3623de960a325181279b1959ad",
		"0xD25db7172BaAf1F3027F02ffF460E468A677c51C": "0xdca667f7b3bcddb67b8b4e7b867bed8ce3e9e7a805e71d2713dc9b308151bff0",
		"0x27eC9fAf3FeE0cB1966471eb48C7b407A9244cf6": "0x1c287d267850f7b0a513a320d69bb30935a3d1f4e8871f266d8176a2014ccb51",
		"0x0e294f70252E87f116CC9bB31366CDdf0f510295": "0xaf58c69852219010077dd60ef874bb94b2582c7c48ead7ebabba2369e39cd2e8",
		"0x91A5418D951E7B19A1cEd84BF442d8d6B36D6132": "0x0d5b3dfb3584255c23322ee32d2d83c1e7337cb0af5204bf1488be981916d6d6",
		"0xAE94Dc22Eb575157A5AF01F8eA714b3Ff1750105": "0x6d91b57da06e6f5b93435e7df2a0a74a609914cecc01117fdd3e52cbaf12220b",
		"0x672aBbc8A016610Edf97A55b8cB8dA2e0c368643": "0x62fbedf1d2e0ae245fbb3c874141d3ed4524fd9ec4cf111dcc91f7fff7070000",
		"0x6Bc1C8a7A6e6c5e08350B308982AD8F15127A269": "0xdd1ed3c76baa4b8e0ba7fdf843dd321a83a5ee1aea4a249ea06ab16d5ca9d448",
	}
	for k, v := range keys {
		var eth Eth
		eth.PublicKey = k
		eth.PrivateKey = v
		_, _ = o.Insert(&eth)
	}
}
