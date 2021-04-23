package main

import (
	"github.com/victorrub/dns-reset/cmd"
	"github.com/victorrub/dns-reset/infra/errors"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		errors.EndAsErr(err, "Couldn't start CLI")
	}
}
