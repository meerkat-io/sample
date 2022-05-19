package rest

import (
	"log"
	"net/http"

	"sample/handlers"
	"sample/repositories"
	"sample/services"
)

func Start(listen string, r repositories.Repository) {
	userHandler := handlers.NewUserHandler(services.NewUserService(r.UserRepo()))

	srv := &http.Server{
		Addr:    listen,
		Handler: router(userHandler),
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
}
