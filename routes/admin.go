/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 15:09:02
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 18:48:46
 */
package routes

import (
	controllers "weego/app/http/controllers/admin/v1"
	"weego/app/http/controllers/admin/v1/auth"
	"weego/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRouter 注册API 相关路由
func RegisterAdminRoutes(r *gin.Engine) {

	// api 路由组，我们所有 v1 版本的路由都将存放在这里

	api := r.Group("admin/v1")
	guard := "admin"
	// 全局限流中间件：每小时限流，这里是所有API（根据IP）请求加起来
	// 作为参考 GITHUB API 每小时最多 60个请求（根据IP）

	api.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := api.Group("/auth")
		// 限流中间键：每小时限流，作为参考github api 每小时最多 60个请求（根据IP）
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 注册用户
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(guard), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(guard), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(guard), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(guard), suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-code/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-code/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-code/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(guard), lgc.LoginByPhone)
			// 使用手机号，Email和用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(guard), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(guard), lgc.RefreshToken)

			//重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.AuthJWT(guard), pwc.ResetByphone)
			authGroup.POST("/password-reset/using-email", middlewares.AuthJWT(guard), pwc.ResetByEmail)

		}
		uc := new(controllers.UsersController)

		// 获取当前用户
		api.GET("/user", middlewares.AuthJWT(guard), uc.CurrentUser)

		usersGroup := api.Group("/users")
		{
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", middlewares.AuthJWT(guard), uc.UpdateProfile)
			usersGroup.PUT("/email", middlewares.AuthJWT(guard), uc.UpdateEmail)
			usersGroup.PUT("/phone", middlewares.AuthJWT(guard), uc.UpdatePhone)
			usersGroup.PUT("/password", middlewares.AuthJWT(guard), uc.UpdatePassword)
			usersGroup.PUT("/avatar", middlewares.AuthJWT(guard), uc.UpdateAvatar)
		}

		cgc := new(controllers.CategoriesController)
		cgcGroup := api.Group("/categories")
		{
			cgcGroup.GET("", cgc.Index)
			cgcGroup.GET("/:id", cgc.Show)
			cgcGroup.POST("", middlewares.AuthJWT(guard), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJWT(guard), cgc.Update)
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(guard), cgc.Delete)
		}

		tpc := new(controllers.TopicsController)
		tpcGroup := api.Group("/topics")
		{
			tpcGroup.GET("", tpc.Index)
			tpcGroup.GET("/:id", tpc.Show)
			tpcGroup.POST("", middlewares.AuthJWT(guard), tpc.Store)
			tpcGroup.PUT("/:id", middlewares.AuthJWT(guard), tpc.Update)
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(guard), tpc.Delete)
		}

		lsc := new(controllers.LinksController)
		linksGroup := api.Group("/links")
		{
			linksGroup.GET("", lsc.Index)
		}

	}

}
