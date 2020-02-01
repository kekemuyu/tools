package config

import (
	"os"
	"path/filepath"

	"gofile/util"

	"github.com/donnie4w/go-logger/logger"
	"github.com/go-ini/ini"
)

var Cfg *ini.File

func init() {
	var err error

	Cfg, err = ini.Load(GetRootdir() + "/config/conf.ini")

	if err != nil {
		logger.Error("Fail to read file: %v", err)
		os.Exit(1)
	}
}

func GetRootdir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		if util.Exist(d + "/config") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	return infer(cwd)
}

func Save() {
	Cfg.SaveTo(GetRootdir() + "/config/conf.ini")
}
