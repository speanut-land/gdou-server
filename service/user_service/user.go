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

func (u *User) ExistByTelephone() int {
	return models.IsTelephoneUsable(u.Telephone)
}

func (u *User) Add() error {
	return models.AddUser(u.Username, u.Password, u.Telephone)
}

func (u *User) ResetPassword() error {
	return models.ResetPassword(u.Telephone, u.Password)
}

func (u *User) Login() (bool, error) {
	return models.CheckLogin(u.Username, u.Password)
}
