package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	"sinblog.cn/FunAnime-Server/serializable/request/video"
	responseVideo "sinblog.cn/FunAnime-Server/serializable/response/video"
	serviceVideo "sinblog.cn/FunAnime-Server/service/video"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/consts"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/logger"
	utilMath "sinblog.cn/FunAnime-Server/util/math"
	"strconv"
)

func GetVideoListForOuter(ctx *gin.Context) {
	req := video.GetVideoListForOuter{}
	err := req.GetRequest(ctx)
	if err != nil {
		logger.Error("get_request_data_failed", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	videoList, count, errNo := serviceVideo.GetVideoList(&req)
	if errNo != errno.Success {
		logger.Error("get_video_list_failed", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errNo)
		return
	}

	common.EchoBaseJson(ctx, http.StatusOK, errno.Success, responseVideo.BuildVideoListResponse(videoList, req.Page, req.Size, count))
}

func GetVideoDetailForOuter(ctx *gin.Context) {
	vIdStr := ctx.Param("id")
	videoId, err := strconv.ParseInt(vIdStr, 10, 64)
	if err != nil {
		logger.Error("params_error", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	loginUserInfo := user.GetUserInfoFromContext(ctx)

	videoDetail, userInfo, collected, errNo := serviceVideo.GetVideoDetail(videoId, loginUserInfo)
	if errNo != errno.Success {
		logger.Error("get_video_detail_failed", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errNo)
		return
	}

	videoResp := responseVideo.VideoDetailResponse{
		VideoName:     videoDetail.VideoName,
		VideoRemoteId: videoDetail.VideoRemoteId,
		CreateTime:    videoDetail.PassTime.Format(consts.TimeFormatYMDHM),
		Category:      fmt.Sprintf("%s/%s", videoDetail.CategoryTopLevelDesc, videoDetail.CategoryNextLevelDesc),
		Pv:            utilMath.GetHumanFormatNumber(videoDetail.Pv),
		IsCollect:     collected,
		Creator:       userInfo.Nickname,
		CreatorImg:    common.BuildImageLink(userInfo.Avatar),
	}
	common.EchoBaseJson(ctx, http.StatusOK, errno.Success, videoResp)
	return
}
