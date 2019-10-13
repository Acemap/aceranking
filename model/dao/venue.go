package dao

import (
	"acemap/data"
	"acemap/database"
	"acemap/model"
	"fmt"
)

func GetFieldVenueMapByIndexAndLevel(index string, level string) map[string][]data.ID {
	fieldVenueMap := make(map[string][]data.ID)

	sql := "SELECT venue_id, field FROM am_venue_category WHERE paper_index = ? AND level = ?;"
	result, err := database.MainDB.Query(sql, index, level)
	if err != nil {
		return fieldVenueMap
	}
	for result.Next() {
		var fieldName string
		var venueID data.ID
		_ = result.Scan(&venueID, &fieldName)
		if fieldVenueMap[fieldName] == nil {
			fieldVenueMap[fieldName] = make([]data.ID, 0)
		}
		fieldVenueMap[fieldName] = append(fieldVenueMap[fieldName], venueID)
	}
	_ = result.Close()
	return fieldVenueMap
}

func GetConferenceAbbrByID(conferenceID data.ID) string {
	//cacheKey := fmt.Sprintf("ConferenceAbbrByID/%d", conferenceID)
	//if f, ok := model.Cache.Get(cacheKey); ok {
	//	abbr := f.(string)
	//	return abbr
	//}

	sql := "SELECT abbreviation FROM am_conference_series WHERE conference_series_id = ?;"
	result, err := database.MainDB.Query(sql, conferenceID)
	if err != nil {
		return ""
	}
	var abbr string
	for result.Next() {
		_ = result.Scan(&abbr)
	}
	_ = result.Close()

	//model.Cache.SetDefault(cacheKey, abbr)
	return abbr
}

func GetJournalNameByID(journalID data.ID) string {
	//cacheKey := fmt.Sprintf("JournalNameByID/%d", journalID)
	//if f, ok := model.Cache.Get(cacheKey); ok {
	//	name := f.(string)
	//	return name
	//}

	sql := "SELECT name FROM am_journal WHERE journal_id = ?;"
	result, err := database.MainDB.Query(sql, journalID)
	if err != nil {
		return ""
	}
	var name string
	for result.Next() {
		_ = result.Scan(&name)
	}
	_ = result.Close()

	//model.Cache.SetDefault(cacheKey, name)
	return name
}

func GetVenueNameByID(venueID data.ID) string {
	t := model.GetIDType(venueID)
	var name string
	if t == "journal" {
		name = GetJournalNameByID(venueID)
	} else if t == "conference" {
		name = GetConferenceAbbrByID(venueID)
	}
	return name
}

func GetFieldByVenueID(typ string, venueID data.ID) string {
	cacheKey := fmt.Sprintf("FieldByVenueID/%s/%d", typ, venueID)
	if f, ok := model.Cache.Get(cacheKey); ok {
		field := f.(string)
		return field
	}
	sql := "SELECT field FROM am_venue_category WHERE venue_id = ? AND paper_index = ?;"
	result, err := database.MainDB.Query(sql, venueID, typ)
	if err != nil {
		return ""
	}
	var field string
	for result.Next() {
		_ = result.Scan(&field)
	}
	_ = result.Close()
	model.Cache.SetDefault(cacheKey, field)
	return field
}
