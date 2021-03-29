package rsinit

import (
	"log"
	"os"

	"github.com/gfm/core/fo"
)

func InitConfData() {
	cfg := fo.LoadConf()
	cfg.Section("").Key("filename").SetValue("error-" + fo.GetTday() + ".log")
	cfg.SaveTo(fo.ConfContent.ConfName)
	cfg.Section("paths").Key("logpath").SetValue(fo.GetMainPath() + fo.GetTday() + "/" + fo.GetFileName())
	cfg.SaveTo(fo.ConfContent.ConfName)
	_, err := os.Stat(fo.GetLog(fo.ConfContent.ConfName))
	if err != nil {
		log.Fatalln(err)
	}
}

func RunInit() {
	InitConfData()
}
