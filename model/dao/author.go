package dao

import (
	"acemap/data"
	"acemap/database"
	"acemap/model"
	"fmt"
)

func GetAuthorNameByID(authorID data.ID) string {
	cacheKey := fmt.Sprintf("AuthorNameByID/%d", authorID)
	if f, ok := model.Cache.Get(cacheKey); ok {
		name := f.(string)
		return name
	}

	sql := "SELECT name FROM am_author WHERE author_id = ?;"
	result, err := database.MainDB.Query(sql, authorID)
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
