/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 23:45:30
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 23:48:33
 */
package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file, example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	os.MkdirAll("app/policies", os.ModePerm)

	filePath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)

	createFileFromStub(filePath, "policy", model)
}
