/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 23:54:46
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 23:55:57
 */
package factories

import (
	"weego/app/models/topic"

	"github.com/bxcodec/faker/v3"
)

func MakeTopics(count int) []topic.Topic {

	var objs []topic.Topic

	// 设置唯一性，如 Topic 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: "3",
			UserID:     "1",
		}
		objs = append(objs, topicModel)
	}

	return objs
}
