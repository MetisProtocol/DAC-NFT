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

		//"0x3F8d1eC00Cbb5CaEffa4c3cf620cA7aBa9D0a7ED": "0x98aed43383320bcd881ebb6008f52cab58fddf3bc7a08f9c92c4339872f9f99c",
		//"0xffF91ec916d9eFFF8119A0ADE3358764f1F87e4E": "0x068c31f042082d826d6ecedb7c92beb5d0dacc1086f0a65a0a4ccf50d22e85a0",
		//"0xD786E63F72c5270Ce7858021939aae0172054A1B": "0x5cd12959dae958f2a746f1fb26dd557fda39584c8590c6c3057438abcf33519d",
		//"0x13E33800F84a1e8e52E305532f51f8095a34AC35": "0xa214f33247fdae837db1c4ecea82756997fd5046cf1da54c1f4efda1056c4a24",
		//"0x3BB05828F4220F89e38754CC5C04994F61759078": "0xfa65a521da07edca2698b5e7e85410530b919e415f53e7d0ec9e17d1b173dc03",
		//"0x1385F3CC51BFa743e2f8C79C8768a8a9CB7BEb05": "0x22484488b8159f072aa64673d605e4a149254d7cc48ceaac204f19ee9506075c",
		//"0x3B84B0fC56A99734fc507159DF9669441C8e8065": "0xeffe12dcf8d0c84f841f90e4ec9824d7d791ce19ea0c38cc1a27d8439f6fc2fd",
		//"0x53324d5e589A0facb6f7508AA4F3baA8B08a69Fb": "0xadab3f406ad1bc49ae79bb39ad3dff0c1c82a29e2792c926b49dfaef7ae4deb4",
		//"0xD78125c0E1D0227f97D4e8DDcDeAFe920e705D67": "0x3aa8a1a2cbba5b7fd1726e433f60785944f55ab25b89defb96bf66e75ae056c8",
		//"0x7366C479D8C1bEA452B67327255FEDB43B3ECc8E": "0xfec45ed1ff8a26bab000010d27baac2001c5e1d8fa27569c244bd4a7a1875bc7",
		//
		//"0x5dDEE435078ab3145f90954aE48267EA21554dDC": "0xfb25aac47fd501142dd5de16cabf2edbf91bd3e5ce4075760965dfe7b7a3efa0",
		//"0x832EB4A25C22D8787db9fFbb5617b5A525Eb1Cb3": "0x0128a3ad251c704180d895dd9a8a32e2728b3a9ae9cd57f734db2459815e7b6c",
		//"0x3a7a90aE229A17014C3d68cd7Eddf95582084D74": "0x72dce90bbe718017d6b26ac779556adbf569ea3623de960a325181279b1959ad",
		//"0xD25db7172BaAf1F3027F02ffF460E468A677c51C": "0xdca667f7b3bcddb67b8b4e7b867bed8ce3e9e7a805e71d2713dc9b308151bff0",
		//"0x27eC9fAf3FeE0cB1966471eb48C7b407A9244cf6": "0x1c287d267850f7b0a513a320d69bb30935a3d1f4e8871f266d8176a2014ccb51",
		//"0x0e294f70252E87f116CC9bB31366CDdf0f510295": "0xaf58c69852219010077dd60ef874bb94b2582c7c48ead7ebabba2369e39cd2e8",
		//"0x91A5418D951E7B19A1cEd84BF442d8d6B36D6132": "0x0d5b3dfb3584255c23322ee32d2d83c1e7337cb0af5204bf1488be981916d6d6",
		//"0xAE94Dc22Eb575157A5AF01F8eA714b3Ff1750105": "0x6d91b57da06e6f5b93435e7df2a0a74a609914cecc01117fdd3e52cbaf12220b",
		//"0x672aBbc8A016610Edf97A55b8cB8dA2e0c368643": "0x62fbedf1d2e0ae245fbb3c874141d3ed4524fd9ec4cf111dcc91f7fff7070000",
		//"0x6Bc1C8a7A6e6c5e08350B308982AD8F15127A269": "0xdd1ed3c76baa4b8e0ba7fdf843dd321a83a5ee1aea4a249ea06ab16d5ca9d448",

		"0xF005a988DD1Ad01D5d771E3e9e9cA3907D2D5a05": "0xaee380352addaf27d6213a378c1dd2d78d17b461356bc6eb461d9ac6cad1a0b6",
		"0x15a0eb8C6a23ae474753a6d5179f8016f131Ce36": "0x0912100f244d404c94b287c9ec7e9115f2d098f1f96e4c3d6cc1cad513c7ed0a",
		"0x1d8e88447E846F4fDe2e74e7a0B9F0B7Ceb2F668": "0xf4889ef6ba96de8c4065080df17fc774d86ff044eb73598f207daece72972b37",
		"0xda7399bFb60A618b08C5B4b9a28673e0C10FaA6C": "0x43aa360ce239f03c35674cf34fa134cb48fd5b9d5776c890f3e3e3e284e42905",
		"0xBfFfCe9C485F478F29bd0742432e07B1972B7A76": "0x9539ddc992fb60191748c967e6c77ec06b2764f40c25c9fae963ec3c2f3a2d18",
		"0xa47B83Db0fE9809A5eBE1d8478F8E37DAB7DcB66": "0xf33bf0104f32f396c51e1be1a1c4a3925b10018365c759471d94b9d5c80441fc",
		"0x43831dE8316a4A05D3510e9a8739Bca671758F20": "0x222a00465826323e1aa31c9744886fde03038e238e06a77e3570c64b6e5e3f72",
		"0x4afdd76b869F9a18cDD0d970B06D4d53D39DDd1f": "0x216df49bea781ae2c3f08fd10b6339795c0988fe107ecd6b1a043b80ac11f03e",
		"0xC54d2e8aD5272d62608d9e993B1762Da90B05657": "0x2dfade394d26c6168e2f0378127aaede6e0e7961ae99ae93a3018aa7509b8c54",
		"0xF23731CD5abD0D0651CF648AB3AA4337727b08F9": "0x92aa06fed08c1f6fe45ea5a5219669fc3470deb788b9d2fd2931659abfcb6729",

		"0xE20aCDC2329f5d887Ab4E7740aeC29F887B7269F": "0xa75da29b7e06553379b94da3a2410a2b946bdf0149bc14de4894e78ef3a2b8ba",
		"0x874d6302693F79dd312017F2Eaa3b37a0625ac66": "0x0ed234e6fdad1c503c26cd6e8639ea9235acc8fd46aa95576f1029518d44ea7d",
		"0xBb08D5A8D4C966877DAb5fB21511D293C4237D29": "0x7ba06083abfada32c234e17a006cd832a3ecc016a2a400b2123a5a64561e9560",
		"0x4C306F81b925a28f1931098d7Fdfe2d278D933fB": "0xddbf5c096b6c017e249a93e1d79f44292cb3c24ae7d7a5362643258b1e646a98",
		"0x6Bb87a5A205f4edee2F0237De5fc8D5a55Cf956e": "0x28af39fa45edd14e97d96fc1ee4ef16f7f02023e2144f80fbf00bab7744039e8",
		"0xB5D9b8690139b49C14106d05143Ab4aD90DB9c5c": "0x34f025c3917b36e81b5442e39becd003001c8116d64bcdf56f9dde86753bb191",
		"0x282f3C69C5D7d4F3e9A2BE5c5268FAf206315E9D": "0x7504ecf463390cc4fc7e69d765740d813efcab9af934e46d58498aaa02329f7f",
		"0x084B47F82372D7c911f3a6f78eFC6E7F700a1449": "0x71dcd66e0134423db377ab264cfc03ed3052261024fb5b5cd840570644b2075a",
		"0x9322D94654e5C44FFd983342Dff4E04f38708e85": "0x9d8fef21158b4c6a7c5c7dda9982f5535944e3d73018fa52631dc48c0557b920",
		"0x8Dc373769e0f781e5949f0fbAdc1F658746E9309": "0xe2090e16cb0d554ed2d2306e532e94f62c83f91699c4fed26af7b0713fcf44a5",

		"0xca8Dd91Ec24c60239fd258c7BC49fACab6e85294": "0x9518a9c72a062686fa17dfd89e6a2b68d0bd24a0d53913a99d210019570d934c",
		"0x8EE9f46Ed8e4309661B88A9e8540553169149Dc0": "0x7f8b2d45f7cc818f321ee7ca4167de92253ef7a315f01e0b8f9edefbc022218a",
		"0x6dbB32E49700959B48fF86E5311af6491f13c245": "0x27126dfd696b6a50e0cfddf421a7123e3ef795cf3f6341909f8b6ab1d4c97c93",
		"0x2b8bEe675b189e6cb0377a61BFfaf3f7cc5A0EFB": "0xb66e3e7b374d1ac74d4490ddb2d1d9857d0209baede6fe7705be8ca27001bf14",
		"0x3724e93A2b400ba5f563aCda6Cb66604e14061A5": "0x4848eab49b6dedf8513ed0008fdb97f18297ec611d52c9038bdfe8f29a90fdc0",
		"0x96CA43209d61A948d36e6c3375Eff3575E6f9Dcf": "0x8f0ebe10a49323b9c57a3601621ebd6d6793b5c19b4abfab42164da721fed12d",
		"0xD0A200711B58D16547F73143a71B2b26F5dAd834": "0x236d8e3bb420fe488d43674ed8f9e5ac6f41fcd4c6538c06b5d5051d88d6eb5b",
		"0x41282f74180b2A2563Dc5E33627dE43f0a7c295D": "0x4247d88a8b23d882f0352856e68b3ed6ac1c0c2451f7bdcc61ac3400b6d155da",
		"0x2f8F1D68d805F6173b8D68F917Ca2726A0E89b57": "0x676051a6ed58fe85f60591707edac38a12b0f1e8367af32ab54b95add00b36c3",
		"0x865f177c6eF159A378337c11744BB50d7FC44D53": "0x38cc2458b84626e5b46e0872e403a038a8c4a0f495ee067d214dec7ac7ff255e",
	}
	for k, v := range keys {
		var eth Eth
		eth.PublicKey = k
		eth.PrivateKey = v
		_, _ = o.Insert(&eth)
	}
}
