package networking

import (
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
	actualMessage := CreateHeader("verack", "")
	
	if !bytes.Equal(expectedResult, actualMessage)  {
		t.Fatalf("Expected %v but got %v", expectedResult, actualMessage)
	}
}
