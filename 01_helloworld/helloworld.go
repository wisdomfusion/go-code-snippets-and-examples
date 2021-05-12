package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		s := fmt.Sprintf("你好，世界！@ %s", time.Now().String())
		_, err := fmt.Fprintf(w, "%v\n", s)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%v\n", s)
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
