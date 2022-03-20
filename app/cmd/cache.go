/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-20 10:56:39
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 10:59:03
 */
package cmd

import (
	"weego/pkg/cache"
	"weego/pkg/console"

	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "clear cache",
	Run:   runCacheClear,
}

func init() {
	CmdCache.AddCommand(CmdCacheClear)
}

func runCacheClear(cmd *cobra.Command, args []string) {

	cache.Flush()
	console.Success("cache cleard.")
}
