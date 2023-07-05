package main

import (
	"analityc_test_task/cmd/providers"
	"analityc_test_task/configs"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const configPath = "configs/config.json"

func main() {
	cnf, err := configs.LoadConfig(configPath)
	if err != nil {
		log.Panic(err)
	}

	logger, err := providers.ProvideConsoleLogger(cnf)
	if err != nil {
		log.Panic(err)
	}

	db, closeDB, err := providers.ProvideDB(cnf, logger)
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	actionRepo := providers.ProvideActionRepository(db, logger)
	analitycsController := providers.ProvideAnalitycsController(ctx, actionRepo, logger)

	server := providers.ProvideHTTPServer(cnf, analitycsController, logger)

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		fmt.Println("Terminating the app")
		fmt.Println("Shutdown workers")
		cancel()

		fmt.Println("Close DB")
		closeDB()

		fmt.Println("Stop Server")
		server.Stop(ctx)
	}()

	server.Start()
}
