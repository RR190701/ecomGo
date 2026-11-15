package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/RR190701/ecomGo/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	appConfig := config{
		dbConfig: dbConfig{
			//always have fallbacks
			dsn: env.GetValueFromEnv("GOOSE_DBSTRING", "host=localhost port=5432 user=shopgo password=shopgo dbname=shopgo sslmode=disable"),
		},
		address: ":8080",
	}
	//structured logging using slog
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//DB connection
	conn, err := pgx.Connect(ctx, appConfig.dbConfig.dsn)

	if err != nil {
		//since the logger is defined after the db connection, we can use log package here
		//panic(err) - does the same thing as log.Fatal(err)
		slog.Error("Failed to connect to the database", "error", err.Error())
		os.Exit(-1)
	}
	defer conn.Close(ctx)

	logger.Info("Database connection established", "dsn", appConfig.dbConfig.dsn)

	api := application{
		config: appConfig,
	}

	handler := api.Server()
	if err := api.Start(handler); err != nil {
		//log.Fatal(err)
		slog.Error("Server failed to start", "error", err.Error())
		os.Exit(1)
	}

}
