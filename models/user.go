package models

import (
	"github.com/jinzhu/gorm"
	"github.com/speanut-land/gdou-server/pkg/e"
	"regexp"
)

type User struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
}

//用户名是否可用
/*
	1.查询用户名是否重复
*/
func ExistUserByName(name string) (bool, error) {
	var user User
	err := db.Select("id").Where("username = ?", name).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// AddUser Add a User
func AddUser(username string, password string, telephone string) error {
	user := User{
		Username:  username,
		Password:  password,
		Telephone: telephone,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

//判断用户能否登录
func CheckLogin(username string, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

//电话号码是否可用
/*
	1.电话号码是否符合规则
	2.是否存在相同的电话号码
*/
func IsTelephoneUsable(phone string) int {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	isPassReg := rgx.MatchString(phone)

	if isPassReg {
		var user User
		err := db.Select("id").Where("telephone = ?", phone).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return e.ERROR
		}
		if err == gorm.ErrRecordNotFound {
			return 0
		}
		if user.ID > 0 {
			return e.ERROR_TELEPHONE_USED
		}
	}
	return e.ERROR_TELEPHONE_FORMAT
}
