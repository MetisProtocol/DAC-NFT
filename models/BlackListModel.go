package models

import (
	"github.com/astaxie/beego/orm"
	"metis-v1.0/conf"
	"time"
)

type BlackList struct {
	Id        int       `orm:"pk;auto;"`
	BlackName string    `orm:"size(50);unique"`
	Status    bool      `default:"1"`
	Ctime     time.Time `orm:"auto_now_add;type(datetime)"`
	Utime     time.Time `orm:"auto_now;type(datetime)"`
}

func (m *BlackList) TableName() string {
	return "black_list"
}

func (m *BlackList) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

func NewBlackList() *BlackList {
	return &BlackList{}
}

func (m *BlackList) InsertBlackList() {
	o := orm.NewOrm()
	blackList := [...]string{
		"Apple",
		"BerkshireHathaway",
		"Amazon",
		"UnitedHealthGroup",
		"McKesson",
		"CVSHealth",
		"AT&T",
		"AmerisourceBergen",
		"Chevron",
		"FordMotor",
		"GeneralMotors",
		"CostcoWholesale",
		"Alphabet",
		"CardinalHealth",
		"WalgreensBootsAlliance",
		"JPMorganChase",
		"VerizonCommunications",
		"Kroger",
		"GeneralElectric",
		"FannieMae",
		"Phillips66",
		"ValeroEnergy",
		"BankofAmerica",
		"Microsoft",
		"HomeDepot",
		"Boeing",
		"WellsFargo",
		"Citigroup",
		"MarathonPetroleum",
		"Comcast",
		"Anthem",
		"DellTechnologies",
		"DowDuPont",
		"StateFarmInsurance",
		"Johnson&Johnson",
		"IBM",
		"Target",
		"FreddieMac",
		"UnitedParcelService",
		"Lowes",
		"Intel",
		"MetLife",
		"Procter&Gamble",
		"UnitedTechnologies",
		"FedEx",
		"PepsiCo",
		"ArcherDanielsMidland",
		"PrudentialFinancial",
		"Centene",
		"Albertsons",
		"WaltDisney",
		"Sysco",
		"HP",
		"Humana",
		"Facebook",
		"Caterpillar",
		"EnergyTransfer",
		"LockheedMartin",
		"Pfizer",
		"GoldmanSachsGroup",
		"MorganStanley",
		"CiscoSystems",
		"Cigna",
		"AIG",
		"HCAHealthcare",
		"AmericanAirlinesGroup",
		"DeltaAirLines",
		"CharterCommunications",
		"NewYorkLifeInsurance",
		"AmericanExpress",
		"Nationwide",
		"BestBuy",
		"LibertyMutualInsuranceGroup",
		"Merck",
		"HoneywellInternational",
		"UnitedContinentalHoldings",
		"TIAA",
		"TysonFoods",
		"Oracle",
		"Allstate",
		"WorldFuelServices",
		"MassachusettsMutualLifeInsurance",
		"TJX",
		"ConocoPhillips",
		"Deere",
		"TechData",
		"EnterpriseProductsPartners",
		"Nike",
		"PublixSuperMarkets",
		"GeneralDynamics",
		"Exelon",
		"PlainsGPHoldings",
		"3M",
		"AbbVie",
		"CHS",
		"CapitalOneFinancial",
		"Progressive",
		"CocaCola",
		"USAA",
		"HewlettPackardEnterprise",
		"AbbottLaboratories",
		"TwentyFirstCenturyFox",
		"MicronTechnology",
		"Travelers",
		"RiteAid",
		"NorthropGrumman",
		"ArrowElectronics",
		"PhilipMorrisInternational",
		"NorthwesternMutual",
		"INTLFCStone",
		"PBFEnergy",
		"Raytheon",
		"KraftHeinz",
		"MondelezInternational",
		"USBancorp",
		"Macys",
		"DollarGeneral",
		"Nucor",
		"Starbucks",
		"DXCTechnology",
		"EliLilly",
		"ThermoFisherScientific",
		"USFoodsHolding",
		"DukeEnergy",
		"Halliburton",
		"Cummins",
		"Amgen",
		"Paccar",
		"Southern",
		"CenturyLink",
		"InternationalPaper",
		"UnionPacific",
		"DollarTree",
		"PenskeAutomotiveGroup",
		"Qualcomm",
		"BristolMyersSquibb",
		"GileadSciences",
		"Jabil",
		"ManpowerGroup",
		"SouthwestAirlines",
		"Aflac",
		"Tesla",
		"AutoNation",
		"CBREGroup",
		"Lear",
		"Whirlpool",
		"McDonalds",
		"Broadcom",
		"MarriottInternational",
		"WesternDigital",
		"Visa",
		"Lennar",
		"WellCareHealthPlans",
		"Kohls",
		"AECOM",
		"Synnex",
		"PNCFinancialServices",
		"Danaher",
		"HartfordFinancialServices",
		"AltriaGroup",
		"BankofNewYorkMellon",
		"Fluor",
		"Avnet",
		"IcahnEnterprises",
		"OccidentalPetroleum",
		"MolinaHealthcare",
		"GenuineParts",
		"FreeportMcMoRan",
		"KimberlyClark",
		"TenetHealthcare",
		"SynchronyFinancial",
		"CarMax",
		"HollyFrontier",
		"PerformanceFoodGroup",
		"SherwinWilliams",
		"EmersonElectric",
		"NGLEnergyPartners",
		"XPOLogistics",
		"EOGResources",
		"AppliedMaterials",
		"PG&E",
		"NextEraEnergy",
		"CHRobinsonWorldwide",
		"Gap",
		"LincolnNational",
		"DaVita",
		"JonesLangLaSalle",
		"WestRock",
		"CDW",
		"AmericanElectricPower",
		"CognizantTechnologySolutions",
		"DRHorton",
		"BectonDickinson",
		"Nordstrom",
		"Netflix",
		"Aramark",
		"TexasInstruments",
		"GeneralMills",
		"Supervalu",
		"ColgatePalmolive",
		"GoodyearTire&Rubber",
		"PayPalHoldings",
		"PPGIndustries",
		"OmnicomGroup",
		"Celgene",
		"JacobsEngineeringGroup",
		"RossStores",
		"Marsh&McLennan",
		"Mastercard",
		"LandOLakes",
		"WasteManagement",
		"IllinoisToolWorks",
		"Ecolab",
		"BookingHoldings",
		"CBS",
		"ParkerHannifin",
		"PrincipalFinancial",
		"DTEEnergy",
		"BlackRock",
		"UnitedStatesSteel",
		"CommunityHealthSystems",
		"KinderMorgan",
		"QurateRetail",
		"Loews",
		"Arconic",
		"StanleyBlack&Decker",
		"Textron",
		"LasVegasSands",
		"EsteeLauder",
		"DISHNetwork",
		"Stryker",
		"Kellogg",
		"Biogen",
		"Alcoa",
		"AnadarkoPetroleum",
		"DominionEnergy",
		"ADP",
		"salesforce",
		"LBrands",
		"HenrySchein",
		"NewellBrands",
		"GuardianLifeInsCoofAmerica",
		"BJsWholesaleClub",
		"BB&TCorp",
		"StateStreetCorp",
		"Viacom",
		"AmeripriseFinancial",
		"CoreMarkHolding",
		"XiJinping",
		"HuJingtao",
		"JiangZemin",
		"DengXiaoping",
		"MaoZedong",
		"ZhaoZiyang",
		"Lipeng",
		"Wenjiabao",
		"WangQishan",
		"ZhouEnlai",
		"LiKeqiang",
		"GongChanDang",
		"abo",
		"Argie",
		"bitch",
		"cunt",
		"Ching Chong",
		"Chink",
		"chinaman",
		"cocksucker",
		"coolie",
		"cunt",
		"dyke",
		"dago",
		"Diegodarkie",
		"fuck",
		"gingo",
		"honky",
		"honkie",
		"Jap",
		"Japie",
		"kike",
		"Kafir",
		"kiki",
		"kraut",
		"boche",
		"sauerkraut",
		"mick",
		"Michealmongol",
		"Mongoloid",
		"Nword",
		"nigger",
		"nignog",
		"Paddy",
		"Polack",
		"Polak",
		"popish",
		"queer",
		"Sambo",
		"spick",
		"spic",
		"spik",
		"slant",
		"twat",
		"wetback",
		"wop",
		"yid",
		"Yank",
		"Yankee",
		"Paki",
		"honky honkie",
		"halfbreed",
		"halfcaste",
		"Hun",
		"Heeb",
		"Hebe",
		"queer",
		"fag",
		"faggot",
		"Redneck",
		"limey",
		"Letterbox",
		"Chinky eyes",
		"Pingpong",
		"beaner",
	}

	for _, v := range blackList {
		var black BlackList
		black.BlackName = v
		_, _ = o.Insert(&black)
	}
}

func (m *BlackList) GetBlackList(dacName string) bool {
	var black BlackList
	var registered bool
	o := orm.NewOrm()
	err := o.QueryTable("metis_black_list").Filter("BlackName__iexact", dacName).One(&black)
	if err == orm.ErrNoRows {
		registered = true
	} else {
		registered = false
	}
	return registered
}
