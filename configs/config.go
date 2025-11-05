package configs

var Conf Config

//	type MySQL struct {
//		Host           string `yaml:"host"`
//		Port           string `yaml:"port"`
//		User           string `yaml:"user"`
//		Password       string `yaml:"password"`
//		Database       string `yaml:"database"`
//		InitDB         int    `yaml:"initDB"`
//		LogLevel       int    `yaml:"logLevel"`
//		SlowThreshold  int    `yaml:"slowThreshold"`
//		DBMaxOpenConns int    `yaml:"dbMaxOpenConns"`
//		DBMaxIdleConns int    `yaml:"dbMaxIdleConns"`
//		DBMaxLifeTime  int    `yaml:"dbMaxLifeTime"`
//	}
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
	MySQL struct {
		Host           string
		Port           string
		User           string
		Password       string
		Database       string
		InitDB         int
		LogLevel       int
		SlowThreshold  int
		DBMaxOpenConns int
		DBMaxIdleConns int
		DBMaxLifeTime  int
	}
	Teltlk struct {
		TestImUrl   string
		ImUrl       string
		CallBackUrl string
	}

	ClientLogReq struct {
		Path string
	}
}
