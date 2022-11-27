package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/usersService/internal/models"
	"github.com/vslitvinov/usersService/internal/storage/psql"
)



type AccountStorage interface {
	Create(ctx context.Context, ac models.Accounty) (string, error)
	FindByID(ctx context.Context, aid string) (models.Accounty, error)
	Verified(ctx context.Context, aid string, verified bool) error
	Archive(ctx context.Context, aid string, archive bool) error
}




type Account struct {
	storage AccountStorage
}

// construct Account
func NewAccountService(db *pgxpool.Pool) *Account {
	return &Account{storage:psql.NewAccountStorage(db) }
}
  
func (a *Account) Create(ctx context.Context, ac models.Accounty) (string, error) {

	var id string

	err := ac.GeneratePasswordHash()
	if err != nil {
		return id, fmt.Errorf("Servise.Account.Create %w", err)
	}

	id, err = a.storage.Create(ctx,ac)
	if err != nil {
		return id, fmt.Errorf("Service.Account.Create %w", err)
	}

	return id, nil 

}

func (a *Account) GetByID(ctx context.Context, aid string) (models.Accounty, error) {

	var ac models.Accounty

	ac, err := a.storage.FindByID(ctx, aid)
	if err != nil {
		return ac, fmt.Errorf("Service.Account.GetByID %w", err)
	}

	return ac, nil

}

func (a *Account) GetByEmail(ctx context.Context, email string) (models.Accounty, error) {

	var ac models.Accounty

	ac, err := a.storage.FindByID(ctx, email)
	if err != nil {
		return ac, fmt.Errorf("Service.Account.GetByEmail %w", err)
	}

	return ac, nil
}

func (a *Account) Delete(ctx context.Context, aid, sid string) error {

	// Archive
	// Finish session

}