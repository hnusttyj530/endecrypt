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

const version = "1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print endecrypt version info.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("endecrypt's version is %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
