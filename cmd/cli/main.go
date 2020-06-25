package main

import (
	"bitcoin-go/pkg/config"
	"bitcoin-go/pkg/networking"
	"encoding/hex"
	"fmt"
)

func main() {
	// fmt.Print("Hello, world!")
	command := "verack"
	payload := []byte("")
	conf := config.GenerateDefaultConfig()
	fmt.Print(hex.EncodeToString(networking.CreateHeader(command, payload, conf)))
}
