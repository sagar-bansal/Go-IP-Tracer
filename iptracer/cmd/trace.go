/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Tracing the IP",
	Long:  `Trace the IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)

			}

		} else {
			fmt.Println("Please enter the IP")
		}
	},
}

type Ip struct {
	IP string `json::"ip"`

	City string `json::"city"`

	Region string `json::"region"`

	Country string `json::"country"`

	Loc string `json::"loc"`

	Timezone string `json::"timezone"`

	Postal string `json::"postal"`
}

func init() {
	rootCmd.AddCommand(traceCmd)

}
func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)
	data := Ip{}
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		fmt.Println("Unable to unmarshal the response")
	} else {
		fmt.Println(data)
	}
}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("not able to get response")
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to read response")
	}
	return responseByte
}
