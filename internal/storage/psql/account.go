package psql

import "github.com/vslitvinov/usersService/internal/models"



type Account struct {
	//pool
	//tables
}

func (u *Account) Create(user models.User)  {}
func (u *Account) Update(user models.User) {}
func (u *Account) Delete(userUUID string) {}
func (u *Account) Get(userUUID string){}
func (u *Account) Find(){}