package main

import (
	"fmt"
	"os"

	"github.com/gfm/utils/setting"
)

func main() {

	args := os.Args
	if args == nil || len(args) <= 2 {
		setting.Help()
	} else if len(args) == 1 {
		switch {
		case args[1] == "version" || args[1] == "--version":
			fmt.Println("v0.1.0")
		case args[1] == "run" || args[1] == "--run":
			setting.Run("pro")
		default:
			setting.Help()
		}
	} else {
		switch {
		case args[1] == "run" && args[2] == "--debug":
			setting.Run("debug")
		case args[1] == "run" && args[2] == "-d":
			setting.Run("debug")
		default:
			setting.Help()
		}
	}
}
