package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	result, _ := base64.StdEncoding.DecodeString(secret)

	whatIsIt = Reverse(string(result))
	fmt.Println(whatIsIt)
}

func Reverse(str string) (result string) {
	for _,v := range str {
	  result = string(v) + result
	}
	return 
}
