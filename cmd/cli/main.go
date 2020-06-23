package main

import (
	"bitcoin-go/pkg/networking"
	"encoding/hex"
	"fmt"
)

func main() {
	command := "verack"
	payload := ""
	fmt.Print(hex.EncodeToString(networking.CreateMessage(command, payload)))
}
