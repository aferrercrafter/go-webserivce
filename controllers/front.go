package controllers

import (
	"net/http"
)

//RegisterControllers registering controllers
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
