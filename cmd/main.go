package main

import (
	"analityc_test_task/cmd/config"
	"fmt"
	"log"
)

const configPath = "cmd/config/config.json"

func main() {
	cnf, err := config.LoadConfiguration(configPath)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(cnf)
}
