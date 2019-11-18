package api

import (
	"github.com/gin-gonic/gin"
	"github.com/speanut-land/gdou-server/pkg/app"
	"github.com/speanut-land/gdou-server/pkg/e"
	"github.com/speanut-land/gdou-server/service/user_service"
	"net/http"
)

type UserForm struct {
	Username  string `form:"username" valid:"Required;MaxSize(50)"`
	Password  string `form:"password" valid:"Required;MaxSize(20)"`
	Telephone string `form:"telephone" valid:"Required;MaxSize(11)"`
	Email     string `form:"email" json:"email"`
}

// @Summary 注册用户
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Param telephone body string true "Telephone"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /register [post]
func Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form UserForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	userService := user_service.User{
		Username:  form.Username,
		Password:  form.Password,
		Telephone: form.Telephone,
	}
	exists, err := userService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USER_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_USER, nil)
		return
	}
	err = userService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
