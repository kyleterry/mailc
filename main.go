package main

import (
	clipkg "github.com/kyleterry/cli"
	"github.com/kyleterry/mailc/mailc"
	"os"
	"fmt"
)

func main() {
	cli := clipkg.NewCLI("mailc", Version)
	cli.Args = os.Args[1:]
	cli.Commands = map[string]clipkg.CommandFactory{
		"testing": func() (clipkg.Command, error) {
			return &mailc.TestCommand{
				UI: &clipkg.BasicUi{Writer: os.Stdout},
			}, nil
		},
	}
	cli.DefaultCommand = func() (clipkg.Command, error) {
		return &mailc.DefaultCommand{
			UI: &clipkg.BasicUi{Writer: os.Stdout},
		}, nil
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Println("fuck an error")
		os.Exit(exitCode)
	}
	os.Exit(exitCode)
}

