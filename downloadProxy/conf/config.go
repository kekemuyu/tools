package config

import (
	"os"
	"path/filepath"

	"github.com/donnie4w/go-logger/logger"
	"github.com/go-ini/ini"
)

var Cfg *ini.File

func init() {
	var err error

	Cfg, err = ini.Load(GetRootdir() + "/conf/conf.ini")

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
		if Exist(d + "/conf") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	return infer(cwd)
}

func Save() {
	Cfg.SaveTo(GetRootdir() + "/conf/conf.ini")
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
