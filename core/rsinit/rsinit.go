package rsinit

import (
	"log"

	"github.com/gfm/core/fo"
	"gopkg.in/ini.v1"
)

func InitData() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatalln(err)
	}
	cfg.Section("paths").Key("logpath").SetValue(fo.GetMainPath() + fo.GetTday() + "/" + fo.GetFileName())
	cfg.SaveTo("conf.ini")
}

func RunInit() {
	InitData()
}
