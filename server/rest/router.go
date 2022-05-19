package rest

import (
	"net/http"

	"sample/handlers"

	"github.com/gorilla/mux"
)

func router(userHandler handlers.UserHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{user_id}", userHandler.ReadUser).Methods(http.MethodGet)
	r.HandleFunc("/user", userHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{user_id}", userHandler.DeleteUser).Methods(http.MethodDelete)

	return r
}
