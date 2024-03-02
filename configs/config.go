package configs

var Conf Config

type Config struct {
	Name string
	Host string
	Port int64
	Mode string

	Token struct {
		Appid  string
		Secret string
		Expire int64
	}
	TestImUrl   string
	ImUrl       string
	CallBackUrl string
}
