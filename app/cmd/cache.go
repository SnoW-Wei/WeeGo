/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-20 10:56:39
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 12:27:41
 */
package cmd

import (
	"fmt"
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

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete Redis key , example: cache forget cache-key",
	Run:   runCacheForget,
}

// forget 命令的选项
var cacheKey string

func init() {
	// 注册cache 命令的子命令
	CmdCache.AddCommand(CmdCacheClear, CmdCacheForget)

	// 设置 cache forget 命令的选项
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	CmdCacheForget.MarkFlagRequired("key")
}

func runCacheClear(cmd *cobra.Command, args []string) {

	cache.Flush()
	console.Success("cache cleard.")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("cache key [%s] deleted", cacheKey))
}
