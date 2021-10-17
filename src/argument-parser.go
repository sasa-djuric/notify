package main

import (
	"os"
	"strings"
)

var args map[string]string

func parseArgs() {
	args = map[string]string{}

	for i, s := range os.Args {
		if strings.HasPrefix(s, "-") {
			if i+1 >= len(os.Args) || strings.HasPrefix(os.Args[i+1], "-") {
				args[s] = "true"
			} else {
				args[s] = os.Args[i+1]
			}
		}
	}
}

func argExists(arg string, isFlag bool) bool {
	if value, exists := args[arg]; exists && ((isFlag && value == "true") || (!isFlag && value != "true")) {
		return true
	}

	return false
}

func getArg(arg string, argAlias string, isFlag bool) string {
	if argExists("--"+arg, isFlag) {
		return args["--"+arg]
	} else if argExists("-"+argAlias, isFlag) {
		return args["-"+argAlias]
	}

	return ""
}
