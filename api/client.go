package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/my-Sakura/SMS/pkg/utils"
	"github.com/spf13/viper"
)

type client struct {
	Name         string `yaml: "name"`
	AppCode      string `yaml: "appCode"`
	TemplateCode string `yaml: "templateCode"`
}

func NewClient() *client {
	var c client

	viper.SetConfigFile("../config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Sprintf("unmarshal error: %v\n", err))
	}

	return &c
}

//Send message to mobile
//length is the length of verification code
func (c *client) Send(mobile string, length int) error {

	switch c.Name {
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
