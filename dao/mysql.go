package dao

import (
	"aceranking/cache"
	"aceranking/model"
	"github.com/jmoiron/sqlx"
	"sort"
)

//===========================================================

func GetAreaList() []string {
	var areaList []string
	sql := "SELECT DISTINCT area FROM am_area;"
	if err := MysqlClient.Select(&areaList, sql); err != nil {
		return nil
	}

	// Set "All" to the first, and sort others.
	sort.Slice(areaList, func(i, j int) bool {
		if areaList[i] == "All" {
			return true
		} else if areaList[j] == "All" {
			return false
		} else {
			return areaList[i] < areaList[j]
		}
	})
	return areaList
}

//===========================================================

type venueField struct {
	VenueID model.ID `db:"venue_id"`
	Field   string   `db:"field"`
	Index   string   `db:"paper_index"`
	Level   string   `db:"level"`
}

func GetVenueField() []venueField {
	var venueFieldList []venueField
	sql := "SELECT venue_id, field, paper_index, level FROM am_venue_category;"
	_ = MysqlClient.Select(&venueFieldList, sql)
	return venueFieldList
}

func GetFieldVenueMapByIndexAndLevel(index string, level string) map[string]model.IDList {
	venueFieldList := cache.Cache(GetVenueField).([]venueField)
	fieldVenueMap := make(map[string]model.IDList)
	for _, row := range venueFieldList {
		if row.Index == index && row.Level == level {
			fieldVenueMap[row.Field] = append(fieldVenueMap[row.Field], row.VenueID)
		}
	}
	return fieldVenueMap
}

func GetFieldByVenueID(typ string, venueID model.ID) string {
	venueFieldList := cache.Cache(GetVenueField).([]venueField)
	for _, row := range venueFieldList {
		if row.VenueID == venueID && row.Index == typ {
			return row.Field
		}
	}
	return "None"
}

//===========================================================

func GetVenueIDToNameMap() map[model.ID]string {
	venueIDToName := make(map[model.ID]string)

	sql := "SELECT conference_series_id, abbreviation FROM am_conference_series;"
	cursor, err := MysqlClient.Query(sql)
	if err != nil {
		return venueIDToName
	}
	for cursor.Next() {
		var id model.ID
		var name string
		_ = cursor.Scan(&id, &name)
		venueIDToName[id] = name
	}

	sql = "SELECT journal_id, name FROM am_journal;"
	cursor, err = MysqlClient.Query(sql)
	if err != nil {
		return venueIDToName
	}
	for cursor.Next() {
		var id model.ID
		var name string
		_ = cursor.Scan(&id, &name)
		venueIDToName[id] = name
	}
	return venueIDToName
}

func GetVenueNameByID(venueID model.ID) string {
	venueIDToName := cache.Cache(GetVenueIDToNameMap).(map[model.ID]string)
	return venueIDToName[venueID]
}

//===========================================================

func GetCountryListByArea(area string) model.IDList {
	var countryList model.IDList
	sql := "SELECT country_id FROM am_area WHERE area = ?;"
	if err := MysqlClient.Select(&countryList, sql, area); err != nil {
		return nil
	}
	return countryList
}

func GetAffiliationListByArea(area string) model.IDList {
	var affiliationList model.IDList
	countryList := GetCountryListByArea(area)
	query, args, err := sqlx.In("SELECT affiliation_id FROM am_affiliation WHERE country_id IN (?);", countryList)
	if err != nil {
		return nil
	}
	if err := MysqlClient.Select(&affiliationList, query, args...); err != nil {
		return nil
	}
	return affiliationList
}

//===========================================================
