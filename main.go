package main

import (
	"fmt"
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
	err = server.Run(config.ListenAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
}
