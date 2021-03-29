package rsinit

import (
	"github.com/gfm/core/fo"
)

func InitData() {
	cfg := fo.LoadConf()
	cfg.Section("paths").Key("logpath").SetValue(fo.GetMainPath() + fo.GetTday() + "/" + fo.GetFileName())
	cfg.SaveTo(fo.ConfContent.ConfName)
}

func RunInit() {
	InitData()
}
