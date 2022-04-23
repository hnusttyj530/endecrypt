/**
 @author:tyj
 @date:2022/4/1
 @note:1670171244@qq.com
 @TODO:
 @Param:
 @return:
**/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/*!
flags option
*/
type Options struct {
	//Input      string
	//Output     string
	Type       string
	Version    string
	Encryption string
	Decryption string
	Key        string
	Public     string
	Private    string
	Eccpublic  string
	Eccprivate string
}

var options Options
var rootCmd = &cobra.Command{
	Use:   "endecrypt",
	Short: "endecrypt",
	Long:  `The tool for encryption & decryption with common algorithm.`,
	//main enterance
	Run: func(cmd *cobra.Command, args []string) {
		//Convert(options)
		fmt.Println("hello world!")
	},
}

func init() {
	rootCmd.Flags().StringVarP(&options.Version, "version", "V", "1.0", "The version of endecrypt.")
}

func Execute() {
	rootCmd.Execute()
}
