![rainyun-go-sdk](https://www.rainyun.com/img/logo.d193755d.png)

## 简介 📜

![GoReport](https://goreportcard.com/badge/github.com/ssdomei232/rainyun-go-sdk)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ssdomei232/rainyun-go-sdk) 
![GitHub](https://img.shields.io/github/license/ssdomei232/rainyun-go-sdk) 
![GitHub tag (with filter)](https://img.shields.io/github/v/tag/ssdomei232/rainyun-go-sdk)

使用`rainyun-go-sdk`来顺畅的接入Rainyun API，物品兑换，个人信息查询...

> [!WARNING]  
> 本项目由第三方开发,不保证可用性和及时性,api接口的变更我们无法即时获知,PR and Issues are welcome.

## 使用 🥑

1. 下载包

```bash
go get github.com/ssdomei232/rainyun-go-sdk
```

2. 导入包

```go
import rain "github.com/ssdomei232/rainyun-go-sdk"
```

3. 使用包

参考 👉[rainyun-go-sdk-example](https://git.mmeiblog.cn/mei/rainyun-go-sdk-example)

4. Client 介绍

```go
// 使用 rainyun api key 创建client
client := rain.NewClient("your-real-api-key")
```

5. 错误码

`client.Dorequest`会处理雨云返回的错误码并返还,相关定义位于 [model.go](https://git.mmeiblog.cn/mei/rainyun-go-sdk/src/branch/main/model.go)

6. 验证码

目前sdk无法处理验证码,[一键气死搬砖狐](https://api.zzwl.top/rainyun-tasks/)是应该是狐狸给加的白名单

## TODO List

按先后顺序:

[ ] 积分兑换
[ ] 域名管理
[ ] rcs管理
[ ] rca管理
[ ] ros管理
[ ] rgs管理
[ ] 工单管理
[ ] 裸金属管理
[ ] 二级面板
