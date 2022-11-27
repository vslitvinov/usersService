package service

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/usersService/internal/models"
)



type Auth struct {
}

// construct Account
func NewAuthService(db *pgxpool.Pool) *Auth {
	return &Auth{}
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
