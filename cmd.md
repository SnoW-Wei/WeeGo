<!--
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 00:13:17
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 00:36:34
-->
## 快捷开发命令
---
> 查看根命令 go run main.go -h

> 查看子命令 go run main.go [command] -h


### 创建数据库

1. 创建模型（创建model）
```
go run main.go make model project

```

2. 创建迁移（构建表结构）
```
go run main.go make migration add_projects_table
```

3. 执行迁移（新建数据）

```
go run main.go migrate up

```

### 填充数据
1. 创建工厂
```
go run main.go make factory project
```
2. 创建数据制造工具
```
go run main.go make seeder project
```
3. 填充数据
```
go run main.go seed SeedLinksTable
```
### 创建控制器
```
go run main.go make apicontroller v1/project
```

### 创建请求
```
go run main.go make request project
```

### 创建命令
```
go run main.go make cmd projectCommand
```

### 创建app_key命令
```
go run main.go key
```

### 缓存命令

1. 清理全部缓存
```
go run main.go cache clear
```

2. 清理key缓存
```
go run main.go cache forget --key=links:all
```

### 创建策略
```
go run main.go make policy project
```


