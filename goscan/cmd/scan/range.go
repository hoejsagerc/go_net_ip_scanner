/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package scan

import (
	"fmt"
	"time"

	"github.com/hoejsagerc/go_net_ip_scanner/goscan/internal/scanner"
	"github.com/spf13/cobra"
)

var (
	startIp string
	endIp   string
	ports   []int
)

// rangeCmd represents the range command
var rangeCmd = &cobra.Command{
	Use:   "range",
	Short: "Scan a range of ip addresses",
	Long:  `This command will scan a range of ip addresses and return the results to the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		results, err := scanner.StartScan(startIp, endIp, (1 * time.Second), ports)
		if err != nil {
			fmt.Println(err)
		}

		for _, result := range results {
			if result.Open {
				if result.Hostname != "" {
					fmt.Printf("IP: %s (Hostname: %s), Port: %d is open\n", result.IP, result.Hostname, result.Port)
				} else {
					fmt.Printf("IP: %s, Port: %d is open\n", result.IP, result.Port)
				}
			}
		}
	},
}

func init() {
	rangeCmd.Flags().StringVarP((&startIp), "start", "s", "", "The start ip address in the range")
	rangeCmd.Flags().StringVarP((&endIp), "end", "e", "", "The end ip address in the range")
	rangeCmd.Flags().IntSliceVarP(&ports, "ports", "p", []int{80, 443}, "The ports to scan")

	if err := rangeCmd.MarkFlagRequired("start"); err != nil {
		fmt.Println(err)
	}

	if err := rangeCmd.MarkFlagRequired("end"); err != nil {
		fmt.Println(err)
	}

	ScanCmd.AddCommand(rangeCmd)
}
