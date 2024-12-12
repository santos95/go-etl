package config

import (
	"fmt"
	b64 "encoding/base64"
)

func DecodePassString(text string) string {

	fmt.Println("Decoding String ... ")
	desStr, _ := b64.StdEncoding.DecodeString(text)

	return string(desStr)
}