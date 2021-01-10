package main

import (
	"github.com/my-Sakura/SMS/api"
)

func main() {
	client := api.NewClient()
	client.Send("17731895913", 6)
}
