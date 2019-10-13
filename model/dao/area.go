package dao

import (
	"acemap/data"
	"acemap/database"
	"sort"
)

func GetAllArea() []string {
	areaList := make([]string, 0)
	sql := "SELECT DISTINCT area FROM am_area;"
	result, err := database.MainDB.Query(sql)
	if err != nil {
		return areaList
	}
	for result.Next() {
		var area string
		_ = result.Scan(&area)
		areaList = append(areaList, area)
	}
	_ = result.Close()

	sort.Slice(areaList, func(i, j int) bool {
		if areaList[i] == "All" {
			return true
		} else if areaList[j] == "All" {
			return false
		} else if areaList[i] == "China" {
			return true
		} else if areaList[j] == "China" {
			return false
		} else {
			return areaList[i] < areaList[j]
		}
	})
	return areaList
}

func GetCountrySetByArea(area string) map[data.ID]bool {
	countrySet := make(map[data.ID]bool)
	sql := "SELECT country_id FROM am_area WHERE area = ?;"
	result, err := database.MainDB.Query(sql, area)
	if err != nil {
		return countrySet
	}
	for result.Next() {
		var countryID data.ID
		_ = result.Scan(&countryID)
		countrySet[countryID] = true
	}
	_ = result.Close()
	return countrySet
}

func GetAffiliationMaskByArea(area string) []bool {
	countrySet := GetCountrySetByArea(area)
	affiliationMask := make([]bool, data.NumOfAffiliations)
	for i, cid := range data.CountryList {
		if countrySet[cid] {
			affiliationMask[i] = true
		}
	}
	return affiliationMask
}

func GetAffiliationListByArea(area string) []int {
	countrySet := GetCountrySetByArea(area)
	affiliationList := make([]int, 0)
	for i, cid := range data.CountryList {
		if countrySet[cid] {
			affiliationList = append(affiliationList, i)
		}
	}
	return affiliationList
}
