![rainyun-go-sdk](https://www.rainyun.com/img/logo.d193755d.png)

## 简介 📜

![GoReport](https://goreportcard.com/badge/github.com/ssdomei232/rainyun-go-sdk)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ssdomei232/rainyun-go-sdk)
![GitHub](https://img.shields.io/github/license/ssdomei232/rainyun-go-sdk)
![GitHub tag (with filter)](https://img.shields.io/github/v/tag/ssdomei232/rainyun-go-sdk)

使用`rainyun-go-sdk`来顺畅的接入Rainyun API，物品兑换，个人信息查询...

> [!WARNING]  
> 本项目由第三方开发,不保证可用性和及时性,api接口的变更我们无法即时获知,PR and Issues are welcome.

> [!TIP]
> 如果您想支持 sdk 的发展，欢迎从[这个链接](https://www.rainyun.com/cat_)注册雨云，还可以在积分商城领取新购7折和续费85折独家优惠券:)
> 直播地址: [https://live.mmeiblog.cn/web/cinema/3205c0020f1e40af80d02e45d772203f](https://live.mmeiblog.cn/web/cinema/3205c0020f1e40af80d02e45d772203f) (周五晚/周六/周天早)

## 使用 🥑

1. 导入包

```go
// 导入核心包
import rain "github.com/ssdomei232/rainyun-go-sdk/rainyun/common"
// 按需导入产品包
import "github.com/ssdomei232/rainyun-go-sdk/rainyun/rcs"
```

2. Client 介绍

```go
// 使用 rainyun api key 创建client
client := rain.NewClient("your-real-api-key")
rcsClient := rcs.Client{Client: client}
detail, err := rcsClient.GetRcsDetails(114514)
 if err != nil {
  log.Println("Error getting RCS details:", err)
 }
 fmt.Println("Backup Status:", detail.Data.RBSList)
 fmt.Println("=================================")
 fmt.Println("EIP Status:", detail.Data.EIPList)
```

3. 错误码

`client.Dorequest`会处理雨云返回的错误码并返还,相关定义位于 [model.go](https://github.com/ssdomei232/rainyun-go-sdk/blob/main/model.go)

4. 数据实例

我们在`data`下存放了部分响应/请求实例，或许会对你有所帮助

## 一些约定 🤝

1. 在游戏云MCSM中，淋雨管游戏叫蛋(Egg)
2. 1个月=31天
3. 结构体注解写的不全面，请优先参考官方写的很烂的文档和`data`下的实例，实在找不到请提issue
4. 用户可以登陆PVE是预期内行为

## TODO List

按先后顺序:

- [x] 积分兑换  
- [x] rcs管理  
- [x] 工单管理  
- [x] rgs管理  
- [x] ros管理  
- [x] 域名管理  
- [ ] 裸金属管理(working)

短期内不会支持的操作(如果你有这种操作的请求和返回示例欢迎发到issue中或提PR):

- [ ] rca管理(淋雨说有大变化,核心部分暂缓)  
