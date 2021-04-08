package feishu

import (
	"errors"
	requestFeishu "sinblog.cn/FunAnime-Server/serializable/request/feishu"
)

const (
	TypeURLVerification string = "url_verification"
	TypeEventCallback   string = "event_callback"
)

type EventHandler interface {
	BasicHandler(request *requestFeishu.EventCallbackRequest) (map[string]interface{}, error)
}

func DispatchCallbackByType(request *requestFeishu.EventCallbackRequest) (map[string]interface{}, error) {
	var b EventHandler
	switch request.Type {
	case TypeURLVerification:
		b = new(URLVerification)

	//case TypeEventCallback:
	//	// todo: implement

	default:
		// todo: log
		return nil, errors.New("type not defined")
	}

	return b.BasicHandler(request)
}
