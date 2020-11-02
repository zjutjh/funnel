package helps

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

func GetEncodePassword(s []byte, password []byte) string {
	p := &RSAPublicKey{}
	_ = json.Unmarshal(s, p)
	nString, _ := base64.StdEncoding.DecodeString(p.Modulus)
	n, _ := new(big.Int).SetString(hex.EncodeToString(nString), 16)
	eString, _ := base64.StdEncoding.DecodeString(p.Exponent)
	e, _ := strconv.ParseInt(hex.EncodeToString(eString), 16, 32)
	pub := rsa.PublicKey{E: int(e), N: n}
	cc, _ := rsa.EncryptPKCS1v15(rand.Reader, &pub, password)
	return base64.StdEncoding.EncodeToString(cc)
}
