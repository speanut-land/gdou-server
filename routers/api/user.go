package api

import (
	"github.com/gin-gonic/gin"
	"github.com/speanut-land/gdou-server/pkg/app"
	"github.com/speanut-land/gdou-server/pkg/e"
	"github.com/speanut-land/gdou-server/pkg/logging"
	"github.com/speanut-land/gdou-server/pkg/redis"
	"github.com/speanut-land/gdou-server/pkg/util"
	"github.com/speanut-land/gdou-server/service/user_service"
	"net/http"
	"strings"
)

type UserForm struct {
	Username  string `form:"username" valid:"Required;MaxSize(50)"`
	Password  string `form:"password" valid:"Required;MaxSize(20)"`
	Telephone string `form:"telephone" valid:"Required;MaxSize(11)"`
	Code      string `form:"code" valid:"Required;MaxSize(6)"`
	Email     string `form:"email" json:"email"`
}

// @Summary 注册用户
// @Tags 用户接口
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param telephone body string true "手机号"
// @Param code body string true "验证码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /user/register [post]
func Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form UserForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, false, nil)
		return
	}
	if redis.Exists("register" + form.Telephone) {
		data, err := redis.Get("register" + form.Telephone)
		temp := string(data)
		tempData := temp[1 : len(temp)-1]
		if err != nil {
			logging.Info(err)
			appG.Response(http.StatusInternalServerError, e.ERROR, false, nil)
			return
		} else if !strings.EqualFold(form.Code, tempData) {
			appG.Response(http.StatusOK, e.ERROR_CODE, false, nil)
			return
		}
	} else {
		appG.Response(http.StatusOK, e.ERROR_CODE, false, nil)
		return
	}
	userService := user_service.User{
		Username:  form.Username,
		Password:  form.Password,
		Telephone: form.Telephone,
	}
	exists, err := userService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USER_FAIL, false, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_USER, false, nil)
		return
	}
	err = userService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, false, nil)
		return
	}

	_, err = redis.Delete("register" + form.Telephone)
	if err != nil {
		logging.Info(err)
	}
	appG.Response(http.StatusOK, e.SUCCESS, true, nil)
}

type UserLoginForm struct {
	Username  string `form:"username" valid:"Required;MaxSize(50)"`
	Password  string `form:"password" valid:"Required;MaxSize(20)"`
	Telephone string `form:"telephone" valid:"MaxSize(11)"`
	Email     string `form:"email" json:"email"`
}

// @Summary 用户登录
// @description 请登录后将token放在请求头上
// @Tags 用户接口
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "用户密码"
// @Success 200 {object} app.Response
// @Router /user/login [post]
func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form UserLoginForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, false, nil)
		return
	}
	userService := user_service.User{
		Username: form.Username,
		Password: form.Password,
	}
	isLogin, err := userService.Login()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USER_FAIL, false, nil)
		return
	}
	if !isLogin {
		appG.Response(http.StatusOK, e.ERROR_LOGIN_FAIL, false, nil)
		return
	}
	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, false, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, true, map[string]string{
		"token": token,
	})
}

type UserResetPasswordForm struct {
	Telephone string `form:"telephone" valid:"Required;MaxSize(11)"`
	Password  string `form:"password" valid:"Required;MaxSize(20)"`
	Code      string `form:"code" valid:"Required;MaxSize(6)"`
}

// @Summary 重置用户密码
// @Tags 用户接口
// @Produce json
// @Param telephone body string true "手机号"
// @Param code body string true "验证码"
// @Param password body string true "用户密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /user/resetPassword [post]
func ResetPassword(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form UserResetPasswordForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, false, nil)
		return
	}
	if redis.Exists("resetPassword" + form.Telephone) {
		data, err := redis.Get("resetPassword" + form.Telephone)
		temp := string(data)
		tempData := temp[1 : len(temp)-1]
		if err != nil {
			logging.Info(err)
			appG.Response(http.StatusInternalServerError, e.ERROR, false, nil)
			return
		} else if !strings.EqualFold(form.Code, tempData) {
			appG.Response(http.StatusOK, e.ERROR_CODE, false, nil)
			return
		}
	} else {
		appG.Response(http.StatusOK, e.ERROR_CODE, false, nil)
		return
	}
	userService := user_service.User{
		Password:  form.Password,
		Telephone: form.Telephone,
	}
	errCode = userService.ExistByTelephone()
	if errCode == e.ERROR {
		appG.Response(http.StatusInternalServerError, errCode, false, nil)
		return
	} else {
		if errCode != e.ERROR_TELEPHONE_USED {
			appG.Response(http.StatusOK, errCode, false, nil)
			return
		}
	}
	err := userService.ResetPassword()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, false, nil)
		return
	}
	_, err = redis.Delete("resetPassword" + form.Telephone)
	if err != nil {
		logging.Info(err)
	}
	appG.Response(http.StatusOK, e.SUCCESS, true, nil)
}
