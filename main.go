package main

import (
	"context"
	"fmt"
	"os"

	"github.com/colt005/TrackPixel/pkg/api"
	"github.com/colt005/TrackPixel/pkg/app"
	"github.com/colt005/TrackPixel/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	logger "github.com/sirupsen/logrus"
)

var (
	BASE_URL = os.Getenv("BASE_URL")

	db  *pgx.Conn
	err error
)

func main() {
	if BASE_URL == "" {
		logger.Fatal("BASE_URL not found")

	}
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}

}

// func run will be responsible for setting up db connections, routers etc
func run() error {

	connectionString := os.Getenv("CONN_STRING")

	if connectionString == "" {
		panic("database connection string(`CONN_STRING`) not defined")
	}

	// setup database connection
	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}

	// create storage dependency
	storage := repository.NewStorage(db)

	// create router dependecy
	router := gin.Default()

	// create tracker service
	trackService := api.NewTrackerService(storage)

	server := app.NewServer(router, trackService)

	// start the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*pgx.Conn, error) {

	ctx := context.TODO()

	db, err = pgx.Connect(ctx, connString)

	if err != nil {
		logger.Error(err)
	}

	err = db.Ping(context.TODO())

	if err != nil {
		logger.Error(err)
	}

	return db, nil

}
