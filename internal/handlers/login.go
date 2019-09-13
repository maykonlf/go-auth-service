// @Summary check service health
// @Description check service health
// @ID healthz
// @Accept json
// @Produce json
// @Success 200
// @Failure 500
// @Route /login [post]
package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/maykonlf/go-auth-service/internal/controllers"
	"github.com/maykonlf/go-auth-service/internal/entities"
	"net/http"
)

type LoginHandlerI interface {
	Router(r chi.Router)
	Post(w http.ResponseWriter, r *http.Request)
}

func NewLoginHandler() LoginHandlerI {
	return &LoginHandler{
		loginController: controllers.NewLoginController(),
	}
}

type LoginHandler struct {
	loginController controllers.LoginControllerI
}

func (h *LoginHandler) Router(r chi.Router) {
	r.Post("/", h.Post)
}

func (h *LoginHandler) Post(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.loginController.Authenticate(&user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	return
}
