package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/usama28232/candid/controllers"
	"github.com/usama28232/candid/db"
	"github.com/usama28232/candid/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("*** starting `candid framework` ***")
	logger.Info("https://github.com/usama28232/candid")
	var port string
	flag.StringVar(&port, "port", "3000", "Specifies Application Port")
	flag.Parse()

	dbconn, dbSErr := db.Init()
	logger.Debug(dbconn)
	if dbSErr == nil {

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

		go startServer(port)

		<-signals

		dbCErr := db.Close()
		logger.Debugw("closing db connection & release resources")
		if dbCErr != nil {
			logger.Error(dbCErr.Error())
		}

	} else {
		logger.Info(dbSErr.Error())
	}
}

func startServer(port string) {
	logging.GetLogger().Infow("Starting Webservice with", "port", port)
	http.ListenAndServe(":"+port, controllers.RegisterRoutes())
}
