package service

import (
	"context"
	"fmt"

	"github.com/vslitvinov/usersService/internal/models"
)


type account interface {
	Create(ctx context.Context, ac models.Accounty) (string, error) 
	GetByID(ctx context.Context, aid string) (models.Accounty, error) 
	GetByEmail(ctx context.Context, email string) (models.Accounty, error) 
	Delete(ctx context.Context, aid, sid string) error 
}
 

type Auth struct {
	as account
	ss session
}

// construct Account
func NewAuthService(as account,ss session) *Auth {
	return &Auth{as,ss}
}

func (a *Auth) EmailSingIn(ctx context.Context, email, password string, d Device)  (models.Session, error) {

	var ms models.Session 

	ma, err := a.as.GetByEmail(ctx,email)
	if err != nil {
		return ms, fmt.Errorf("service.auth.EmailSingIn.GetByEmail %w", err)
	}

	err = ma.CompareHashAndPassword(password)
	if err != nil {
		return ms, fmt.Errorf("service.auth.EmailSingIn.CompareHashAndPassword %w", err)
	}

	ms, err = a.ss.Create(ctx,ma.ID,"email",d)
	if err != nil {
		return ms, fmt.Errorf("service.auth.EmailSingIn.CreateSession %w", err)
	}

	return ms, nil
}

func (a *Auth) SingUp(models.Accounty)  {


}

func (a *Auth) LogOut(ctx context.Context, sid string) error {
	if err := a.ss.Finish(ctx, sid, ""); err != nil {
		return fmt.Errorf("authService.LogOut.session.Finish: %w", err)
	}

	return nil
}
