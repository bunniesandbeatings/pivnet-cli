package util

import (
	"runtime"
	"os"
)

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("USERPROFILE")
		if home == "" {
			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		}

		if home == "" {
			panic("could not detect home directory for .pivnetrc")
		}

		return home
	}

	return os.Getenv("HOME")
}