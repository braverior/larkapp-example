package larkapp

import (
	"context"
	"fmt"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	httpserverext "github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	dispatcher "github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
	"net/http"
	"strconv"
)

type LarkApp struct {
	config          *LarkAppConfig
	eventDispatcher *dispatcher.EventDispatcher
	eventHandler    IEventHandler
}

func NewLarkApp(handler IEventHandler, larkAPpConfig *LarkAppConfig) *LarkApp {
	larkApp := &LarkApp{}
	config := larkAPpConfig
	larkApp.config = config
	larkApp.eventHandler = handler
	return larkApp
}

func (app *LarkApp) Start() error {
	//处理Event事件
	if err := app.EventDispatcher(); err != nil {
		return fmt.Errorf("event dispatcher error: %v", err)
	}
	//处理Card事件
	if err := app.CardDispatcher(); err != nil {
		return fmt.Errorf("card dispatcher error: %v", err)
	}
	return nil
}

// EventDispatcher 处理消息回调
func (app *LarkApp) EventDispatcher() error {
	// 注册事件回调
	eventDispatcher := dispatcher.NewEventDispatcher(app.config.VerificationToken, app.config.EncryptKey)
	// 处理消息回传
	eventDispatcher = eventDispatcher.OnP2MessageReceiveV1(app.eventHandler.OnP2MessageReceiveV1)
	// 处理底部按钮
	eventDispatcher = eventDispatcher.OnP2BotMenuV6(app.eventHandler.OnP2BotMenuV6)
	// 保存事件回调
	app.eventDispatcher = eventDispatcher

	// 创建Client
	wsCli := larkws.NewClient(app.config.AppId, app.config.AppSecret,
		larkws.WithEventHandler(app.eventDispatcher),
		larkws.WithLogLevel(larkcore.LogLevel(app.config.LogLevel)),
	)
	// 启动客户端
	return wsCli.Start(context.Background())
}

// CardDispatcher 处理卡片消息回调
// 必须以http的形式运行，无法通过长连接形式, 需要配置 CardEventPort
func (app *LarkApp) CardDispatcher() error {
	if app.config.CardEventPort == 0 {
		return nil
	}
	cardHandler := larkcard.NewCardActionHandler(app.config.VerificationToken, app.config.EncryptKey, app.eventHandler.OnCardAction)
	// 注册处理器
	http.HandleFunc(app.config.CardEventPath, httpserverext.NewCardActionHandlerFunc(cardHandler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))
	// 启动 http 服务
	return http.ListenAndServe(strconv.Itoa(app.config.CardEventPort), nil)
}
