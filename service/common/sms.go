package serviceCommon

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	"sinblog.cn/FunAnime-Server/util/logger"
)

var SuccessSend = "Ok"

func SendSms(phone, smsCode, duration string) error {
	credential := common.NewCredential(
		viper.GetString("tencent_api.secret_id"),
		viper.GetString("tencent_api.secret_key"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, err := sms.NewClient(credential, "ap-beijing", cpf)
	if err != nil {
		logger.Error("send_sms_new_client_failed", logger.Fields{"err": err, "client": client, "cpf": cpf, "credential": credential})
		return err
	}

	templateId := viper.GetString("tencent_api.sms_template_id")
	smsSign := viper.GetString("tencent_api.sms_sign")
	sdkAppId := viper.GetString("tencent_api.sms_app_id")
	newPhone := fmt.Sprintf("+86%s", phone)

	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = []*string{&newPhone}
	request.TemplateID = &templateId
	request.Sign = &smsSign
	request.TemplateParamSet = []*string{&smsCode, &duration}
	request.SmsSdkAppid = &sdkAppId

	response, apiErr := client.SendSms(request)
	if apiErr != nil {
		logger.Error("send_sms_failed", logger.Fields{"err": apiErr, "request": request, "response": response, "client": client, "cpf": cpf, "credential": credential})
		return apiErr
	}

	respCode := response.Response.SendStatusSet[0].Code
	if len(response.Response.SendStatusSet) < 0 || *respCode != SuccessSend {
		logger.Error("send_sms_api_error", logger.Fields{"err": apiErr, "request": request, "response": response, "client": client, "cpf": cpf, "credential": credential})
		return errors.New("send_sms_api_error")
	}

	logger.Info("send_sms_response", logger.Fields{"request": request, "response": response, "client": client, "cpf": cpf, "credential": credential})
	return nil
}
