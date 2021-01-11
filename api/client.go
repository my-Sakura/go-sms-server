package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/my-Sakura/go-sms-server/pkg/utils"
)

type client struct {
	SMSProvider  string
	AppCode      string
	TemplateCode string
}

//New a client
func NewClient(smsProvider, appCode, templateCode string) *client {
	return &client{
		SMSProvider:  smsProvider,
		AppCode:      appCode,
		TemplateCode: templateCode,
	}
}

//Send message to mobile
//length is the length of verification code
func (c *client) Send(mobile string, length int) error {

	switch c.SMSProvider {
	case "dingxin":
		code := utils.GetCode(length)
		url := fmt.Sprintf("http://dingxin.market.alicloudapi.com/dx/sendSms?mobile=%s&param=code:%s&tpl_id=%s", mobile, code, c.TemplateCode)

		client := &http.Client{}

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "APPCODE"+" "+c.AppCode)
		resp, _ := client.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response is %v\n", string(b))

		return nil
	case "tianyan":
		code := utils.GetCode(length)
		url := fmt.Sprintf("https://smssend.shumaidata.com/sms/send?receive=%s&tag=%s&templateId=%s", mobile, code, c.TemplateCode)

		client := &http.Client{}

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "APPCODE"+" "+c.AppCode)
		resp, _ := client.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response is %v\n", string(b))

		return nil
	}

	return nil
}
