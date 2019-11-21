package util

import (
	"fmt"
	"github.com/speanut-land/gdou-server/pkg/setting"
	"math/rand"
	"time"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

//生成6位随机验证码
func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
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
