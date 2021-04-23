package errors

import (
	"fmt"
	"os"
	"time"
)

// EndAsErr logs the error and ends the service process
func EndAsErr(err error, message string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		fmt.Fprintln(os.Stdout, message)
		time.Sleep(time.Millisecond * 50) // needed for printing all messages before exiting
		os.Exit(1)
	}
}
