package fo

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var ConfContent = struct {
	ConfName string `json:"confname"`
}{ConfName: "conf.ini"}

func LoadConf() *ini.File {
	cfg, err := ini.Load(ConfContent.ConfName)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}

func GetTday() string {
	td := time.Now().Format("2006-01-02")
	return td
}

func GetMainPath() string {
	cfg := LoadConf()
	lmp := cfg.Section("").Key("logmainpath").String()
	return lmp
}

func GetFileName() string {
	cfg := LoadConf()
	fn := cfg.Section("").Key("filename").String()
	return fn
}

func GetApi(conf string) string {
	cfg := LoadConf()
	ddapi := cfg.Section("apis").Key("ddapi").String()
	return ddapi
}

func GetLog(conf string) string {
	cfg := LoadConf()
	logpath := cfg.Section("paths").Key("logpath").String()
	return logpath
}

func GetMsg(conf string) string {
	cfg := LoadConf()
	reg_str := cfg.Section("data").Key("reg_str").String()
	return reg_str
}
