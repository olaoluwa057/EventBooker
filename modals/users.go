package modals

import (
	"errors"
	"example.com/event-booker/db"
	"example.com/event-booker/utils"
)

type User struct {
	ID       int64
	Name     string `binding:"required"`
	EMAIL    string `binding:"required"`
	PASSWORD string `binding:"required"`
	IsAdmin  bool
}

func (u *User) Save() error {
	query := `INSERT INTO users(name, email, password) VALUES(?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	password, err := utils.Hash(u.PASSWORD)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.EMAIL, password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	if id == 1 {
		err := setAdmin()
		if err != nil {
			return err
		}
	}

	u.ID = id

	return nil
}

func setAdmin() error {
	query := `UPDATE users SET is_admin = TRUE WHERE id = 1`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Validate() (error, User) {
	query := `SELECT id, is_admin, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, user.EMAIL)

	var hashedPassword string

	err := row.Scan(&user.ID, &user.IsAdmin, &hashedPassword)

	if err != nil {
		return errors.New("credentials incorrect"), User{}
	}
	err = utils.ComparePassword(user.PASSWORD, hashedPassword)

	if err != nil {
		return errors.New("invalid password"), User{}
	}
	return nil, User{ID: user.ID, IsAdmin: user.IsAdmin}

}

func GetUser(id int64) (User, error) {
	query := `SELECT id, name, email, is_admin FROM users WHERE id = ?`

	row := db.DB.QueryRow(query, id)

	var user User

	err := row.Scan(&user.ID, &user.Name, &user.EMAIL, &user.IsAdmin)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
