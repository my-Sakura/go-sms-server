package main

import (
	"log"

	"github.com/my-Sakura/go-sms-server/api"
)

//smsProvider: tianyan
//appCode: xxx
//templateCode:
func main() {
	client := api.NewClient("dingxin", "4c576b3341754ec2a8db38e6f70c1da9", "TP1711063")
	err := client.Send("15032371660", 6)
	if err != nil {
		log.Println(err)
	}
}

//your_templateCode
//your_appCode
