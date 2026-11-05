package main

import (
	"log/slog"
	"os"
)

func main() {

	appConfig := config{
		dbConfig: dbConfig{},
		address:  ":8080",
	}

	api := application{
		config: appConfig,
	}

	//structured logging using slog
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	handler := api.Server()
	if err := api.Start(handler); err != nil {
		//log.Fatal(err)
		slog.Error("Server failed to start", err)
		os.Exit(1)
	}

}
