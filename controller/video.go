package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	"sinblog.cn/FunAnime-Server/serializable/request/video"
	responseVideo "sinblog.cn/FunAnime-Server/serializable/response/video"
	serviceCommon "sinblog.cn/FunAnime-Server/service/common"
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

	barrage, errNo := serviceVideo.GetBarrageList(videoId)
	if errNo != errno.Success {
		logger.Error("upload_video_failed", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errNo)
		return
	}

	videoResp := responseVideo.VideoDetailResponse{
		VideoName:     videoDetail.VideoName,
		VideoDesc:     videoDetail.VideoDesc,
		VideoRemoteId: videoDetail.VideoRemoteId,
		CreateTime:    videoDetail.PassTime.Format(consts.TimeFormatYMDHM),
		Category:      fmt.Sprintf("%s/%s", videoDetail.CategoryTopLevelDesc, videoDetail.CategoryNextLevelDesc),
		Pv:            utilMath.GetHumanFormatNumber(videoDetail.Pv),
		IsCollect:     collected,
		Creator:       userInfo.Nickname,
		CreatorImg:    common.BuildImageLink(userInfo.Avatar),
		BarrageList:   responseVideo.BuildBarrageArrayResp(barrage),
	}
	common.EchoBaseJson(ctx, http.StatusOK, errno.Success, videoResp)
	return
}

func GetVideoUploadSign(ctx *gin.Context) {
	common.EchoSuccessJson(ctx, gin.H{"sign": serviceCommon.GetVideoUploadSign()})
}

func UploadVideo(ctx *gin.Context) {
	request := &video.UploadRequest{}
	err := request.GetRequest(ctx)
	if err != nil || request.UserInfo == nil {
		logger.Error("get_upload_video_params_error", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errno.UnknownError)
		return
	}

	errNo := serviceVideo.UploadVideo(request)
	if errNo != errno.Success {
		logger.Error("upload_video_failed", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errNo)
		return
	}

	common.EchoSuccessJson(ctx, gin.H{})
	return
}

func GetManageVideoList(ctx *gin.Context) {
	loginInfo := user.GetUserInfoFromContext(ctx)
	fmt.Println(loginInfo)
	list, count, errNo := serviceVideo.GetFixUserVideoList(loginInfo, 1, 20)
	if errNo != errno.Success {
		logger.Error("get_manage_video_list_failed", logger.Fields{"err": errNo})
		common.EchoFailedJson(ctx, errNo)
		return
	}

	common.EchoBaseJson(ctx, http.StatusOK, errNo, responseVideo.BuildVideoManageListResponse(list, 1, 100, count))
}

func GetBarrageList(ctx *gin.Context) {
	vIdStr := ctx.Param("id")
	videoId, err := strconv.ParseInt(vIdStr, 10, 64)
	if err != nil {
		logger.Error("params_error", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	barrage, errNo := serviceVideo.GetBarrageList(videoId)
	if errNo != errno.Success {
		logger.Error("upload_video_failed", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errNo)
		return
	}

	common.EchoBaseJson(ctx, http.StatusOK, errno.Success, responseVideo.BuildBarrageArrayResp(barrage))
	return
}
