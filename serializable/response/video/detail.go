package responseVideo

type VideoDetailResponse struct {
	VideoName     string `json:"video_name"`
	VideoRemoteId string `json:"video_remote_id"`
	CreateTime    string `json:"create_time"`
	Category      string `json:"category"`
	Pv            string `json:"pv"`
	IsCollect     bool   `json:"is_collect"`
	Creator       string `json:"creator"`
	CreatorImg    string `json:"creator_img"`
}
