package util

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"
	"regexp"
)

func ExtractJsonStr(jsonStr, jsonKey string) string {
	re := regexp.MustCompile(jsonKey + `:'([^']+)'`)
	matches := re.FindStringSubmatch(jsonStr)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// RSAEncryptWithPublicKey 使用模数和指数加密数据。
func RSAEncryptWithPublicKey(plaintext, modulusB64, exponentB64 string) (string, error) {
	if plaintext == "" || modulusB64 == "" || exponentB64 == "" {
		return "", fmt.Errorf("plaintext, modulusB64 and exponentB64 cannot be empty")
	}

	modulusBytes, err := base64.StdEncoding.DecodeString(modulusB64)
	if err != nil {
		return "", fmt.Errorf("modulus base64 解码失败: %v", err)
	}

	exponentBytes, err := base64.StdEncoding.DecodeString(exponentB64)
	if err != nil {
		return "", fmt.Errorf("exponent base64 解码失败: %v", err)
	}

	pubKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(modulusBytes),
		E: int(new(big.Int).SetBytes(exponentBytes).Int64()),
	}

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(plaintext))
	if err != nil {
		return "", fmt.Errorf("rsa.EncryptPKCS1v15 加密失败: %v", err)
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}
