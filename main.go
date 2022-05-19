package main

import (
	"os"
	"os/signal"
	"syscall"

	"sample/infrastructures/logger"
	"sample/repositories"
	"sample/server/rest"
)

func main() {
	_ = logger.SetLevel(logger.DEBUG)

	listen := ":8080"
	dbURL := "mongodb://localhost:27017"
	dbName := "sample"

	r, err := repositories.NewRepository(dbURL, dbName)
	if err != nil {
		logger.Criticalf("connect to database failed %s", err.Error())
		os.Exit(1)
	}

	rest.Start(listen, r)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logger.Infof("service started on %s", listen)
	<-c
	logger.Info("service stopping...")
}
