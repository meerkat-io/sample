package rest

import (
	"log"
	"net/http"

	"sample/handlers"
	"sample/repositories"
	"sample/services"

	"github.com/gorilla/mux"
)

func Start(listen string, r repositories.Repository) {
	userHandler := handlers.NewUserHandler(services.NewUserService(r.UserRepo()))

	router := mux.NewRouter()

	router.HandleFunc("/user", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{user_id}", userHandler.ReadUser).Methods(http.MethodGet)
	router.HandleFunc("/user", userHandler.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{user_id}", userHandler.DeleteUser).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:    listen,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
}
