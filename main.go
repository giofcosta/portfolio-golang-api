package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/giofcosta/portfolio-golang-api/config"
	"github.com/giofcosta/portfolio-golang-api/server"
)

func main() {

	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)
	server.Init()
}
