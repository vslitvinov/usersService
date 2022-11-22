package psql

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/usersService/internal/models"
)

type Account struct {
	pool  *pgxpool.Pool
	cache sync.Map
}

func NewAccount(db *pgxpool.Pool) *Account {
	return &Account{db, sync.Map{}}
}

func (a *Account) Create(ctx context.Context, user models.User) error {
	sql := `INSERT INTO users (uuid, firstname, lastname, displayname, phone, email, password)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING user_id`
	
	a.pool.QueryRow(ctx, sql,
		user.UUID,
		user.FirstName,
		user.LastName,
		user.DisplayName,
		user.Phone,
		user.Email,
		user.Password,
	)
	// var id uint64
	// err := row.Scan(&id)
	// if err != nil {
	// 	return err
	// }
	// return id, nil
	return nil
}

func (a *Account) Update(user models.User) {

}

func (a *Account) Delete(userUUID string) {}

func (a *Account) Get(ctx context.Context, userUUID string) {
		sql := `SELECT * FROM "users"`
	rows, err := a.pool.Query(ctx,sql)
	if err != nil {
		// return nil,err
	}
	var data []models.User

	for rows.Next() {
		d := models.User{}
		err = rows.Scan(&d.UUID,&d.FirstName,&d.LastName,&d.DisplayName,&d.Email,&d.Password)
		if err != nil {
			// log.Println(err)
		}	
		data = append(data,d)
	}
}

func (a *Account) GetAll() {}

func (a *Account) Find() {}
