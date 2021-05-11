package main

import (
	"github.com/WisdomFusion/go_examples/26_filelistingserver/filelisting"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			logrus.Warnf("error handling request: %s", err.Error())
			code := http.StatusOK

			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(w, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/ls/", errWrapper(filelisting.HandleFileList))

	logrus.Fatalln(http.ListenAndServe(":8888", nil))
}
