package larkapp

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"time"
)

type LarkClient struct {
	config *LarkAppConfig
	client *lark.Client
}

func NewLarkClient(config *LarkAppConfig) *LarkClient {
	larkClient := &LarkClient{}
	larkClient.config = config
	larkClient.client = lark.NewClient(
		larkClient.config.AppId, larkClient.config.AppSecret,
		lark.WithLogLevel(larkcore.LogLevel(larkClient.config.LogLevel)),
		lark.WithReqTimeout(time.Duration(larkClient.config.ReqTimeout)*time.Second),
		lark.WithEnableTokenCache(larkClient.config.EnableTokenCache),
	)
	return larkClient
}

// SendTextMessage 发送文本消息
func (c *LarkClient) SendTextMessage(receiveOpenId string, text, tenantKey string, atUserId string, inChat bool) error {
	textBuilder := larkim.NewMessageTextBuilder().Text(text)
	if atUserId == "all" {
		textBuilder = textBuilder.AtAll()
	} else if atUserId != "" {
		textBuilder = textBuilder.AtUser(atUserId, "")
	}
	msgType := larkim.MsgTypeText
	rsp, err := c.SendMessage(receiveOpenId, msgType, textBuilder.Build(), tenantKey, inChat)
	if err != nil {
		return err
	}
	if rsp.Code != 0 {
		return fmt.Errorf("SendTextMessage failed, code %d, msg %s", rsp.Code, rsp.Msg)
	}
	return nil
}

// SendCardTemplateMessage  发送卡片消息
func (c *LarkClient) SendCardTemplateMessage(receiveOpenId string, templateId, contentJson, tenantKey string, inChat bool) error {
	templateContent := `{"type":"template","data":{"template_id":"%s","template_variable":%s}}`
	body := fmt.Sprintf(templateContent, templateId, contentJson)
	msgType := larkim.MsgTypeInteractive
	rsp, err := c.SendMessage(receiveOpenId, msgType, body, tenantKey, inChat)
	if err != nil {
		return err
	}
	if rsp.Code != 0 {
		return fmt.Errorf("SendCardTemplateMessage failed, code %d, msg %s", rsp.Code, rsp.Msg)
	}
	return nil
}

// SendMessage 基础的消息发送
func (c *LarkClient) SendMessage(receiveOpenId string, msgType, textBuilder, tenantKey string, inChat bool) (*larkim.CreateMessageResp, error) {
	//内容builder
	bodyBuilder := larkim.NewCreateMessageReqBodyBuilder().MsgType(msgType).
		ReceiveId(receiveOpenId).Content(textBuilder).Build()

	openIdType := larkim.ReceiveIdTypeOpenId
	if inChat {
		openIdType = larkim.ReceiveIdTypeChatId
	}
	//请求builder
	reqBuilder := larkim.NewCreateMessageReqBuilder().ReceiveIdType(openIdType).Body(bodyBuilder).Build()
	//给指定租户发送消息
	return c.client.Im.Message.Create(context.Background(), reqBuilder, larkcore.WithTenantKey(tenantKey))
}
