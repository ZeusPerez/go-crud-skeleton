package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ZeusPerez/go-crud-skeleton/internal/adapters/storage"
	"github.com/ZeusPerez/go-crud-skeleton/internal/adapters/transport"
	"github.com/ZeusPerez/go-crud-skeleton/internal/services"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

const listerAddr = ":8000"

func main() {

	// Inint MySQl adapter
	mysqlCfg := storage.MySQLConfig{}
	err := envconfig.Process("DEVS_CRUD", &mysqlCfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbAdapter, err := storage.NewMySQLDev(mysqlCfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer dbAdapter.Close()

	// Create Devs CRUD service
	devsService := services.NewDevs(dbAdapter)

	// Init HTTP adapter
	httpCfg := transport.HttpConfig{}
	err = envconfig.Process("DEVS_CRUD", &httpCfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	httpAdapter := transport.NewHttpAdapter(httpCfg, devsService)

	mux := http.NewServeMux()
	httpAdapter.AddHandlers(mux)

	// Start the HTTP server
	log.Infof("Starting server on %s", listerAddr)
	err = http.ListenAndServe(":8000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
