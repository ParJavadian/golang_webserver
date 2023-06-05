package service

import (
	"database/sql"
	"errors"
)

type User struct {
	Name      string `json:"name"`
	Family    string `json:"family"`
	Id        int64  `json:"id"`
	Age       int64  `json:"age"`
	Sex       string `json:"sex"`
	CreatedAt int64  `json:"createdAt"`
}

func (p *User) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (p *User) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (p *User) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (p *User) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func getUsers(db *sql.DB, start, count int) ([]User, error) {
	return nil, errors.New("Not implemented")
}
