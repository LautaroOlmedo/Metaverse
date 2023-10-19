package handlers

import (
	"encoding/json"
	"fmt"
	"metaverse/internal/core/services"
	"metaverse/internal/infrastructure/http/DTOs"
	"net/http"
	"regexp"
)

var (
	listUserRe   = regexp.MustCompile(`^\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/users\/(\d+)$`) // ---> /users/123
	createUserRe = regexp.MustCompile(`^\/users[\/]*$`)
	updateUserRe = regexp.MustCompile(`^\/users\/(\d+)$`) // ---> /users/123
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: &userService,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {

	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		h.Register(w, r)
		return

	default:
		http.NotFound(w, r)
		return
	}

}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var user DTOs.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.userService.Register(user.Name, user.DNI, user.Username, user.Email, user.Password, user.Age)
	if err != nil {
		fmt.Println(err)
		internalServerError(w, r)

	}
	w.WriteHeader(http.StatusCreated)
}

func userNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`error: user not found`))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`error: internal server error`))
}
