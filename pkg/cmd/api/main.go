package main

import (
	"log"
	"os"

	"github.com/hpi/measurement/pkg/api"
	"github.com/hpi/measurement/pkg/datastore"
)

const (
	envDbHost   = "DB_HOST"
	envDbPort   = "DB_PORT"
	envHTTPPort = "HTTP_SERVER_PORT"
)

func main() {
	dbHost := os.Getenv(envDbHost)
	port := os.Getenv(envDbPort)
	if dbHost == "" || port == "" {
		log.Fatal("Database host/port not set in the Environment. Exiting..")
		os.Exit(0)
	}

	sessFactory, err := datastore.NewSessionFactory(dbHost, port)
	if err != nil {
		log.Fatalf("Unable to connect to mongo db : %s. Exiting., host : %s, port %s", err.Error(), "mongo", "27017")
		os.Exit(1)
	}

	httpPort := os.Getenv(envHTTPPort)
	if httpPort == "" {
		log.Fatal("HTTP port not set in environment. Exiting.")
		os.Exit(1)
	}

	ch := make(chan error)
	server := api.NewServer(datastore.NewMongo(sessFactory))

	server.Start(httpPort, ch)
	if err := <-ch; err == nil {
		log.Fatalf("HTTP server failure. Error : %s", err.Error())
		os.Exit(1)
	} else {
		log.Println("HTTP server was shut down.")
		os.Exit(0)
	}
}
