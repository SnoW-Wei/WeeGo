/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 15:48:02
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 14:20:59
 */
package routes

import (
	controllers "weego/app/http/controllers/api/v1"
	"weego/app/http/controllers/api/v1/auth"
	"weego/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRouter 注册API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// V1 路由组，我们所有 v1 版本的路由都将存放在这里
	v1 := r.Group("v1")

	// 全局限流中间件：每小时限流，这里是所有API（根据IP）请求加起来
	// 作为参考 GITHUB API 每小时最多 60个请求（根据IP）

	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")
		// 限流中间键：每小时限流，作为参考github api 每小时最多 60个请求（根据IP）
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 注册用户
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-code/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-code/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-code/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 使用手机号，Email和用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			//重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByphone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)

		}
		uc := new(controllers.UsersController)

		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)

		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", middlewares.AuthJWT(), uc.UpdateProfile)
			usersGroup.PUT("/email", middlewares.AuthJWT(), uc.UpdateEmail)
		}

		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			cgcGroup.GET("", cgc.Index)
			cgcGroup.GET("/:id", cgc.Show)
			cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.GET("", tpc.Index)
			tpcGroup.GET("/:id", tpc.Show)
			tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
		}

		lsc := new(controllers.LinksController)
		linksGroup := v1.Group("/links")
		{
			linksGroup.GET("", lsc.Index)
		}

	}

}
