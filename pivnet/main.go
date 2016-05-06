package main

import (
	"github.com/bunniesandbeatings/pivnet-cli/commands"
	"github.com/jessevdk/go-flags"
	"os"
	"fmt"
)

func main() {
	parser := flags.NewParser(&commands.Pivnet, flags.HelpFlag | flags.PassDoubleDash)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		if err == commands.ErrShowHelpMessage {
			helpParser := flags.NewParser(&commands.Pivnet, flags.HelpFlag)
			helpParser.NamespaceDelimiter = "-"
			helpParser.ParseArgs([]string{"-h"})
			helpParser.WriteHelp(os.Stdout)
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
		}
		os.Exit(1)
	}
}

