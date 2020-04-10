package service

import (
	"aceranking/cache"
	"aceranking/dao"
	"aceranking/model"
	"errors"
	"fmt"
	"sort"
)

func Venue(req *model.VenueReq) (*model.VenueResp, error) {
	var fieldVenueMap map[string]model.IDList
	typ := req.Type
	if typ == "CCF" {
		fieldVenueMap = dao.GetFieldVenueMapByIndexAndLevel("CCF", "A")
	} else if typ == "SCI" {
		fieldVenueMap = dao.GetFieldVenueMapByIndexAndLevel("SCI", "1")
	} else if typ == "IEEE Society" {
		fieldVenueMap = dao.GetFieldVenueMapByIndexAndLevel("IEEE Society", "1")
	} else if typ == "THU" {
		fieldVenueMap = dao.GetFieldVenueMapByIndexAndLevel("THU", "A")
	} else if typ == "ACM Society" {
		fieldVenueMap = dao.GetFieldVenueMapByIndexAndLevel("ACM Society", "1")
	} else {
		return nil, errors.New("unknown type: " + typ)
	}

	var vids model.IDList
	for _, venueList := range fieldVenueMap {
		vids = append(vids, venueList...)
	}
	cntMap := dao.GetVenuePaperCount(vids, &req.Filter)

	fieldNames := make([]string, 0, len(fieldVenueMap))
	for fieldName := range fieldVenueMap {
		fieldNames = append(fieldNames, fieldName)
	}
	sort.Slice(fieldNames, func(i, j int) bool {
		return fieldNames[i] < fieldNames[j]
	})

	allCnt := 0
	var fieldList []model.VenueContainer
	for id, fieldName := range fieldNames {
		venueIDList := fieldVenueMap[fieldName]
		venueInfoList := make([]model.VenueInfo, 0, len(venueIDList))
		cntSum := 0
		for _, v := range venueIDList {
			cnt := cntMap[v]
			cntSum += cnt
			venueInfoList = append(venueInfoList, model.VenueInfo{
				VenueID:  fmt.Sprintf("%d_%d", v, id),
				Type:     v.GetIDType(),
				Name:     cache.Cache(dao.GetVenueNameByID, v).(string),
				Count:    cnt,
				CountStr: ShadowCount(cnt),
			})
		}
		sort.Slice(venueInfoList, func(i, j int) bool {
			return venueInfoList[i].Count > venueInfoList[j].Count
		})
		fieldList = append(fieldList, model.VenueContainer{
			ID:       fieldName,
			Name:     fieldName,
			Count:    cntSum,
			CountStr: ShadowCount(cntSum),
			List:     venueInfoList,
		})
		allCnt += cntSum
	}
	sort.Slice(fieldList, func(i, j int) bool {
		return fieldList[i].Count > fieldList[j].Count
	})
	resp := &model.VenueResp{
		{
			ID:       "0",
			Name:     "All",
			Count:    allCnt,
			CountStr: ShadowCount(allCnt),
			List:     fieldList,
		},
	}

	return resp, nil
}
