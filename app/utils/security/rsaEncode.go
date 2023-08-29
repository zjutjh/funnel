package security

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

type RSAPublicKey struct {
	Modulus  string `json:"modulus"`
	Exponent string `json:"exponent"`
}

func GetEncodePassword(publickey []byte, password []byte) (string, error) {
	K := &RSAPublicKey{}
	err := json.Unmarshal(publickey, K)
	if err != nil {
		return "", err
	}

	nString, _ := base64.StdEncoding.DecodeString(K.Modulus)
	n, _ := new(big.Int).SetString(hex.EncodeToString(nString), 16)
	eString, _ := base64.StdEncoding.DecodeString(K.Exponent)
	e, _ := strconv.ParseInt(hex.EncodeToString(eString), 16, 32)
	pub := rsa.PublicKey{E: int(e), N: n}
	cc, err := rsa.EncryptPKCS1v15(rand.Reader, &pub, password)
	return base64.StdEncoding.EncodeToString(cc), err
}

func GetEncryptPassword(publicKey []byte, password string) (string, error) {
	K := &RSAPublicKey{}
	err := json.Unmarshal(publicKey, K)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	e, _ := strconv.ParseInt(K.Exponent, 16, 32)
	pub := publicKeyFromHex(K.Modulus, int(e))

	encode := encryptString(password, pub)
	return encode, err
}

func publicKeyFromHex(modulusHexString string, exponent int) *rsa.PublicKey {
	modulus := new(big.Int)
	modulus.SetString(modulusHexString, 16)

	return &rsa.PublicKey{
		N: modulus,
		E: exponent,
	}
}

func encryptString(msg string, pub *rsa.PublicKey) string {
	result := encrypt([]byte(msg), pub)
	hexString := hex.EncodeToString(result[0])
	for i := 1; i < len(result); i++ {
		hexString += " " + hex.EncodeToString(result[i])
	}
	return hexString
}

func chunkSize(pub *rsa.PublicKey) int {
	return 2 * (pub.N.BitLen()/16 - 1)
}

func max(l int, r int) int {
	if l > r {
		return l
	}
	return r
}

// encrypt encrypts msg using pub and returns the ciphertext as a big integer.
// The length of msg must be smaller than pub.N.
func encrypt(msg []byte, pub *rsa.PublicKey) [][]byte {
	chunkSize := chunkSize(pub)
	// paddedMsg := pad(msg, chunkSize)
	msgLen := len(msg)
	var result [][]byte

	// make compatible with big-endian byte order
	// reverse(paddedMsg)

	// encrypt each chunk
	for i := msgLen; i > 0; i -= chunkSize {
		// convert the bytes to a big integer
		msgChunk := new(big.Int)
		msgChunk.SetBytes(msg[max(i-chunkSize, 0):i])

		msgChunk.Exp(msgChunk, big.NewInt(int64(pub.E)), pub.N) // cipher = msg^e mod n

		// append the chunk to the result
		result = append(result, msgChunk.Bytes())
	}

	return result
}
