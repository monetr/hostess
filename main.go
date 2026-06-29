// hostess is command-line utility for managing your /etc/hosts file. Works on
// Unixes and Windows.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/monetr/hostess/hostess"
)

const help = `An idempotent tool for managing %s

Commands

    fmt                  Reformat the hosts file

    add <hostname> <ip>  Add or overwrite a hosts entry
    rm <hostname>        Remote a hosts entry
    on <hostname>        Enable a hosts entry
    off <hostname>       Disable a hosts entry

    ls                   List hosts entries
    has                  Exit 0 if entry present in hosts file, 1 if not

    dump                 Export hosts entries as JSON
    apply                Import hosts entries from JSON

    All commands that change the hosts file will implicitly reformat it.

Flags

    -n will preview changes but not rewrite your hosts file

Configuration

    HOSTESS_FMT may be set to unix or windows to force that platform's syntax
    HOSTESS_PATH may be set to point to a file other than the platform default

About

    Copyright 2015-2020 Chris Bednarski <chris@monetr.com>; MIT Licensed
    Portions Copyright the Go authors, licensed under BSD-style license
    Bugs and updates via https://github.com/monetr/hostess
`

var (
	Version           = "dev"
	ErrInvalidCommand = errors.New("invalid command")
)

func ExitWithError(err error) {
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
}

// printUsage writes the help text but does NOT exit. wrappedMain calls this so
// the no-command path can just return nil and stay testable; main() turns a nil
// error into a clean exit 0 anyway, so the observable behavior is the same.
func printUsage() {
	// help has a %s placeholder for the platform's hosts path, so this has to be
	// Printf and not Print otherwise we print the literal "%s" and tack the path
	// on the end. go vet catches it if someone swaps it back.
	fmt.Printf(help, hostess.GetHostsPath())
}

// Usage is the flag.Usage callback. flag invokes it on a parse error so it still
// exits, but wrappedMain itself no longer calls os.Exit in the middle of a run.
func Usage() {
	printUsage()
	os.Exit(0)
}

func CommandUsage(command string) error {
	return fmt.Errorf("Usage: %s %s <hostname>", os.Args[0], command)
}

func wrappedMain(args []string) error {
	cli := flag.NewFlagSet(args[0], flag.ExitOnError)
	preview := cli.Bool("n", false, "preview")
	cli.Usage = Usage

	command := ""
	if len(args) > 1 {
		command = args[1]
	}

	// args[2:] panics when hostess is run with no command at all (the original
	// reason Usage() exited here), so only slice off the command when there
	// actually is one. An empty command falls through to the help case below.
	rest := []string{}
	if len(args) > 2 {
		rest = args[2:]
	}

	if err := cli.Parse(rest); err != nil {
		return err
	}

	options := &Options{
		Preview: *preview,
	}

	switch command {

	case "-v", "--version", "version":
		fmt.Println(Version)
		return nil

	case "", "-h", "--help", "help":
		// Print help and return nil instead of exiting. main() exits 0 for a nil
		// error so `hostess` with no args still prints help and exits cleanly,
		// but now this path can run under a test without killing the process.
		printUsage()
		return nil

	case "fmt":
		return Format(options)

	case "add":
		if len(cli.Args()) != 2 {
			return fmt.Errorf("Usage: %s add <hostname> <ip>", cli.Name())
		}
		return Add(options, cli.Arg(0), cli.Arg(1))

	case "rm":
		if cli.Arg(0) == "" {
			return CommandUsage(command)
		}
		return Remove(options, cli.Arg(0))

	case "on":
		if cli.Arg(0) == "" {
			return CommandUsage(command)
		}
		return Enable(options, cli.Arg(0))

	case "off":
		if cli.Arg(0) == "" {
			return CommandUsage(command)
		}
		return Disable(options, cli.Arg(0))

	case "ls":
		return List(options)

	case "has":
		if cli.Arg(0) == "" {
			return CommandUsage(command)
		}
		return Has(options, cli.Arg(0))

	case "dump":
		return Dump(options)

	case "apply":
		if cli.Arg(0) == "" {
			return fmt.Errorf("Usage: %s apply <filename>", args[0])
		}
		return Apply(options, cli.Arg(0))

	default:
		return ErrInvalidCommand
	}
}

func main() {
	ExitWithError(wrappedMain(os.Args))
}
