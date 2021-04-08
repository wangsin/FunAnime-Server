package feishu

import "github.com/gin-gonic/gin"

type Request interface {
	GetRequest(ctx *gin.Context) error
}

type EventCallbackRequest struct {
	Challenge string       `json:"challenge,omitempty"`
	UUID      string       `json:"uuid"`
	Token     string       `json:"token"`
	Ts        string       `json:"ts"`
	Type      string       `json:"type"`
	Event     EventElement `json:"event"`
}

func (e *EventCallbackRequest) GetRequest(ctx *gin.Context) error {
	if err := ctx.Bind(e); err != nil {
		return err
	}

	return nil
}

type EventElement struct {
	Type             string `json:"type"`
	AppID            string `json:"app_id"`
	TenantKey        string `json:"tenant_key"`
	RootID           string `json:"root_id"`
	ParentID         string `json:"parent_id"`
	OpenChatID       string `json:"open_chat_id"`
	ChatType         string `json:"chat_type"`
	MsgType          string `json:"msg_type"`
	OpenID           string `json:"open_id"`
	OpenMessageID    string `json:"open_message_id"`
	IsMention        bool   `json:"is_mention"`
	Text             string `json:"text"`
	TextWithoutAtBot string `json:"text_without_at_bot"`
}
