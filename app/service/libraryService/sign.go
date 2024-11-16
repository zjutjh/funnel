package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/service/libraryService/request"
	"funnel/app/utils/security"
)

// OauthLoginGetPublicKey 获取加密密钥
func OauthLoginGetPublicKey(client request.Client) ([]byte, error) {
	resp, err := client.Request().
		Get(library.OAuthPublicKey)
	if err != nil {
		return []byte{}, err
	}
	return resp.Body(), nil
}

// GetEncryptedPwd 密码加密
func GetEncryptedPwd(client request.Client, password string) (string, error) {
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
