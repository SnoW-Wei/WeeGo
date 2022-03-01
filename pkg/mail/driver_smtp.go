/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-01 13:16:45
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-01 18:36:00
 */
package mail

import (
	"fmt"
	"net/smtp"
	"weego/pkg/logger"

	emailPKG "github.com/jordan-wright/email"
)

type SMTP struct{}

func (s *SMTP) Send(email Email, config map[string]string) bool {

	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v%>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发件详情", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),

		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)

	if err != nil {
		logger.DebugString("发送邮件", "发件出错", err.Error())
		return false
	}

	logger.DebugString("发送邮件", "发件成功", "")
	return true
}
