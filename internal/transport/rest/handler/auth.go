package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/vslitvinov/usersService/internal/models"
	"github.com/vslitvinov/usersService/internal/service"
)

// описание поведения структуры обьявленой в сервисе
type AuthService interface {
	EmailSingIn(ctx context.Context, email, password string, d service.Device)  (models.Session, error)
}

type AuthHandler struct {
	service AuthService 
}

// construct AccountHandler
func NewAuthHandler(s AuthService) *AuthHandler{
	return &AuthHandler{s}
}

// gin.Context or standert http mux ???
func (h *AuthHandler) SingIn(w http.ResponseWriter, r *http.Request){

	if r.Method != "POST" {
		return 
	}

	r.ParseForm()                    
    email := r.Form.Get("email")
    password := r.Form.Get("password")

    d := service.Device{
    	UserAgent: r.Header.Get("User-Agent"),
    	IP:        r.RemoteAddr,
    }

	s, err := h.service.EmailSingIn(r.Context(),email,password,d)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w,s.ID)

}
func (h *AuthHandler) SingUp(w http.ResponseWriter, r *http.Request){

}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request){

}