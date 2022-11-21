package service

import "github.com/vslitvinov/usersService/internal/models"


type AccountStorage interface {
	Create(models.User)
	Update(models.User)
	Delete(string)
	Get(string)
	Find()
}

type Account struct {
	storage AccountStorage
}

// construct Users

func NewUsers (storage AccountStorage) *Account{
	return &Account{storage}
}
