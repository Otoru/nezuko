package main

import (
	"os"

	"github.com/otoru/nezuko/cmd"
)

func main() {
	code := 0

	if err := cmd.Execute(); err != nil {
		code = 1
	}

	os.Exit(code)
}
