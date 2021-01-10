package main

import (
	"github.com/my-Sakura/go-sms-server/api"
)

func main() {
	client := api.NewClient()
	client.Send("17731895913", 6)
}
