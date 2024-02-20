package main

import (
	"ImSdk/configs"
	"ImSdk/routers"
	"ImSdk/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()
	var c configs.Config
	conf.MustLoad(*configFile, &c)

	configs.Conf = c
	svc.Ctx = svc.NewServiceContext(c)

	router := routers.InitRouter()
	prot := fmt.Sprint(":", c.Port)
	router.Run(prot)
}
