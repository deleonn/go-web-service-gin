package main

import (
	"flag"
	"fmt"
	"os"

	"web-service-gin/config"
	"web-service-gin/db"
	"web-service-gin/server"
)

func main() {
	environment := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Connect()
	server.Init()
}
