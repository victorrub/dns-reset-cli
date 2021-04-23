package network

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/victorrub/dns-reset/infra/errors"
)

const pingCMD = "ping"

// CheckConnection .
func (c *Communicator) CheckConnection(domains []string, requestLimit int) (err error) {

	for _, domain := range domains {
		err = c.pingRequest(domain, requestLimit)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}

func (c *Communicator) pingRequest(domain string, requestLimit int) (err error) {

	cmd := exec.Command(pingCMD, domain, "-c", strconv.Itoa(requestLimit))

	output, err := cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err)
	}

	err = cmd.Start()
	if err != nil {
		return errors.Wrap(err)
	}

	matchTimeOut := regexp.MustCompile("cannot resolve")
	matchPackets := regexp.MustCompile("packets")

	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		message := scanner.Text()

		if matchTimeOut.MatchString(message) {
			fmt.Println("Cannot connect to", domain)
			fmt.Println("  *Error Message*  ", message)
			break
		}

		if matchPackets.MatchString(message) {
			fmt.Println("Ping statistics for", domain)
			fmt.Println(message)
			break
		}
	}

	err = cmd.Wait()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
