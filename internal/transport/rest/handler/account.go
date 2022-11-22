package handler

import "net/http"

// описание поведения структуры обьявленой в сервисе
type AccountService interface {

}

type AccountHandler struct {
	service AccountService
}

// construct AccountHandler
func NewAccountHandler(s AccountService) *AccountHandler{
	return &AccountHandler{s}
}

// gin.Context or standert http mux ???
func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request){

}