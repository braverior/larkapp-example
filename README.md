# larkapp-example

飞书APP示例，以SDK的形式接入

## 说明
本项目优点：
> 统一封装了oapi-sdk-go初始化、事件回调、消息发送、卡片发送等能力，只需要在event_handler目录里实现按照的IEventHandler接口即可



按照飞书目前(2024年5月1日前)开放的能力，
* oapi-sdk-go版本v3.2.4
* 接收消息和按钮采用了长连接的方式
* 卡片回调消息回调仅支持HTTP形式

如果无需卡片回调，可以不初始化`LarkAppConfig.CardEventPort`和`LarkAppConfig.CardEventPath`配置

注意：
该项目仅支持部分飞书能力，如需要更强能力可以催更，也可以clone自行增加并实现IEventHandler接口

## 运行

`go run cmd/main.go`
