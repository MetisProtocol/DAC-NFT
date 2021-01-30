package conf

import "github.com/astaxie/beego"

const RegisterEmail string = "var registerEmail string= \"Hello there,\\n\\nIn order to create an account on the blockchain with Metis, we need to verify your email address.\\n\\nPlease enter the verification code xxxxxx in the webpage to continue.\\n\\nBest Regards,\\nMetis Team\"\n"
const RegisterSuccessEmail string = "Hello there,\n\nCongratulations! And welcome to the decentralized world!\n\nYou have successfully created your account on the blockchain with Metis.\nLogin email address: \nAccount address: xxxxx(以太坊地址)\n\nWith Metis, you can create your exclusive and coolest decentralized companies on the blockchain, mint NFT, play games, and collaborate with your friends, or maybe someone else you have never met before. More fun will be unveiled! Please stay tuned with us! Thanks! \n\nBest Regards,\nMetis Team"
const PrivateKeyEmail string = "Hello there,\n\nThis email is sent to you as you are requesting the private key of your account on Metis.\n\nPlease be aware that the private key can be used to access your account. DO NOT share this key with anyone! And you are fully responsible for your account after you request your private key. Metis is not responsible for any potential loss of your account.\n\nYour private key is: xxxx\n\nPlease save it somewhere safe and secret, and for you use ONLY!\n\nBest Regards,\nMetis Team"

func GetDatabasePrefix() string {
	return beego.AppConfig.DefaultString("db_prefix", "metis_")
}
