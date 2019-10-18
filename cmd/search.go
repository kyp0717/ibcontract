package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ibaccount/client"
	//"ibaccount/models"
	// "log"
	"encoding/json"
	"github.com/fatih/color"
)

// RespAccountList -- respond struct for endpt
type RespAccountList struct {
	// Response Struct for endpt ....
	Accounts []string `json:"accounts"`
	Aliases  struct {
	} `json:"aliases"`
	SelectedAccount string `json:"selectedAccount"`
}

// getstatusCmd represents the getstatus command
var listAcctCmd = &cobra.Command{
	Use:   "list",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, _ := listAccts()
		resp.Print()
	},
}

func init() {
	rootCmd.AddCommand(listAcctCmd)
}

// GetAccts -- return a list of accounts based on the login session
func listAccts() (RespAccountList, error) {
	url := client.BaseUrl + client.Accts
	data, _ := client.IbGet(url)
	var accts RespAccountList
	json.Unmarshal([]byte(data), &accts)
	return accts, nil
}

// Print to console
func (a RespAccountList) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" Accounts: %s \n", cyan(a.Accounts))
}
