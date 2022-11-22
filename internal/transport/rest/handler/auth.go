package handler

import (
	"fmt"
	"net/http"
)

// описание поведения структуры обьявленой в сервисе
type AuthService interface {

}

type AuthHandler struct {
	service AccountService
}

// construct AccountHandler
func NewAuthHandler(s AccountService) *AuthHandler{
	return &AuthHandler{s}
}

// gin.Context or standert http mux ???
func (h *AuthHandler) SingIn(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"test")
}
func (h *AuthHandler) SingUp(w http.ResponseWriter, r *http.Request){

}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request){

}