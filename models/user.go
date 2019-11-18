package models

import (
	"github.com/jinzhu/gorm"
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
	err := db.Where("username = ?", name).First(&user).Error
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

//TODO 增加工具类

//电话号码是否可用
/*
	1.电话号码是否符合规则
	2.是否存在相同的电话号码
*/
func IsTelephoneUsable(phone string) (bool, error) {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	isPassReg := rgx.MatchString(phone)

	if isPassReg {
		var user User
		err := db.Where("telephone = ?", phone).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return false, err
		}
		if user.ID > 0 {
			return true, nil
		}
	}
	return false, nil
}

//用户密码是否可用
/*
	1.查询用户密码是否不少于6位
*/
func IsPasswordUsable(password string) bool {
	length := len(password)
	if length >= 6 {
		return true
	}
	return false
}
