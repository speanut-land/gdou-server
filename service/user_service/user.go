package user_service

import "github.com/speanut-land/gdou-server/models"

type User struct {
	ID        int
	Username  string
	Password  string
	Telephone string
	Email     string
}

func (u *User) ExistByName() (bool, error) {
	return models.ExistUserByName(u.Username)
}

func (u *User) Add() error {
	return models.AddUser(u.Username, u.Password, u.Telephone)
}
