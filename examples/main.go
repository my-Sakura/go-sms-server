package main

import (
	"log"

	"github.com/my-Sakura/go-sms-server/api"
)

func main() {
	client := api.NewClient("tianyan", "your_appCode", "your_templateCode")
	err := client.Send("mobile", 6)
	if err != nil {
		log.Println(err)
	}
}
