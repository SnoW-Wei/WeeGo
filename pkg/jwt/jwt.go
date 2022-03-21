/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-05 10:09:36
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 21:29:10
 */
package jwt

import (
	"errors"
	"strings"
	"time"
	"weego/pkg/app"
	"weego/pkg/cache"
	"weego/pkg/config"
	"weego/pkg/helpers"
	"weego/pkg/logger"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalFormed         error = errors.New("令牌格式错误")
	ErrTokenInvalid           error = errors.New("令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalfromed        error = errors.New("请求头中 Authorization 格式错误！")
)

// JWT 定义一个jwt 对象
type JWT struct {
	// 看守器 admin api user等
	Guards string

	// 秘钥，用以加密 jwt ,读取配置信息 app.key
	SignKey []byte

	// 刷新 token 的最大过期时间
	MaxRefresh time.Duration
}

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了 Valid() 方法
	// JWT 规定了7个官方自断，提供使用：
	// - iss(issuer): 发布者
	// - sub(subject):主题
	// - iat(Issued At):生成签名的时间
	// - exp(expiration time):过期时间
	// - aud(audience): 观众，相当于接受者
	// - nbf(Not Before):生效时间
	// - jti(JWT ID): JWT ID
	jwtpkg.StandardClaims
}

func NewJWT(gurad string) *JWT {
	return &JWT{
		Guards:     gurad,
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt("jwt.max_refresh_time")) * time.Second,
	}
}

// ParserToken 解析 token，中间键中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {

	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 1. 调用 jwt 库解析用户传参的Token
	token, err := jwt.parseTokenString(tokenString)

	// 2. 解析出错
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalFormed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {

		// 看守器判断
		if claims.Subject == jwt.Guards {
			return claims, nil
		}
	}
	return nil, ErrTokenInvalid
}

// RefreshToken 更新 token，用以提供 refresh token 接口
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {

	// 1. 从header 里获取token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的Token
	token, err := jwt.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)

		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*JWTCustomClaims)

	// 5. 检查是否过了“最大允许刷新的时间”
	x := app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		// 修改过期时间
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	}
	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken  生成 Token，在登录成功时调用
func (jwt *JWT) IssueToken(userID string, userName string) string {

	// 1.构造用户 claims 信息（负荷）
	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		userID,
		userName,
		expireAtTime,
		jwtpkg.StandardClaims{
			NotBefore: app.TimenowInTimezone().Unix(), // 签名生效时间
			IssuedAt:  app.TimenowInTimezone().Unix(), // 首次签名时间（后续刷新token 不会更新）
			ExpiresAt: expireAtTime,                   // 过期时间
			Issuer:    config.GetString("app.name"),   // 签名颁发者
			Subject:   jwt.Guards,                     //看守器
		},
	}

	// 2. 根据claims 生成token 对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// 使用 hs256 算法进行 token 生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {
	timenow := app.TimenowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}
	expire := time.Duration(expireTime) * time.Minute
	return timenow.Add(expire).Unix()
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// getTokenFromHeader 从 jwtpkg.ParseWithClaims 里获取 token
// Authorization: Bearer <token>
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	// 解析 Authorization header
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrTokenMalFormed
	}

	return parts[1], nil
}

// redis 黑名单实现token 过期 ：用户主动注销、强制登出(禁止登陆)、忘记密码、修改密码、JWT续签、踢出下线 等  userID:guard token

func (jwt *JWT) BlackListCache(c *gin.Context, args string) (bool, error) {
	// 1. 从header 里获取token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return false, ErrTokenInvalid
	}

	// 2. 调用 jwt 库解析用户传参的Token
	token, err := jwt.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)

		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return false, ErrTokenInvalid
		}
	}

	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*JWTCustomClaims)

	cacheKey := "token:" + claims.UserID + ":" + claims.Subject

	// 设置过期时间
	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}
	expire := time.Duration(expireTime) * time.Minute

	// check 数据
	if !helpers.Empty(cache.Get(cacheKey)) && args == "c" {
		if claims.StandardClaims.ExpiresAt < app.TimenowInTimezone().Add(cache.TTL(cacheKey)).Unix() {
			return false, ErrTokenInvalid
		}
	}
	if args == "s" {
		// set 缓存
		cache.Set(cacheKey, claims.Subject, expire)
	}

	return true, nil
}
