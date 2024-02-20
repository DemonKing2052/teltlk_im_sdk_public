package svc

import (
	"ImSdk/configs"
)

var Ctx *ServiceContext

type ServiceContext struct {
	Config configs.Config
}

func NewServiceContext(c configs.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}
