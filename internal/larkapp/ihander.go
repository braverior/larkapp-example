package larkapp

import (
	"context"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkapplication "github.com/larksuite/oapi-sdk-go/v3/service/application/v6"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// IEventHandler EventInterface 用户收到的消息处理
// TODO 目前仅添加常用的事件，如果需要可以继续增加
type IEventHandler interface {
	// OnP2MessageReceiveV1 收到用户消息
	OnP2MessageReceiveV1(ctx context.Context, event *larkim.P2MessageReceiveV1) error
	// OnP2BotMenuV6 用户点击窗口底部按钮
	OnP2BotMenuV6(ctx context.Context, event *larkapplication.P2BotMenuV6) error
	// OnP2PersonCreatedV1 用户聊天窗口建立
	OnP2PersonCreatedV1(ctx context.Context, event *larkim.P1P2PChatCreatedV1) error
	// OnCardAction 卡片回调
	OnCardAction(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error)
}
