package aceranking

import (
	"acemap/data"
	"acemap/model"
	"acemap/model/dao"
	"errors"
	"fmt"
	"sort"
	"strconv"
)

type Container struct {
	ID       string         `json:"venue_id"`
	Name     string      `json:"name"`
	Count    int         `json:"-"`
	CountStr string      `json:"count"`
	List     interface{} `json:"list"`
}

type VenueInfo struct {
	VenueID  string  `json:"venue_id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Count    int     `json:"-"`
	CountStr string  `json:"count"`
}

func GetVenueCategory(typ, area string, startYear, endYear int) ([]Container, error) {
	cacheKey := fmt.Sprintf("GetVenueCategory/%s/%s/%d/%d", typ, area, startYear, endYear)
	if value, ok := model.Cache.Get(cacheKey); ok {
		all := value.([]Container)
		return all, nil
	}

	var fieldVenueMap map[string][]data.ID
	if typ == "CCF" {
		fieldVenueMap = dao.GetFieldVenueMapByIndexAndLevel("CCF", "A")
	} else if typ == "SCI" {
		fieldVenueMap = dao.GetFieldVenueMapByIndexAndLevel("SCI", "1")
	} else if typ == "IEEE Society" {
		fieldVenueMap = dao.GetFieldVenueMapByIndexAndLevel("IEEE Society", "1")
	} else {
		return nil, errors.New("wrong type, must be \"CCF\" or \"SCI\" or \"IEEE\"")
	}

	affiliationMask := dao.GetAffiliationMaskByArea(area)

	fieldList := make([]Container, 0)
	allCnt := 0
	id := 1
	for fieldName, venueIDList := range fieldVenueMap {
		venueInfoList := make([]VenueInfo, len(venueIDList))
		cntSum := 0
		for i, v := range venueIDList {
			venueIndex := data.VenueIDToIndex[v]
			cnt := getCountByAffAndYear(data.VenuePaperMap[venueIndex], affiliationMask, startYear, endYear)
			cntSum += cnt
			venueInfoList[i].Type = model.GetIDType(v)
			venueInfoList[i].VenueID = fmt.Sprintf("%d_%d", v, id)
			venueInfoList[i].Count = cnt
			venueInfoList[i].CountStr = model.ShadowCount(cnt)
			venueInfoList[i].Name = dao.GetVenueNameByID(v)
		}
		sort.Slice(venueInfoList, func(i, j int) bool {
			return venueInfoList[i].Count > venueInfoList[j].Count
		})
		fieldList = append(fieldList, Container{
			ID:       strconv.FormatInt(int64(id), 10),
			Name:     fieldName,
			Count:    cntSum,
			CountStr: model.ShadowCount(cntSum),
			List:     venueInfoList,
		})
		id++
		allCnt += cntSum
	}
	sort.Slice(fieldList, func(i, j int) bool {
		return fieldList[i].Count > fieldList[j].Count
	})
	all := []Container{
		{
			ID:       "0",
			Name:     "All",
			Count:    allCnt,
			CountStr: model.ShadowCount(allCnt),
			List:     fieldList,
		},
	}
	model.Cache.SetDefault(cacheKey, all)
	return all, nil
}

func getCountByAffAndYear(paperIndexList []int, affiliationMask []bool, startYear, endYear int) int {
	cnt := 0
	for _, paperIndex := range paperIndexList {
		info := &data.PaperInfoList[paperIndex]
		if info.Year < startYear || info.Year > endYear {
			continue
		}
		for _, pair := range info.AuthorAffList {
			affIndex := pair.AffiliationIndex
			if affiliationMask[affIndex] {
				cnt++
				break
			}
		}
	}
	return cnt
}
