package main

import (
	"fmt"
	"github.com/WisdomFusion/go_examples/24_interfaces_v2/retriever"
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
