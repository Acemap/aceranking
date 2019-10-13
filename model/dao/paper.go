package dao

import (
	"acemap/data"
	"acemap/database"
)

func GetPaperTitleByID(paperID data.ID) string {
	sql := "SELECT title FROM am_paper WHERE paper_id = ?;"
	result, err := database.MainDB.Query(sql, paperID)
	if err != nil {
		return ""
	}
	var title string
	for result.Next() {
		_ = result.Scan(&title)
	}
	_ = result.Close()
	return title
}
