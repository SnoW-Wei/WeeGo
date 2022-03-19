/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 16:44:12
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 16:48:19
 */
package config

import "weego/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {

		return map[string]interface{}{
			// 默认每页条数
			"perpage": 10,
			// URL中泳衣分辨多少页的参数
			// 此值若修改需一并修改请求验证规则
			"url_query_page": "page",
			// URL 中用以分辨排序的参数（使用 id 或者其他）
			// 此值若修改需一并修改请求验证规则
			"url_query_sort": "sort",
			// URL 中用以分辨排序规则的参数（辨别时正序还是倒序）
			// 此值若修改需一并修改请求验证规则
			"url_query_order": "order",
			// URL 中用以分辨每页条数的参数
			//此值若修改需一并修改请求验证规则
			"url_query_per_page": "per_page",
		}
	})
}
