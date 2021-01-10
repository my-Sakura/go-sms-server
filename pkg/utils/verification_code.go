package utils

import (
	"math/rand"
	"strconv"
	"time"
)

//GetCode generate rand verification code
func GetCode(length int) string {
	var code string
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}
