package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	encode "github.com/wisdomfusion/go-encode"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("https://www.qq.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	e, _, _, err := encode.Find(resp.Body)
	if err != nil {
		panic(err)
	}

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	result, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result)
}
