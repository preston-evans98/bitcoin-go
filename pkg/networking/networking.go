package networking

import (
	"bitcoin-go/pkg/config"
	"bitcoin-go/pkg/constants"
	"bitcoin-go/pkg/crypto"
	"encoding/binary"
)

// Service is a type for Network Service identifiers
// used in version messsages
type Service uint64

// Exported service types
const (
	Unnamed            Service = 0x00
	NodeNetwork        Service = 0x01
	NodeGetUtxo        Service = 0x02
	NodeBloom          Service = 0x04
	NodeWitness        Service = 0x08
	NodeXthin          Service = 0x10
	NodeNetworkLimited Service = 0x400
)

// Uint64 converts a service to a uint64
func (s Service) Uint64() uint64 {
	return uint64(s)
}

func main() {
	// conn, err := net.Dial("tcp", "192.168.1.3:8333")
	// if err != nil {
	//     log.Fatal("Could not connect to pi")
	// }

	// conn.write("")

}

// CreateHeader transforms creates a message header for a given command and payload
func CreateHeader(command string, payload []byte, config *config.NodeConfig) []byte {

	header := make([]byte, 24)

	// https://developer.bitcoin.org/glossary.html#term-Start-string
	binary.BigEndian.PutUint32(header, networkToStartString(config.GetNetwork()))
	copy(header[4:16], []byte(command))
	binary.BigEndian.PutUint32(header[16:20], uint32(len(payload)))
	checksum := crypto.DoubleSha(payload)

	copy(header[20:], checksum)
	return header
}

// // CreateVersionBody creates the body of a "version header"
// func CreateVersionBody(supportedServices []Service, conf *config.NodeConfig) {

// 	// Do Setup
// 	message := make([]byte, 85)
// 	// Or supportedServices to construct service field
// 	var services Service = 0x0000
// 	for _, service := range supportedServices {
// 		services |= service
// 	}

// 	// TODO: implement user agent
// 	// https://github.com/bitcoin/bips/blob/master/bip-0014.mediawiki

// 	// Construct Message
// 	// Version (4 bytes)
// 	binary.BigEndian.PutUint32(message, constants.ProtocolVersion)
// 	// Services (8 bytes)
// 	binary.BigEndian.PutUint64(message[4:12], services.Uint64())
// 	// Timestamp (8 bytes)
// 	binary.BigEndian.PutUint64(message[12:20], uint64(time.Now().Unix()))
// 	// Receiver Services (8 bytes)
// 	// TODO: make accurate. Although BitcoinJ has same hardcoded behavior
// 	// No-op

// 	// Receiver IP address (16 bytes)
// 	// TODO: make accuarate. Again, BitcoinJ doesn't
// 	copy(message[20:36], []byte("::ffff:127.0.0.1"))

// 	// Port (2 bytes)
// 	binary.BigEndian.PutUint16(message[36:38], networkToPort(conf.GetNetwork()))

// }

func networkToPort(network int) uint16 {
	if network == constants.Mainnet {
		return constants.MainnetDefaultPort
	} else if network == constants.Testnet {
		return constants.TestnetDefaultPort
	} else if network == constants.Regtest {
		return constants.RegtestDefaultPort
	} else {
		panic("Unkown network identifier: %v")
	}
}

func networkToStartString(network int) uint32 {
	if network == constants.Mainnet {
		return constants.MainnetStartString
	} else if network == constants.Testnet {
		return constants.TestnetStartString
	} else if network == constants.Regtest {
		return constants.RegtestStartString
	} else {
		panic("Unkown network identifier")
	}
}
