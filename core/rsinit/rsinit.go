package rsinit

import (
	"log"
	"os"
	"runtime"

	"github.com/gfm/core/fo"
)

func InitConfData() {
	cfg := fo.LoadConf()
	cfg.Section("").Key("filename").SetValue("error-" + fo.GetTday() + ".log")
	cfg.SaveTo(fo.ConfContent.ConfName)

	if runtime.GOOS == "linux" {
		cfg.Section("paths").Key("logpath").SetValue(fo.GetMainPath() + fo.GetTday() + "/" + fo.GetFileName())
		cfg.SaveTo(fo.ConfContent.ConfName)
	} else if runtime.GOOS == "windows" {
		cfg.Section("paths").Key("logpath").SetValue(fo.GetMainPath() + "\\" + "\\" + fo.GetTday() + "\\" + "\\" + fo.GetFileName())
		cfg.SaveTo(fo.ConfContent.ConfName)
	} else {
		log.Fatalln("OS 操作系统不可用！")
	}
	_, err := os.Stat(fo.GetLog(fo.ConfContent.ConfName))
	if err != nil {
		log.Fatalln(err)
	}
}

func RunInit() {
	InitConfData()
}
