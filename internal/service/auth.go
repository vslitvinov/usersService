package service

import "github.com/vslitvinov/usersService/internal/models"

type AuthStorage interface {
	Create(string) (error, string)
	Update(string) error
	Delete(string) error
	Find(string) (error, models.Session)
}

type Auth struct {
	storage AuthStorage
}

// construct Account
func NewAuthService(storage AuthStorage) *Auth {
	return &Auth{storage}
}

//Аутентификация пользовател    
func (s *Auth) SingIn(username,password string) string {

	return ""
}

//Регистрация пользовател
func (s *Auth) SingUp(models.User)  {

	// s.storage.Create 

}

// 
func (s *Auth) LogOut() {}
