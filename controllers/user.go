package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gopher, from userController..."))
}

func newuserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}