package commands

type PivnetCommand struct {
	Help HelpCommand `command:"help" description:"Print this help message"`
	Search SearchCommand `command:"search" description:"list or search packages"`
}

var Pivnet PivnetCommand