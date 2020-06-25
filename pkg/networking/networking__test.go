package networking

import (
	"bitcoin-go/pkg/config"
	"bytes"
	"encoding/hex"
	"testing"
)

func TestCreateHeaderVerack(t *testing.T) {
	// Expected taken from https://developer.bitcoin.org/reference/p2p_networking.html#version
	expectedResult, err := hex.DecodeString("f9beb4d976657261636b000000000000000000005df6e0e2")
	if err != nil {
		t.Fatalf("Could not decode expectedResult: %v", err)
	}
	conf := config.GenerateDefaultConfig()
	actualMessage := CreateHeader("verack", []byte(""), conf)

	if !bytes.Equal(expectedResult, actualMessage) {
		t.Fatalf("Expected %v but got %v", expectedResult, actualMessage)
	}
}

func TestCreateVersion(t *testing.T) {

}

// func TestCreateVersion(t *testing.T) {
// 	expectedResult, err := hex.DecodeString("7111010001d39f608a7775b537729884d4e6633bb2105e55a16a14d31b00000000000000000000000000000000000000000000000000000000000000000000000000000000")
// 	if err != nil {
// 		t.Fatalf("Could not decode expectedResult: %v", err)
// 	}
// 	actualMessage := CreateVersionMessage("d39f608a7775b537729884d4e6633bb2105e55a16a14d31b0000000000000000")

// }
