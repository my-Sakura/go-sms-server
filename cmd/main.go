package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Dispatcher interface {
	Send(mobile, code string) error
}

type client struct {
	Name         string
	AppCode      string
	TemplateCode string
}

func NewClient() *client {
	var c client

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config/")
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

func (c *client) Send(mobile string) error {

	switch c.Name {
	case "tianyan":
		code := GetCode(6)
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

func GetCode(length int) string {
	var code string
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

func main() {
	c := NewClient()

	err := c.Send("17731895913")
	if err != nil {
		fmt.Println(err)
	}
}
