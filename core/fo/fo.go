package fo

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

func GetTday() string {
	td := time.Now().Format("2006-01-02")
	return td
}

func GetMainPath() string {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatalln(err)
	}
	lmp := cfg.Section("").Key("logmainpath").String()
	return lmp
}

func GetFileName() string {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatalln(err)
	}
	fn := cfg.Section("").Key("filename").String()
	return fn
}

func GetApi(conf string) string {
	cfg, err := ini.Load(conf)
	if err != nil {
		log.Fatalln(err)
	}
	ddapi := cfg.Section("apis").Key("ddapi").String()
	return ddapi
}

func GetLog(conf string) string {
	cfg, err := ini.Load(conf)
	if err != nil {
		log.Fatalln(err)
	}
	logpath := cfg.Section("paths").Key("logpath").String()
	return logpath
}

func GetMsg(conf string) string {
	cfg, err := ini.Load(conf)
	if err != nil {
		log.Fatalln(err)
	}
	reg_str := cfg.Section("data").Key("reg_str").String()
	return reg_str
}
