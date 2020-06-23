package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)


func TestDoubleHashHello(t *testing.T) {
	// Expected taken from https://en.bitcoin.it/wiki/Protocol_documentation#Hashes
	expectedResult, _ := hex.DecodeString("9595c9df90075148eb06860365df33584b75bff782a510c6cd4883a419833d50")
	actualResult := DoubleHash([]byte("hello"))
	if !bytes.Equal(expectedResult, actualResult)  {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}
