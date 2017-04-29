package utils

import (
	commonpallete "github.com/beldpro-ci/common/pallete"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func AssertIntFlagsSet(command string, c *cli.Context, flags ...string) {
	assertCommandSet(command)
	for _, flag := range flags {
		if c.Int(flag) == 0 && c.GlobalInt(flag) == 0 {
			printUsage(c, command)
			printErrorMessage(flag)
			os.Exit(1)
		}
	}
}

func AssertStringFlagsSet(command string, c *cli.Context, flags ...string) {
	assertCommandSet(command)
	for _, flag := range flags {
		if c.String(flag) == "" && c.GlobalString(flag) == "" {
			printUsage(c, command)
			printErrorMessage(flag)
			os.Exit(1)
		}
	}
}

func AssertStringSliceFlagsSet(command string, c *cli.Context, flags ...string) {
	assertCommandSet(command)
	for _, flag := range flags {
		if len(c.StringSlice(flag)) == 0 && len(c.GlobalStringSlice(flag)) == 0 {
			printUsage(c, command)
			printErrorMessage(flag)
			os.Exit(1)
		}
	}
}

func assertCommandSet(command string) {
	if command == "" {
		panic("A command must be specified.")
	}
}

func printUsage(c *cli.Context, command string) {
	if command == "_app" {
		cli.ShowAppHelp(c)
	} else {
		cli.ShowCommandHelp(c, command)
	}
}

func printErrorMessage(flag string) {
	commonpallete.BoldRed.Print("    Error:")
	commonpallete.Red.Print("\n")
	commonpallete.Red.Print("        Required argument ")
	commonpallete.BoldRed.Printf("%s", flag)
	commonpallete.Red.Print(" not set.")
	commonpallete.Red.Print("\n\n")
	commonpallete.Red.Print("        Aborting.")
	commonpallete.Red.Print("\n\n")
}
