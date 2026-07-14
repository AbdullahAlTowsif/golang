package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
	// List() ([]*User, error)
	// Delete(userId int) error
	// Update(user User) (*User, error)
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

// func (r userRepo) Create(user User) (*User, error) {
// 	if user.ID != 0 {
// 		return &user, nil
// 	}

// 	user.ID = len(r.users) + 1

// 	r.users = append(r.users, user)
// 	return &user, nil
// }

func (r userRepo) Create(user User) (*User, error) {
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

func (r *userRepo) Find(email, password string) (*User, error) {
	var user User

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

// func (r *userRepo) Find(email, password string) (*User, error) {
// 	for _, u := range r.users {
// 		if u.Email == email && u.Password == password {
// 			return &u, nil
// 		}
// 	}
// 	return nil, nil
// }
