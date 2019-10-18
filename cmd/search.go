package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ibcontract/client"
	// "log"
	"encoding/json"
	"github.com/fatih/color"
)

/*PostIserverSecdefSearchBody post iserver secdef search body
swagger:model PostIserverSecdefSearchBody
*/
type ReqSearchBody struct {
	// should be true if the search is to be performed by name. false by default.
	Name bool `json:"name,omitempty"`
	// If search is done by name, only the assets provided in this field will be returned. Currently, only STK is supported.
	SecType string `json:"secType,omitempty"`
	// symbol or name to be searched
	// Required: true
	Symbol *string `json:"symbol"`
}

type RspSearchBody struct {
	// company header
	CompanyHeader string `json:"companyHeader,omitempty"`
	// company name
	CompanyName string `json:"companyName,omitempty"`
	// conid
	Conid int64 `json:"conid,omitempty"`
	// description
	Description string `json:"description,omitempty"`
	// opt
	Opt string `json:"opt,omitempty"`
	// sections
	Sections []interface{} `json:"sections"`
	// symbol
	Symbol string `json:"symbol,omitempty"`
	// war
	War string `json:"war,omitempty"`
}

// getstatusCmd represents the getstatus command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, _ := search2("AAPL")
		fmt.Println(string(resp))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

// GetAccts -- return a list of accounts based on the login session
func search(symbol string) (RspSearchBody, error) {
	url := client.BaseUrl + client.ContractSearch
	reqBody := map[string]string{
		"symbol": symbol,
	}

	data, _ := client.IbPost(url, reqBody)
	var rsp RspSearchBody
	json.Unmarshal([]byte(data), &rsp)
	return rsp, nil
}

// GetAccts -- return a list of accounts based on the login session
func search2(symbol string) ([]byte, error) {
	url := client.BaseUrl + client.ContractSearch
	reqBody := map[string]string{
		"symbol": symbol,
	}

	data, _ := client.IbPost(url, reqBody)
	return data, nil
}

// Print to console
func (a RspSearchBody) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" ConID: %s \n", cyan(a.Symbol))
}
