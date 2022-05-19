package handlers

import (
	"encoding/json"
	"net/http"

	"sample/dtos"
	"sample/services"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type UserHandler interface {
	CreateUser(http.ResponseWriter, *http.Request)
	ReadUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

type userHandlerImpl struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandlerImpl{
		userService: userService,
	}
}

func (u *userHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	userData := &dtos.UserData{}
	if err := json.NewDecoder(r.Body).Decode(userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := govalidator.ValidateStruct(userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := u.userService.CreateUser(r.Context(), userData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *userHandlerImpl) ReadUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["user_id"]
	userData, err := u.userService.ReadUser(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (u *userHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userData := &dtos.UserData{}
	if err := json.NewDecoder(r.Body).Decode(userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := govalidator.ValidateStruct(userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := u.userService.UpdateUser(r.Context(), userData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (u *userHandlerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["user_id"]
	if err := u.userService.DeleteUser(r.Context(), userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
