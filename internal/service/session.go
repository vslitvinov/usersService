package service

import (
	"context"

	"github.com/vslitvinov/usersService/internal/models"
)


type SessionStorage interface {
	Create(ctx context.Context, ms models.Session)
	FindByID(ctx context.Context, sid string) (models.Session, error)
	FindAll(ctx context.Context, aid string) ([]models.Session, error)
	Delete(ctx context.Context, sid string) error 
	DeleteAll(ctx context.Context, aid, sid string) error
}

type Session struct {
	storage SessionStorage
}

func NewSessionStorage() *Session {
	return &Session{}
}

// Device represents data transfer object with user device data
type Device struct {
	UserAgent string
	IP        string
}

// Create new session
func (s *Session) Create(ctx context.Context, aid, provider string, d Device) (models.Session, error) {
	return models.Session{}, nil
}

// GetByID session.
func (s *Session) GetByID(ctx context.Context, sid string) (models.Session, error) {
	return models.Session{}, nil
}

// GetAll account sessions using provided account id.
func (s *Session) GetAll(ctx context.Context, aid string) ([]models.Session, error){
	return []models.Session{}, nil
}

// Finish session by id excluding current session with id.
func (s *Session) Finish(ctx context.Context, sid, currSid string) error {
	return nil
}

// FinishAll account sessions excluding current session with id.
func (s *Session) FinishAll(ctx context.Context, aid, sid string) error {
	return nil
}