package main

import (
	"fmt"

	"github.com/wisdomfusion/go_examples/basics/interfaces2/retriever"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = retriever.Retriever{}
	fmt.Println(download(r))
}
