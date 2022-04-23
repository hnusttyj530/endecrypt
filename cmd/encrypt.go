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
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"goencrypt/imp"
	"io"
	"os"
	"strings"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "The command of encryption.",
	Run: func(cmd *cobra.Command, args []string) {
		RunEncrypt(options)
	},
}

func init() {
	encryptCmd.Flags().StringVarP(&options.Type, "type", "t", "aescbc", "The type of encryption[aescbc/aesctr/des/3des/sha256hex/sha512hex/ecc/rsa]")
	encryptCmd.Flags().StringVarP(&options.Encryption, "text", "e", "", "The text to be encryted.")
	encryptCmd.Flags().StringVarP(&options.Key, "key", "k", "1234567887654321", "The key to be encryted.")
	encryptCmd.Flags().StringVarP(&options.Public, "public", "p", "public.pem", "The public key file to be encryted.")
	encryptCmd.Flags().StringVarP(&options.Eccpublic, "eccpublic", "c", "eccpublic.pem", "The eccpublic key file to be encryted.")
	rootCmd.AddCommand(encryptCmd)
}

func GetFileStr(fileName string) string {
	result := ""
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Errorf("Error:%s", err.Error())
		return result
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		result = result + line
		if err == io.EOF { //读取结束，会报EOF
			break
		}
	}
	return result
}
func RunEncrypt(op Options) {
	var str string
	plainText := []byte(strings.TrimSpace(op.Encryption))
	key := []byte(op.Key)
	switch op.Type {
	case "aescbc":
		str = imp.GetAesCbcEncrypt(plainText, key, []byte(""))
		break
	case "aesctr":
		str = imp.GetAesCtrEncrypt(plainText, key, []byte(""))
		break
	case "des":
		key = []byte("12345678")
		str = imp.GetDesEncrypt(plainText, key, []byte(""))
		break
	case "3des":
		key = []byte("123456788765432112345678")
		str = imp.GetTripleDesEncrypt(plainText, key, []byte(""))
		break
	case "sha256hex":
		str = imp.GetSha256Hex(plainText)
		break
	case "sha512hex":
		str = imp.GetSha512Hex(plainText)
		break
	case "ecc":
		if op.Encryption == "" {
			imp.PrintEccPrivateAndPublic()
		} else {
			public := []byte(GetFileStr(op.Eccpublic))
			str = imp.GetEccEncrypt(plainText, public)
		}
		break
	case "rsa":
		if op.Encryption == "" {
			imp.PrintRSAPrivateAndPublic()
		} else {
			public := []byte(GetFileStr(op.Public))
			str = imp.GetRSAEncrypt(plainText, public)
		}
		break
	default:
		break
	}
	fmt.Println(str)
}
