package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"ibcontract/client"
	"io/ioutil"
	"net/http"
	// "log"
	"encoding/json"

	"github.com/fatih/color"
)

/*PostIserverSecdefSearchBody post iserver secdef search body
swagger:model PostIserverSecdefSearchBody
*/
type ReqSecurityBody struct {
	// should be true if the search is to be performed by name. false by default.
	Name bool `json:"name,omitempty"`
	// If search is done by name, only the assets provided in this field will be returned. Currently, only STK is supported.
	SecType string `json:"secType,omitempty"`
	// symbol or name to be searched
	// Required: true
	Symbol *string `json:"symbol"`
}

type ReqContractBody struct {
	Conids [1]string `json:"conids"`
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

type Contract struct {
	Conid          int    `json:"conid"`
	Name           string `json:"name"`
	AssetClass     string `json:"assetClass"`
	Expiry         string `json:"expiry"`
	LastTradingDay string `json:"lastTradingDay"`
	Group          string `json:"group"`
	PutOrCall      string `json:"putOrCall"`
	Sector         string `json:"sector"`
	SectorGroup    string `json:"sectorGroup"`
	Strike         int    `json:"strike"`
	Ticker         string `json:"ticker"`
	UndConid       int    `json:"undConid"`
	FullName       string `json:"fullName"`
	PageSize       int    `json:"pageSize"`
}

type RspSecurity struct {
	Securities []Security
}

type RspContract struct {
	Contracts []Contract
}

// getstatusCmd represents the getstatus command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, _ := searchSymbol("AAPL")
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
func searchSymbol(symbol string) ([]Security, error) {
	url := client.BaseUrl + client.ContractSymbol
	reqBody := map[string]string{
		"symbol": symbol,
	}

	data, _ := client.IbPost(url, reqBody)
	securities := make([]Security, 0)
	//var rsp RspSecurityArray
	json.Unmarshal([]byte(data), &securities)
	return securities, nil
}

// Search by ConID-- return a list of accounts based on the login session
func searchConId(conid string) ([]Contract, error) {
	url := client.BaseUrl + client.ContractConId
	listconids := [1]string{conid}
	req := ReqContractBody{listconids}

	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	resp, _ := http.Post(url, "application/jso", bytes.NewBuffer(reqBody))
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	contracts := make([]Contract, 0)
	//var rsp RspSecurityArray
	json.Unmarshal([]byte(data), &contracts)
	return contracts, nil
}

// Print to console
func (a Security) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" ConID: %s \n", cyan(a.Conid))
	fmt.Printf(" Company: %s \n", cyan(a.CompanyName))

}
