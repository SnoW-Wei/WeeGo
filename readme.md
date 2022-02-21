## 功能模块
---
1. 注册
2. 登录
3. 找回密码
4. 注销账号
5. 修改密码

## 开发模块
---
- Restful Api
- 数据库ORM
- Redis
- 缓存
- 命令行
- 代码生成器（make 命令）
- 日志和错误处理
- 路由
- 数据库迁移
- 数据填充（faker）
- 安全验证码（短信、邮箱验证）
- 图片验证码
- 分页
- 授权策略
- 请求验证（JSON、表单、URI query 请求）
- 图片上传
- 图片裁剪
- 分页
- 限流（令牌桶）

## 框架基础内容
---
```
1. 配置信息（使用 Viper，支持 .env 和 config 目录 ）
2. API 版本
3. API 错误码
4. API 限流
    1)支持秒、分钟、小时、天级的请求限制
    2)支持返回 API 请求量标头（限制数，剩余量、重置时间）
5.注册登录
    1)注册
        判断手机是否注册
        判断 Email 是否注册
        支持手机 + 短信验证码进行注册
        支持使用邮箱注册账号
    2)登录
        支持手机 + 短信进行登录
        支持密码登录（手机号、Email、用户名任选）
        支持更加安全的 Token Refresh 机制
    3)找回密码
        支持使用手机 + 短信验证码找回
        支持使用邮箱 + 邮箱验证码找回
    4)注销账号
    5)修改密码    
6.JWT 授权
7.整个应用使用命令行模式（默认运行 web 服务）
8.内置命令行（ cobra，对比 cli 和 cobra）
    1)key 命令生成 app key
    2)make 命令
        make seeder —— 生成数据填充
        make policy —— 生成授权文件
        make apicontroller —— 生成 Restful API 控制器
        make model —— 生成模型文件
        make request —— 生成请求验证文件
        make factory —— 生成模型工厂文件
        make cmd —— 生成自定义命令文件
        make migration —— 生成数据库迁移文件
    3)seed 数据填充
        seed 所有数据
        seed 单条数据
        支持使用 faker 填充假数据
        支持模型工厂（ factory ）
    4)migrate 数据库迁移
        up —— 执行迁移
        rollback (down) —— 回滚上一步执行的迁移
        fresh —— 删除所有表，然后执行所有迁移
        reset —— 回滚所有迁移
        refresh —— 回滚所有迁移，然后再执行所有迁移
    5)cache 缓存处理
        cache clear —— 清除缓存
        cache forget —— 忘记某个 KEY 对应的缓存
9.分页
    支持返回上下页链接，方便客户端调用
10.Cache 缓存包
    支持 redis 缓存
    使用 interface ，支持使用多驱动
11.Redis 操作
12.安全验证码
    1)Email （发送邮箱，使用 Mailhog 进行测试）
    2)手机验证码（发送手机短信）
    3)内置 Redis 驱动，以接口方式编写，支持多驱动
13.图片验证码，防机器人滥用
    1)支持通过配置信息自定义复杂度
    2)内置 Redis 驱动，以接口方式编写，支持多驱动
14.日志记录
    1)集成 zap 高性能日志库
    2)支持命令行记录（方便开发时快速定位问题）
    3)命令行日志高亮
    4)支持文件记录（多文件和按日期分隔）
    5)记录 gorm 的 query log
    6)记录 HTTP 请求 log
    7)Panic Recovery 中间件
    8)合理的日志等级（debug, info, error, panic, fatal）
15.Policy 授权策略结构
16.Request 请求验证方案
    支持 JSON 请求、表单请求、URL Query
17.API 图片上传
18.图片裁切
19.数据库支持 mysql 和 sqlite
```

## 目录结构
---

```
.├── app                            // 程序具体逻辑代码
│   ├── cmd                         // 命令
│   │   ├── cache.go                
│   │   ├── cmd.go
│   │   ├── key.go
│   │   ├── make                    // make 命令及子命令
│   │   │   ├── make.go
│   │   │   ├── make_apicontroller.go
│   │   │   ├── make_cmd.go
│   │   │   ├── make_factory.go
│   │   │   ├── make_migration.go
│   │   │   ├── make_model.go
│   │   │   ├── make_policy.go
│   │   │   ├── make_request.go
│   │   │   ├── make_seeder.go
│   │   │   └── stubs               // make 命令的模板
│   │   │       ├── apicontroller.stub
│   │   │       ├── cmd.stub
│   │   │       ├── factory.stub
│   │   │       ├── migration.stub
│   │   │       ├── model
│   │   │       │   ├── model.stub
│   │   │       │   ├── model_hooks.stub
│   │   │       │   └── model_util.stub
│   │   │       ├── policy.stub
│   │   │       ├── request.stub
│   │   │       └── seeder.stub
│   │   ├── migrate.go
│   │   ├── play.go
│   │   ├── seed.go
│   │   └── serve.go
│   ├── http                        // http 请求处理逻辑
│   │   ├── controllers             // 控制器，存放 API 和视图控制器
│   │   │   ├── api                 // API 控制器，支持多版本的 API 控制器
│   │   │   │   └── v1              // v1 版本的 API 控制器
│   │   │   │       ├── users_controller.go
│   │   │   │       └── ...
│   │   └── middlewares             // 中间件
│   │       ├── auth_jwt.go
│   │       ├── guest_jwt.go
│   │       ├── limit.go
│   │       ├── logger.go
│   │       └── recovery.go
│   ├── models                      // 数据模型
│   │   ├── user                    // 单独的模型目录
│   │   │   ├── user_hooks.go       // 模型钩子文件
│   │   │   ├── user_model.go       // 模型主文件
│   │   │   └── user_util.go        // 模型辅助方法
│   │   └── ...
│   ├── policies                    // 授权策略目录
│   │   ├── category_policy.go
│   │   └── ...
│   └── requests                    // 请求验证目录（支持表单、标头、Raw JSON、URL Query）
│       ├── validators              // 自定的验证规则
│       │   ├── custom_rules.go
│       │   └── custom_validators.go
│       ├── user_request.go
│       └── ...
├── bootstrap                       // 程序模块初始化目录
│   ├── app.go  
│   ├── cache.go
│   ├── database.go
│   ├── logger.go
│   ├── redis.go
│   └── route.go
├── config                          // 配置信息目录
│   ├── app.go
│   ├── captcha.go
│   ├── config.go
│   ├── database.go
│   ├── jwt.go
│   ├── log.go
│   ├── mail.go
│   ├── pagination.go
│   ├── redis.go
│   ├── sms.go
│   └── verifycode.go
├── database                        // 数据库相关目录
│   ├── database.db                 // sqlite 数据文件（加入到 .gitignore 中）
│   ├── factories                   // 模型工厂目录
│   │   ├── user_factory.go
│   │   └── ...
│   ├── migrations                  // 数据库迁移目录
│   │   ├── 2021_12_21_102259_create_users_table.go
│   │   ├── 2021_12_21_102340_create_categories_table.go
│   │   └── ...
│   └── seeders                     // 数据库填充目录
│       ├── users_seeder.go
│       ├── ...
├── pkg                             // 内置辅助包
│   ├── app
│   ├── auth
│   ├── cache
│   ├── captcha
│   ├── config
│   └── ...
├── public                          // 静态文件存放目录
│   ├── css
│   ├── js
│   └── uploads                     // 用户上传文件目录
│       └── avatars                 // 用户上传头像目录
├── routes                          // 路由
│   ├── api.go
│   └── web.go
├── storage                         // 内部存储目录
│   ├── app
│   └── logs                        // 日志存储目录
│       ├── 2021-12-28.log
│       ├── 2021-12-29.log
│       ├── 2021-12-30.log
│       └── logs.log
└── tmp                             // air 的工作目录
├── .env                            // 环境变量文件
├── .env.example                    // 环境变量示例文件
├── .gitignore                      // git 配置文件
├── .air.toml                       // air 配置文件
├── .editorconfig                   // editorconfig 配置文件
├── go.mod                          // Go Module 依赖配置文件
├── go.sum                          // Go Module 模块版本锁定文件
├── main.go                         // Gohub 程序主入口
├── Makefile                        // 自动化命令文件
├── readme.md                       // 项目 readme
```


## 第三方依赖
---
使用到的开源库：

- [gin](https://github.com/gin-gonic/gin) —— 路由、路由组、中间件
- [zap](https://github.com/gin-contrib/zap) —— 高性能日志方案
- [gorm](https://github.com/go-gorm/gorm) —— ORM 数据操作
- [cobra](https://github.com/spf13/cobra) —— 命令行结构
- [viper](https://github.com/spf13/viper) —— 配置信息
- [cast](https://github.com/spf13/cast) —— 类型转换
- [redis](https://github.com/go-redis/redis/v8) —— Redis 操作
- [jwt](https://github.com/golang-jwt/jwt) —— JWT 操作
- [base64Captcha](https://github.com/mojocn/base64Captcha) —— 图片验证码
- [govalidator](https://github.com/thedevsaddam/govalidator) —— 请求验证器
- [limiter](https://github.com/ulule/limiter/v3) —— 限流器
- [email](https://github.com/jordan-wright/email) —— SMTP 邮件发送
- [aliyun-communicate](https://github.com/KenmyZhang/aliyun-communicate) —— 发送阿里云短信
- [ansi](https://github.com/mgutz/ansi) —— 终端高亮输出
- [strcase](https://github.com/iancoleman/strcase) —— 字符串大小写操作
- [pluralize](https://github.com/gertd/go-pluralize) —— 英文字符单数复数处理