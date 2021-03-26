package init

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type Env struct {
	Foo string `ini:"foo"`
}

func InitConfig() {
	cfg, err := ini.Load([]byte("[env]\nfoo = ${test}\n"))
	if err != nil {
		log.Fatalln(err)
	}
	cfg.ValueMapper = os.ExpandEnv
	// ...
	env := &Env{}
	err = cfg.Section("env").MapTo(env)
	log.Fatalln(err)
}

func RunInit() {
	InitConfig()
}
