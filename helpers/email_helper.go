package helpers

import "fmt"

func RegisterEmail(code string) string {
	return fmt.Sprintf("Hello there,<br><br>In order to create an account on the blockchain with Metis, we need to verify your email address.<br><br>Please enter the verification code %s in the webpage to continue.<br><br>Best Regards,<br>Metis Team", code)
}

func RegisterSuccessEmail(dacAddr string) string {
	return fmt.Sprintf("Hello there,<br><br>Congratulations! And welcome to the decentralized world!<br><br>You have successfully created your account on the blockchain with Metis.<br>Login email address: <br>Account address: %s <br><br>With Metis, you can create your exclusive and coolest decentralized companies on the blockchain, mint NFT, play games, and collaborate with your friends, or maybe someone else you have never met before. More fun will be unveiled! Please stay tuned with us! Thanks! <br><br>Best Regards,<br>Metis Team", dacAddr)
}

func RequestPrivateKey(privateKey string) string {
	return fmt.Sprintf("Hello there,<br><br>This email is sent to you as you are requesting the private key of your account on Metis.<br><br>Please be aware that the private key can be used to access your account. DO NOT share this key with anyone! And you are fully responsible for your account after you request your private key. Metis is not responsible for any potential loss of your account.<br><br>Your private key is: %s <br><br>Please save it somewhere safe and secret, and for you use ONLY!<br><br>Best Regards,<br>Metis Team", privateKey)
}
