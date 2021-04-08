package feishu

import requestFeishu "sinblog.cn/FunAnime-Server/serializable/request/feishu"

type URLVerification struct {

}

func (U *URLVerification) BasicHandler(request *requestFeishu.EventCallbackRequest) (map[string]interface{}, error) {
	resultMap := map[string]interface{}{
		"challenge": request.Challenge,
	}

	return resultMap, nil
}

