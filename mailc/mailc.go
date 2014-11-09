package mailc

import (
	cli "github.com/mitchellh/cli"
	"fmt"
)

type DefaultCommand struct {
	UI   cli.Ui
}

func (d *DefaultCommand) Run(args []string) int {
	fmt.Println("Default command called")
	return 0
}

func (d *DefaultCommand) Synopsis() string {
	return "Launches mailc"
}

func (d *DefaultCommand) Help() string {
	help := `
No help yet.
`
	return help
}

type TestCommand struct {
	UI   cli.Ui
}

func (d *TestCommand) Run(args []string) int {
	fmt.Println("Test command called")
	return 0
}

func (d *TestCommand) Synopsis() string {
	return "Launches test"
}

func (d *TestCommand) Help() string {
	help := `
No help yet.
`
	return help
}
