package contrrolers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ServerHTTP(W http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello form User Controller!"))
}

func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
