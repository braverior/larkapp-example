package event_handler

import (
	"context"
	"encoding/json"
	"fmt"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkapplication "github.com/larksuite/oapi-sdk-go/v3/service/application/v6"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"larkapp-example/conf"
)

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

type TestCard struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (c *TestCard) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c *TestCard) GetTemplateId() string {
	return conf.GlobalConfig.TestCardTemplateId
}

func (e *EventHandler) OnP2PersonCreatedV1(ctx context.Context, event *larkim.P1P2PChatCreatedV1) error {
	// 处理消息 event，这里简单打印消息的内容
	fmt.Println(larkcore.Prettify(event))
	fmt.Println(event.RequestId())
	return nil
}

func (e *EventHandler) OnP2MessageReceiveV1(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	// 处理消息 event，这里简单打印消息的内容
	fmt.Println(larkcore.Prettify(event))
	fmt.Println(event.RequestId())
	return nil
}

func (e *EventHandler) OnP2BotMenuV6(ctx context.Context, event *larkapplication.P2BotMenuV6) error {
	// 处理消息 event，这里简单打印消息的内容
	fmt.Println(larkcore.Prettify(event))
	fmt.Println(event.RequestId())
	return nil
}

func (e *EventHandler) OnCardAction(ctx context.Context, event *larkcard.CardAction) (interface{}, error) {
	// 处理消息 event，这里简单打印消息的内容
	fmt.Println(larkcore.Prettify(event))
	fmt.Println(event.RequestId())
	return nil, nil
}
