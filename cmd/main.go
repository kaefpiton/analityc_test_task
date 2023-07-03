package main

import (
	"analityc_test_task/cmd/config"
	"analityc_test_task/cmd/providers"
	"log"
)

const configPath = "cmd/config/config.json"

func main() {
	cnf, err := config.LoadConfig(configPath)
	if err != nil {
		log.Panic(err)
	}
	server := providers.ProvideHTTPServer(cnf)

	server.Start()
}
