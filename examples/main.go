package main

import (
	"log"

	"github.com/my-Sakura/go-sms-server/api"
)

//smsProvider: tianyan
//appCode: xxx
//templateCode:
func main() {
	client := api.NewClient("smsProvider", "your_templateCode", "your_appCode")
	err := client.Send("15032371660", 6)
	if err != nil {
		log.Println(err)
	}
}
