package main

import (
	"os"
	"strings"
)

func parseArgs() map[string]string {
	args := map[string]string{}

	for i, s := range os.Args {
		if strings.HasPrefix(s, "-") {
			args[s] = os.Args[i+1]
		}
	}

	return args
}

func getArg(args map[string]string, arg string, argAlias string) string {
	if val, ok := args["--"+arg]; ok {
		return val
	} else if aliasVal, ok := args["-"+argAlias]; ok {
		return aliasVal
	}

	return ""
}
