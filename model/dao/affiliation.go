package dao

import (
	"acemap/data"
	"acemap/database"
	"acemap/model"
	"fmt"
)

func GetAffiliationNameByID(affiliationID data.ID) string {
	cacheKey := fmt.Sprintf("AffiliationNameByID/%d", affiliationID)
	if f, ok := model.Cache.Get(cacheKey); ok {
		name := f.(string)
		return name
	}

	sql := "SELECT name FROM am_affiliation WHERE affiliation_id = ?;"
	result, err := database.MainDB.Query(sql, affiliationID)
	if err != nil {
		return ""
	}
	var name string
	for result.Next() {
		_ = result.Scan(&name)
	}
	_ = result.Close()

	model.Cache.SetDefault(cacheKey, name)
	return name
}

func GetAffiliationAbbrByID(affiliationID data.ID) string {
	cacheKey := fmt.Sprintf("AffiliationAbbrByID/%d", affiliationID)
	if f, ok := model.Cache.Get(cacheKey); ok {
		abbr := f.(string)
		return abbr
	}

	sql := "SELECT abbreviation FROM am_affiliation WHERE affiliation_id = ?;"
	result, err := database.MainDB.Query(sql, affiliationID)
	if err != nil {
		return ""
	}
	var abbr string
	for result.Next() {
		_ = result.Scan(&abbr)
	}
	_ = result.Close()

	if abbr == "" {
		abbr = GetAffiliationNameByID(affiliationID)
	}

	model.Cache.SetDefault(cacheKey, abbr)
	return abbr
}
