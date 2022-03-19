/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 23:12:03
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 23:41:34
 */
package v1

import (
	"weego/app/models/topic"
	"weego/app/policies"
	"weego/app/requests"
	"weego/pkg/auth"
	"weego/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Index(c *gin.Context) {
	topics := topic.All()
	response.Data(c, topics)
}

func (ctrl *TopicsController) Show(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, topicModel)
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		response.Data(c, topicModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// func (ctrl *TopicsController) Delete(c *gin.Context) {

// 	topicModel := topic.Get(c.Param("id"))
// 	if topicModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}

// 	if ok := policies.CanModifyTopic(c, topicModel); !ok {
// 		response.Abort403(c)
// 		return
// 	}

// 	rowsAffected := topicModel.Delete()
// 	if rowsAffected > 0 {
// 		response.Success(c)
// 		return
// 	}

// 	response.Abort500(c, "删除失败，请稍后尝试~")
// }
