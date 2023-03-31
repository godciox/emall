package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"math/rand"
	"time"
	cache "userservice/db/redisInit"
)

const (
	accessKey    = "LTAI5tCLKvo5pBgQNpD6BCDC"
	accessSecret = "TXb98JSbEFBG5htThYcqsuWjjDsRt8"
)

func SendMessage(phoneNumber string) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-shenzhen", accessKey, accessSecret)
	/* use STS Token
	client, err := dysmsapi.NewClientWithStsToken("cn-qingdao", "<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phoneNumber     //接收短信的手机号码
	request.SignName = "阿里云短信测试"    //短信签名名称
	request.TemplateCode = "SMS_154950909" //短信模板ID
	var randCode string = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	request.TemplateParam = "{\"code\":\"" + randCode + "\"}"
	response, err := client.SendSms(request)

	if err != nil {
		fmt.Print(err.Error())
	}
	if response.Code == "OK" {
		cache.SetCaptcha(phoneNumber, randCode)
	}
	fmt.Printf("response is %#v\n", response)
}
