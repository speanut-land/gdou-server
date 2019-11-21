package sendcode

import (
	"github.com/gin-gonic/gin"
	"github.com/speanut-land/gdou-server/models"
	"github.com/speanut-land/gdou-server/pkg/app"
	"github.com/speanut-land/gdou-server/pkg/e"
	"github.com/speanut-land/gdou-server/pkg/redis"
	"github.com/speanut-land/gdou-server/pkg/sms"
	"github.com/speanut-land/gdou-server/pkg/util"
	"net/http"
)

// @Summary 发送注册短信验证码
// @Tags 短信接口
// @Produce json
// @Param telephone body string true "Telephone"
//@Success 200 {object} app.Response
//@Failure 500 {object} app.Response
// @Router /sendCode/register [post]
func Register(c *gin.Context) {
	appG := app.Gin{C: c}
	telephone := c.PostForm("telephone")
	errCode := models.IsTelephoneUsable(telephone)
	if errCode == e.ERROR {
		appG.Response(http.StatusInternalServerError, errCode, false, nil)
		return
	} else {
		if errCode != e.ERROR_TELEPHONE_UNREGISTER {
			appG.Response(http.StatusOK, errCode, false, nil)
			return
		}
	}
	code := util.CreateCaptcha()
	_ = redis.Set("register"+telephone, code, 60*5)
	defer sms.SendSms(telephone, code)
	appG.Response(http.StatusOK, e.SUCCESS, true, nil)
}

// @Summary 发送重置密码的短信验证码
// @Tags 短信接口
// @Produce json
// @Param telephone body string true "Telephone"
//@Success 200 {object} app.Response
//@Failure 500 {object} app.Response
// @Router /sendCode/resetPassword [post]
func ResetPassword(c *gin.Context) {
	appG := app.Gin{C: c}
	telephone := c.PostForm("telephone")
	errCode := models.IsTelephoneUsable(telephone)
	if errCode == e.ERROR {
		appG.Response(http.StatusInternalServerError, errCode, false, nil)
		return
	} else {
		if errCode != e.ERROR_TELEPHONE_USED {
			appG.Response(http.StatusOK, errCode, false, nil)
			return
		}
	}
	code := util.CreateCaptcha()
	println("resetPassword"+telephone, code)
	_ = redis.Set("resetPassword"+telephone, code, 60*5)
	defer sms.SendSms(telephone, code)
	appG.Response(http.StatusOK, e.SUCCESS, true, nil)
}
