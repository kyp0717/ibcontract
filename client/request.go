/******

*****/
package client

import (
	"fmt"
	"io/ioutil"
	// "log"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
)

func setPolicy() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func IbGet(url string) ([]byte, error) {
	setPolicy()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP Req failed with error %s\n", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}

func IbPost(url string, req map[string]string) ([]byte, error) {
	setPolicy()
	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("HTTP Req failed with error %s\n", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("HTTP Req failed with error %s\n", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}
