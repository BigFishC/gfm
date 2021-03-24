package main

import (
	"fmt"
	"os"

	"github.com/gfm/utils/setting"
)

func main() {

	args := os.Args
	if args == nil || len(args) < 2 {
		setting.Help()
	} else {
		if args[1] == "help" || args[1] == "--help" {
			setting.Help()
		} else if args[1] == "version" || args[1] == "--version" {
			fmt.Println("v0.1.0")
		} else if args[1] == "run" || args[1] == "--run" {
			setting.Run()
		} else {
			setting.Help()
		}
	}
}
