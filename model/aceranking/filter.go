package aceranking

import (
	"acemap/model/dao"
	"time"
)

type Filter struct {
	AreaList []string `json:"area_list"`
	MinYear  int      `json:"min_year"`
	MaxYear  int      `json:"max_year"`
}

func GetFilter() Filter {
	f := Filter{
		AreaList: dao.GetAllArea(),
		MinYear:  2000,
		MaxYear:  time.Now().Year(),
	}
	return f
}
