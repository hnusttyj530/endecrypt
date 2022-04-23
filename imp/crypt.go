/**
  @author:tyj
  @date:2022/4/21
  @note:1670171244@qq.com
  @TODO:
  @Param:
  @return:
**/
package imp

import (
	"encoding/base64"
	"fmt"
	"github.com/wumansgy/goEncrypt"
)

func GetSha256Hex(data []byte) string {
	return goEncrypt.Sha256Hex(data)
}

func GetSha512Hex(data []byte) string {
	return goEncrypt.Sha512Hex(data)
}

func GetAesCtrEncrypt(text, key, data []byte) string {
	cryptText, err := goEncrypt.AesCtrEncrypt(text, key, data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(cryptText)
}
func GetAesCtrDecrypt(text, key, data []byte) string {
	cryptText, err := goEncrypt.AesCtrDecrypt(text, key, data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(cryptText)
}

func GetAesCbcEncrypt(text, key, data []byte) string {
	cryptText, err := goEncrypt.AesCbcEncrypt(text, key, data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(cryptText)
}

func GetAesCbcDecrypt(text, key, data []byte) string {
	cryptText, err := goEncrypt.AesCbcDecrypt(text, key, data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(cryptText)
}

func GetTripleDesEncrypt(text, key, data []byte) string {
	cryptText, err := goEncrypt.TripleDesEncrypt(text, key, data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(cryptText)
}

func GetTripleDesDecrypt(text, key, data []byte) string {
	cryptText, err := goEncrypt.TripleDesDecrypt(text, key, data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(cryptText)
}

func GetDesEncrypt(text, key, data []byte) string {
	cryptText, err := goEncrypt.DesCbcEncrypt(text, key, data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(cryptText)
}

func GetDesDecrypt(text, key, data []byte) string {
	cryptText, err := goEncrypt.DesCbcDecrypt(text, key, data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(cryptText)
}
