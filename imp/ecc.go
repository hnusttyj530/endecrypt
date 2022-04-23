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
	"encoding/hex"
	"fmt"
	"github.com/wumansgy/goEncrypt"
)

func PrintEccPrivateAndPublic() {
	goEncrypt.GetEccKey()
}
func getEccSign(msg, private []byte) string {
	rtext, stext, err := goEncrypt.EccSign(msg, private)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return hex.EncodeToString(rtext) + hex.EncodeToString(stext)
}
func IsEccSign(msg, public, rtext, stext []byte) bool {
	return goEncrypt.EccVerifySign(msg, public, rtext, stext)
}

func GetEccEncrypt(text, public []byte) string {
	cryptText, err := goEncrypt.EccEncrypt(text, public)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return hex.EncodeToString(cryptText)
}

func GetEccDecrypt(text, private []byte) string {
	cryptText, err := goEncrypt.EccDecrypt(text, private)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(cryptText)
}
