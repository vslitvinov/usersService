package psql

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/usersService/internal/models"
)


type Auth struct {
	pool *pgxpool.Pool
	//tables
	cache sync.Map
}

func NewAuthStorage(db *pgxpool.Pool) *Auth{
	return &Auth{db,sync.Map{}}
}

func (a *Auth) Create(UUIDUser string) (error, string) {
	
	token := uuid.NewString()

	// todo check user 

	a.cache.Store(token,models.Session{
		UUIDUser: UUIDUser,
		Expiry:   time.Now().Add(1200 * time.Second),
	})

	return nil, token

}

func (a *Auth) Find(token string) (error, models.Session){

	s, ok := a.cache.Load(token)
	if !ok {
		return errors.New("FIND SESSION - not fount session"), models.Session{}
	}

	return nil, s.(models.Session)

}

func (a *Auth) Update(token string) error {

	s, ok := a.cache.Load(token)
	if !ok {
		return errors.New("UPDATE SESSION - not fount session")
	}

	a.cache.Store(token, models.Session{
		UUIDUser: s.(models.Session).UUIDUser,
		Expiry: time.Now().Add(1200 * time.Second),
	})

	return nil

}

func (a *Auth) Delete(token string) error {

	_, ok := a.cache.LoadAndDelete(token)
	if !ok {
		return errors.New("DELETE SESSION - not fount session")
	}

	return nil
}
