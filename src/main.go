package main

import (
	"github.com/gen2brain/beeep"
)

func main() {
	args := parseArgs()
	subject := getArg(args, "subject", "s")
	title := getArg(args, "title", "t")
	description := getArg(args, "description", "d")

	if subject == "" {
		subject = "Process"
	}

	if title == "" {
		title = subject + " finished"
	}

	NotifyErr := beeep.Notify(title, description, "")

	if NotifyErr != nil {
		return
	}
}
