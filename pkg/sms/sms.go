package sms

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/speanut-land/gdou-server/pkg/logging"
	"github.com/speanut-land/gdou-server/pkg/setting"
)

func SendSms(telephone string, code string) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", setting.AliAccessSetting.AccessKeyId,
		setting.AliAccessSetting.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = telephone
	request.SignName = setting.AliAccessSetting.SignName
	request.TemplateCode = setting.AliAccessSetting.TemplateCode
	request.TemplateParam = fmt.Sprintf("{'code':%s}", code)

	_, err = client.SendSms(request)
	if err != nil {
		logging.Info(err.Error())
	}
}
