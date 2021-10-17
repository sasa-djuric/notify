package main

import (
	"fmt"
	"strconv"

	"github.com/gen2brain/beeep"
)

func help() {
	fmt.Println("--title 	-t	Notification title")
	fmt.Println("--description 	-d 	Notification description")
	fmt.Println("--subject 	-s 	Subject of notification")
	fmt.Println("--version 	-v 	Tool version")
}

func version() {
	fmt.Println("1.0.0")
}

func notification(subject string, title string, description string) {
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

func main() {
	parseArgs()

	isHelp, _ := strconv.ParseBool(getArg("help", "h", true))
	isVersion, _ := strconv.ParseBool(getArg("version", "v", true))
	subject := getArg("subject", "s", false)
	title := getArg("title", "t", false)
	description := getArg("description", "d", false)

	if isHelp {
		help()
	} else if isVersion {
		version()
	} else {
		notification(subject, title, description)
	}
}
