package helpers

type User struct {
	UserName        string `form:"username"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
}

type Dac struct {
	DacName    string `form:"dac_name"`
	DacProduce string `form:"dac_produce"`
	DacLogo    string `form:"dac_logo"`
}

type Email struct {
	ToEmail string `form:"email"`
}

type Code struct {
	EmailCode string `form:"code"`
}
