/**
  @author:tyj
  @date:2022/4/23
  @note:1670171244@qq.com
  @TODO:
  @Param:
  @return:
**/
package cmd

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"goencrypt/imp"
	"strings"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "The command of decryption.",
	Run: func(cmd *cobra.Command, args []string) {
		RunDecrypt(options)
	},
}

func init() {
	decryptCmd.Flags().StringVarP(&options.Type, "type", "t", "aescbc", "The type of decryption[aescbc/aesctr/des/3des/ecc/rsa]")
	decryptCmd.Flags().StringVarP(&options.Decryption, "password", "d", "", "The text to be decryted.")
	decryptCmd.Flags().StringVarP(&options.Key, "key", "k", "1234567887654321", "The key to be decryted.")
	decryptCmd.Flags().StringVarP(&options.Private, "private", "p", "private.pem", "The private key file to be decryted.")
	decryptCmd.Flags().StringVarP(&options.Eccprivate, "eccprivate", "C", "eccprivate.pem", "The eccprivate key file to be decryted.")
	rootCmd.AddCommand(decryptCmd)
}

func RunDecrypt(op Options) {
	var str string
	key := []byte(op.Key)
	switch op.Type {
	case "aescbc":
		password, _ := base64.StdEncoding.DecodeString(strings.TrimSpace(op.Decryption))
		str = imp.GetAesCbcDecrypt(password, key, []byte(""))
		break
	case "aesctr":
		password, _ := base64.StdEncoding.DecodeString(strings.TrimSpace(op.Decryption))
		str = imp.GetAesCtrDecrypt(password, key, []byte(""))
		break
	case "des":
		key = []byte("12345678")
		password, _ := base64.StdEncoding.DecodeString(strings.TrimSpace(op.Decryption))
		str = imp.GetDesDecrypt(password, key, []byte(""))
		break
	case "3des":
		password, _ := base64.StdEncoding.DecodeString(strings.TrimSpace(op.Decryption))
		key = []byte("123456788765432112345678")
		str = imp.GetTripleDesDecrypt(password, key, []byte(""))
		break
	case "ecc":
		password, _ := hex.DecodeString(strings.TrimSpace(op.Decryption))
		if op.Decryption == "" {
			imp.PrintEccPrivateAndPublic()
		} else {
			private := []byte(GetFileStr(op.Eccprivate))
			str = imp.GetEccDecrypt(password, private)
		}
		break
	case "rsa":
		password, _ := hex.DecodeString(strings.TrimSpace(op.Decryption))
		if op.Decryption == "" {
			imp.PrintRSAPrivateAndPublic()
		} else {
			private := []byte(GetFileStr(op.Private))
			str = imp.GetRSADecrypt(password, private)
		}
		break
	default:
		break
	}
	fmt.Println(str)
}
