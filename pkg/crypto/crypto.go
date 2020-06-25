package crypto

import "crypto/sha256"

// DoubleSha returns sha256(sha256(message))
func DoubleSha(message []byte) []byte {
	hash1 := sha256.New()
	hash1.Write(message)
	hash2 := sha256.New()
	hash2.Write(hash1.Sum(nil))
	return hash2.Sum(nil)
}
