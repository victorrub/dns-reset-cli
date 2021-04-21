package main

import (
	"fmt"
	"os"

	"github.com/victorrub/dns-reset/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
