package cmd

import (
	"fmt"
	"ibcontract/client"

	"github.com/spf13/cobra"

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

type Security struct {
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

type RspSecurityArray struct {
	Payload []Security
}

// getstatusCmd represents the getstatus command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, _ := search("AAPL")
		var stock Security
		stock = resp[0]
		//fmt.Println(string(resp))
		// fmt.Println(resp)
		stock.Print()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

// GetAccts -- return a list of accounts based on the login session
func search(symbol string) ([]Security, error) {
	url := client.BaseUrl + client.ContractSearch
	reqBody := map[string]string{
		"symbol": symbol,
	}

	data, _ := client.IbPost(url, reqBody)
	securities := make([]Security, 0)
	//var rsp RspSecurityArray
	json.Unmarshal([]byte(data), &securities)
	return securities, nil
}

// Print to console
func (a Security) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" ConID: %s \n", cyan(a.Conid))
	fmt.Printf(" Company: %s \n", cyan(a.CompanyName))

}
