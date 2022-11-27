package psql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/usersService/internal/models"
)

type Session struct {
	pool  *pgxpool.Pool
}

func NewSessionStorage(db *pgxpool.Pool) *Session{
	return &Session{
		pool: db,
	}
}

// Create new session in DB.
func (s *Session) Create(ctx context.Context, ms models.Session) error {
	return nil
}

// FindByID session.
func (s *Session) FindByID(ctx context.Context, sid string) (models.Session, error){
	return models.Session{}, nil
}

// FindAll accounts sessions by provided account id.
func (s *Session) FindAll(ctx context.Context, aid string) ([]models.Session, error){
	return []models.Session{}, nil
}

// Delete session by id.
func (s *Session) Delete(ctx context.Context, sid string) error {
	return nil
}

// DeleteAll account sessions by provided account id excluding current session.
func (s *Session) DeleteAll(ctx context.Context, aid, sid string) error {
	return nil
}