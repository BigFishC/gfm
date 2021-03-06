package main

import (
	"fmt"
	"os"

	"github.com/gfm/core/rsinit"
	"github.com/gfm/utils/setting"
)

func main() {

	args := os.Args
	if args == nil || len(args) == 1 {
		setting.Help()
	} else if len(args) < 4 {
		switch {
		case args[1] == "version" || args[1] == "--version":
			fmt.Println("v0.1.0")
		case args[1] == "run" || args[1] == "--run":
			rsinit.RunInit()
			setting.Run("pro")
		case args[1] == "run" && args[2] == "--debug":
			rsinit.RunInit()
			setting.Run("debug")
		default:
			setting.Help()
		}
	} else {
		setting.Help()
	}
}
