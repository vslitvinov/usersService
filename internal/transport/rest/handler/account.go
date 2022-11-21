package handler


// описание поведения структуры обьявленой в сервисе 
type UsersService interface {

}

type AccountHandler struct {
	service UsersService
}

// construct UserHandler

func NewUsersHandler(s UsersService) *UsersHandler{
	return &UsersHandler{s}
}


func 