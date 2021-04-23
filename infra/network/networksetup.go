package network

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/victorrub/dns-reset/infra/errors"
)

const netCmd = "networksetup"

// GetCurrentLocation .
func (c *Communicator) GetCurrentLocation() (location string, err error) {

	out, err := exec.Command(netCmd, "-getcurrentlocation").Output()
	if err != nil {
		return location, errors.Wrap(err)
	}

	return c.getMessage(out), nil
}

// ListLocations .
func (c *Communicator) ListLocations() (locations []string, err error) {

	locations = make([]string, 0)

	cmd := exec.Command(netCmd, "-listlocations")

	output, err := cmd.StdoutPipe()
	if err != nil {
		return locations, errors.Wrap(err)
	}

	err = cmd.Start()
	if err != nil {
		return locations, errors.Wrap(err)
	}

	scanner := bufio.NewScanner(output)

	for scanner.Scan() {
		location := scanner.Text()

		locations = append(locations, location)
	}

	err = cmd.Wait()
	if err != nil {
		return locations, errors.Wrap(err)
	}

	return locations, nil
}

// SwitchLocation .
func (c *Communicator) SwitchLocation(name string) (err error) {

	output, _ := exec.Command(netCmd, "-switchtolocation", name).Output()

	matchSuccessMessage, _ := regexp.MatchString("found it!", c.getMessage(output))

	if !matchSuccessMessage {
		errMessage := fmt.Sprintf("Err: %s", c.getMessage(output))
		return errors.NewApplicationError("networksetup.SwitchLocation", errMessage)
	}

	return nil
}

// CreateLocation .
func (c *Communicator) CreateLocation(name string) (err error) {

	output, _ := exec.Command(netCmd, "-createlocation", name, "populate").Output()

	matchSuccessMessage, _ := regexp.MatchString("populated", c.getMessage(output))

	if !matchSuccessMessage {
		errMessage := fmt.Sprintf("Err: %s", c.getMessage(output))
		return errors.NewApplicationError("networksetup.CreateLocation", errMessage)
	}

	return nil
}

// DeleteLocation .
func (c *Communicator) DeleteLocation(name string) (err error) {

	output, _ := exec.Command(netCmd, "-deletelocation", name).Output()

	matchSuccessMessage, _ := regexp.MatchString("found it!", c.getMessage(output))

	if !matchSuccessMessage {
		errMessage := fmt.Sprintf("Err: %s", c.getMessage(output))
		return errors.NewApplicationError("networksetup.DeleteLocation", errMessage)
	}

	return nil
}

func (c *Communicator) getMessage(cmdOutput []byte) string {
	return strings.TrimSpace(string(cmdOutput))
}
