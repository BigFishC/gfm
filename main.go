package main

import (
	"fmt"
	"os"

	"github.com/gfm/core/init"
	"github.com/gfm/core/init/init"
	"github.com/gfm/utils/setting"
)

func main() {

	args := os.Args
	if args == nil || len(args) == 1 {
		setting.Help()
		init.RunInit()
	} else if len(args) < 4 {
		switch {
		case args[1] == "version" || args[1] == "--version":
			fmt.Println("v0.1.0")
		case args[1] == "run" || args[1] == "--run":
			setting.Run("pro")
		case args[1] == "run" && args[2] == "--debug":
			setting.Run("debug")
		default:
			setting.Help()
		}
	} else {
		setting.Help()
	}
}
