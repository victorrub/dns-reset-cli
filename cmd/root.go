package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/victorrub/dns-reset/domain"
	"github.com/victorrub/dns-reset/domain/contract"
	"github.com/victorrub/dns-reset/domain/service"
	"github.com/victorrub/dns-reset/infra/errors"
	"github.com/victorrub/dns-reset/infra/network"
)

var (
	ArgName            string
	ArgCheckConnection bool
)

var rootCmd = &cobra.Command{
	Use:   "dns-reset",
	Short: "DNS-Reset automates the process of setting up a new network location on Mac OS",
	Run: func(cmd *cobra.Command, args []string) {

		svc, err := initNetworkService()
		if err != nil {
			errors.EndAsErr(err, "Could not create service structure")
		}

		if ArgName == "" {
			t := time.Now()
			ArgName = fmt.Sprintf("%s_%s", domain.NewDNSPrefix, t.Format("20060102150405"))
		}

		err = svc.ResetLocation(ArgName)
		if err != nil {
			errors.EndAsErr(err, "Could not reset location")
		}

		if ArgCheckConnection {
			err = svc.CheckConnection()
			if err != nil {
				errors.EndAsErr(err, "Could not check the network connection")
			}
		}

		fmt.Println("DNS was successfully reset!")
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&ArgName, "name", "n", "", "Name for network location")
	rootCmd.Flags().BoolVarP(&ArgCheckConnection, "check", "c", false, "Performs ping tests after creating a new network location")
}

func initNetworkService() (contract.NetworkService, error) {
	net := network.Communicator{}

	svc, err := service.New(&net)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return service.NewNetworkService(svc), nil
}
