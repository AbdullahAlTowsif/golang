package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	// users []User
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
	INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
	VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
	RETURNING id
	`

	// execute named query
	var userId int
	rows, err := r.dbCon.NamedQuery(query, user)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userId)
	}

	user.ID = userId
	return &user, nil
}

func (r *userRepo) Find(email, password string) (*domain.User, error) {
	var user domain.User

	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1
	`

	err := r.dbCon.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
