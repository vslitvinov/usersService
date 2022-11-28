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
	SingUp(ctx context.Context, ma models.Accounty) (string,error)
}

type AuthHandler struct {
	service AuthService 
}

// construct AccountHandler
func NewAuthHandler(s AuthService) *AuthHandler{

	handler := AuthHandler{s}

	return &handler
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

	if r.Method != "POST" {
		return 
	}

	r.ParseForm()                    
    email := r.Form.Get("email")
    password := r.Form.Get("password")
    firstname := r.Form.Get("firstname")
    lastname := r.Form.Get("lastname")
    username := r.Form.Get("username")

    ac := models.Accounty{
    	FirstName:  firstname,
    	LastName:   lastname,
    	Email:      email,
    	Phone:      "",
    	Username:   username,
    	Password:   password,
    }

    id, err := h.service.SingUp(r.Context(),ac)

     if err != nil {
    	log.Println(err)
    }

	fmt.Fprintf(w,id)

}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request){

}