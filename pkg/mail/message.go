package mail

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/spf13/viper"
)

type MessageTemplateParam struct {
	Code string `json:"code"`
}

func SendMessageCode(code string, recipient string) error {
	regionId := viper.GetString("aliyun_regin_id")
	accessKeyId := viper.GetString("aliyun_access_key_id")
	accessKeySecret := viper.GetString("aliyun_access_key_secret")
	signName := viper.GetString("aliyun_sign_name")
	tempCode := viper.GetString("aliyun_temp_code")

	client, err := dysmsapi.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = recipient
	request.SignName = signName
	request.TemplateCode = tempCode
	par := MessageTemplateParam{code}
	b, err := json.Marshal(par)
	if err != nil {
		return err
	}
	request.TemplateParam = string(b)

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}
