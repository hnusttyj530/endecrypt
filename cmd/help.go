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
	"os"
	"os/exec"
)

func ExecuteCommand(name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)
	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()
	return string(bytes), err
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}
