package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/utils/security"
	"github.com/go-resty/resty/v2"
)

// OauthLoginGetPublicKey 获取加密密钥
func OauthLoginGetPublicKey(client *resty.Client) ([]byte, error) {
	resp, err := client.R().
		EnableTrace().
		Get(library.OAuthPublicKey)
	if err != nil {
		return []byte{}, err
	}
	return resp.Body(), nil
}

// GetEncryptedPwd 密码加密
func GetEncryptedPwd(client *resty.Client, password string) (string, error) {
	key, err := OauthLoginGetPublicKey(client)
	if err != nil {
		return "", err
	}

	encrypted, err := security.GetEncryptPassword(key, password)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}
