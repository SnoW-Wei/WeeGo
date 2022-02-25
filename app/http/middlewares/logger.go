/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-25 13:24:28
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-25 13:51:38
 */
package middlewares

import (
	"bytes"
	"io/ioutil"
	"time"

	"weego/pkg/helpers"
	"weego/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {

	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer对象，只能读取一次
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body , 以供后续的其他操作
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		c.Next()

		// 开始记录日志的逻辑
		cost := time.Since(start)
		responStatus := c.Writer.Status()

		logFiles := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {

			// 请求内容
			logFiles = append(logFiles, zap.String("Request Body", string(requestBody)))

			// 响应的内容
			logFiles = append(logFiles, zap.String("Request Body", w.body.String()))
		}

		if responStatus > 400 && responStatus <= 499 {
			// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404 ，开发时都要注意
			logger.Warn("HTTP Warning"+cast.ToString(responStatus), logFiles...)
		} else if responStatus >= 500 && responStatus <= 599 {
			// 除了内部错误，记录error
			logger.Error("HTTP Error"+cast.ToString(responStatus), logFiles...)
		} else {
			logger.Debug("HTTP Access Log", logFiles...)
		}
	}
}
