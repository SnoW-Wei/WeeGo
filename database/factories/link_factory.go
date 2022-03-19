/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-20 00:19:40
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 00:19:49
 */
package factories

import (
	"weego/app/models/link"

	"github.com/bxcodec/faker/v3"
)

func MakeLinks(count int) []link.Link {

	var objs []link.Link

	// 设置唯一性，如 Link 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		linkModel := link.Link{
			Name: faker.Username(),
			URL:  faker.URL(),
		}
		objs = append(objs, linkModel)
	}

	return objs
}
