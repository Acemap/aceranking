package service

import (
	"aceranking/dao"
	"aceranking/model"
	"time"
)

func Filter() *model.FilterResp {
	areaList := dao.GetAreaList()
	resp := &model.FilterResp{
		AreaList:       areaList,
		MinYear:        2000,
		MaxYear:        time.Now().Year(),
		MinAcademicAge: 0,
		MaxAcademicAge: 80,
	}
	return resp
}
