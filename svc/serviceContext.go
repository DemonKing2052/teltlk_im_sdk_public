package svc

import (
	"ImSdk/configs"
	"os"
	"path/filepath"
	"sync"
)

var Ctx *ServiceContext

type ServiceContext struct {
	Config       configs.Config
	ClientLogReq struct {
		FhMutex sync.Mutex
		Fh      *os.File
	}
}

func NewServiceContext(c configs.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		ClientLogReq: struct {
			FhMutex sync.Mutex
			Fh      *os.File
		}{
			Fh: ensureLogFile(c.ClientLogReq.Path),
		},
	}
}

func ensureLogFile(p string) *os.File {
	dir := filepath.Dir(p)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil
	}
	f, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil
	}
	return f
}
