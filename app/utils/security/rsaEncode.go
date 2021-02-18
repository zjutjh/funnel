package security

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
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
