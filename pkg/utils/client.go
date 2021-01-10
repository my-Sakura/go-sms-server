package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type client struct {
	Name         string
	AppCode      string
	TemplateCode string
}

func NewClient() *client {
	var c client

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		panic(fmt.Sprintf("unmarshal error: %v\n", err))
	}

	return &c
}
