package networking

import (
	"bitcoin-go/pkg/crypto"
	"encoding/binary"
)

func main() {
	// conn, err := net.Dial("tcp", "192.168.1.3:8333")
	// if err != nil {
	//     log.Fatal("Could not connect to pi")
	// }

	// conn.write("")

}

// CreateHeader transforms creates a message header for a given command and payload
func CreateHeader(command string, payload string) []byte {

	header := make([]byte, 24)
	bytePayload := []byte(payload)

	// https://developer.bitcoin.org/glossary.html#term-Start-string
	var startString uint32 = 0xf9beb4d9
	binary.BigEndian.PutUint32(header, startString)
	copy(header[4:16], []byte(command))
	binary.BigEndian.PutUint32(header[16:20], uint32(len(payload)))
	checksum := crypto.DoubleHash(bytePayload)

	copy(header[20:], checksum)
	return header
}
