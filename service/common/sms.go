package serviceCommon

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func SendSms(phone, smsCode, duration string) error {
	credential := common.NewCredential(
		viper.GetString("tencent_api.secret_id"),
		viper.GetString("tencent_api.secret_key"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, err := sms.NewClient(credential, "ap-beijing", cpf)
	if err != nil {
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

	response, err := client.SendSms(request)
	if err != nil {
		return err
	}

	fmt.Printf("%s", response.ToJsonString())
	return nil
}
