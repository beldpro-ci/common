package utils

import (
	log "github.com/Sirupsen/logrus"
	commonpallete "github.com/beldpro-ci/common/pallete"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func AssertIntFlagsSet(command string, c *cli.Context, logger *log.Entry, flags ...string) {
	assertCommandSet(command)
	for _, flag := range flags {
		if c.Int(flag) != 0 {
			logger.
				WithField("command", command).
				WithField(flag, c.Int(flag)).
				Debug("Local Flag set")
			continue
		}

		if c.GlobalInt(flag) == 0 {
			logger.
				WithField("command", command).
				WithField(flag, c.GlobalInt(flag)).
				Debug("Global Flag set")
			continue
		}

		logger.
			WithField("command", command).
			WithField("flag", flag).
			Warn("Flag not set")

		printUsage(c, command)
		printErrorMessage(flag)
		os.Exit(1)
	}
}

func AssertInt64FlagsSet(command string, c *cli.Context, logger *log.Entry, flags ...string) {
	assertCommandSet(command)
	for _, flag := range flags {
		if c.Int64(flag) != 0 {
			logger.
				WithField("command", command).
				WithField(flag, c.Int64(flag)).
				Debug("Local Flag set")
			continue
		}

		if c.GlobalInt64(flag) == 0 {
			logger.
				WithField("command", command).
				WithField(flag, c.GlobalInt64(flag)).
				Debug("Global Flag set")
			continue
		}

		logger.
			WithField("command", command).
			WithField("flag", flag).
			Warn("Flag not set")

		printUsage(c, command)
		printErrorMessage(flag)
		os.Exit(1)
	}
}

func AssertStringFlagsSet(command string, c *cli.Context, logger *log.Entry, flags ...string) {
	assertCommandSet(command)
	for _, flag := range flags {
		if len(c.String(flag)) != 0 {
			logger.
				WithField("command", command).
				WithField(flag, c.String(flag)).
				Debug("Local Flag set")
			continue
		}

		if len(c.GlobalString(flag)) != 0 {
			logger.
				WithField("command", command).
				WithField(flag, c.GlobalString(flag)).
				Debug("Global Flag set")
			continue
		}

		logger.
			WithField("command", command).
			WithField("flag", flag).
			Warn("Flag not set")

		printUsage(c, command)
		printErrorMessage(flag)
		os.Exit(1)
	}
}

func AssertStringSliceFlagsSet(command string, c *cli.Context, logger *log.Entry, flags ...string) {
	assertCommandSet(command)
	for _, flag := range flags {
		if len(c.StringSlice(flag)) != 0 {
			logger.
				WithField("command", command).
				WithField(flag, c.String(flag)).
				Debug("Local Flag set")
			continue
		}

		if len(c.GlobalStringSlice(flag)) != 0 {
			logger.
				WithField("command", command).
				WithField(flag, c.GlobalString(flag)).
				Debug("Global Flag set")
			continue
		}

		logger.
			WithField("command", command).
			WithField("flag", flag).
			Warn("Flag not set")

		printUsage(c, command)
		printErrorMessage(flag)
		os.Exit(1)
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
