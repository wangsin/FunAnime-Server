package config

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/model"
	"sinblog.cn/FunAnime-Server/serializable/request"
	serviceStruct "sinblog.cn/FunAnime-Server/service/struct"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/logger"
	"time"
)

func GetBasicConfig(ctx *gin.Context) ([]*serviceStruct.CarouselInfo, []*serviceStruct.BasicRouter, string, int64) {
	configData, err := model.GetBasicConfigList(map[string]interface{}{
		"config_type=?": 0,
	})
	if err != nil || len(configData) <= 0 {
		logger.Warn("get_config_list_failed_at_GetBasicConfig", logger.Fields{"err": err, "ctx": ctx.Request.Context()})
		return nil, nil, "", errno.DBOpError
	}

	cConfig := serviceStruct.MainConfig{}
	err = cConfig.FromJson(configData[0].ConfigData)
	if err != nil {
		logger.Warn("unmarshal_json_data_failed_at_GetBasicConfig", logger.Fields{"err":err, "ctx": ctx.Request.Context(), "data": configData})
		return nil, nil, "", errno.DBOpError
	}

	cConfig.BuildImgLink()
	return cConfig.CarouselImg, cConfig.Router, cConfig.SearchText, errno.Success
}

func SetBasicConfig(ctx *gin.Context, req *request.BasicConfig) int64 {
	db, err := model.GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_faile_at_SetBasicConfig", logger.Fields{"err":err, "req":req})
		return errno.DBOpError
	}

	tx := db.Begin()
	err = model.UpdateConfigWithTrans(tx, map[string]interface{}{
		"status": model.FaConfigStatusValid,
		"config_type": req.ConfigType,
	}, map[string]interface{}{
		"status": model.FaConfigStatusDeleted,
		"modify_time": time.Now(),
	}, 1)
	if err != nil {
		tx.Rollback()
		logger.Error("update_config_with_trans_failed", logger.Fields{"err":err})
		return errno.DBOpError
	}

	configData := serviceStruct.MainConfig{
		CarouselImg: req.CarouselImg,
		SearchText:  req.SearchArea,
		Router:      req.HeadRouter,
	}
	err = model.CreateConfigWithTrans(tx, &model.FaConfig{
		ConfigType: req.ConfigType,
		Status:     model.FaConfigStatusValid,
		ConfigData: configData.ToJson(),
		ConfigUser: 0,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	})
	if err != nil {
		tx.Rollback()
		logger.Error("create_config_with_trans_failed", logger.Fields{"err":err})
		return errno.DBOpError
	}

	tx.Commit()
	return errno.Success
}