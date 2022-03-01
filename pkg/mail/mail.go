/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-01 18:36:53
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-01 19:11:54
 */
package mail

import (
	"sync"
	"weego/pkg/config"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

// NewMailer 单例模式获取
func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})

	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, config.GetStringMapString("mail.smtp"))
}
