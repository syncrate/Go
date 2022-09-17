package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpClient()
}
func httpClient() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func del() {

}