package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zjyl1994/unilinkd/config"
	"github.com/zjyl1994/unilinkd/server"
)

func main() {
	var err error
	// load config
	err = config.LoadConf("config.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	// run server
	err = http.ListenAndServe(config.ListenAddr, server.UnilinkdServer{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
}
