package model

import (
	"acemap/data"
	"sort"
	"strconv"
	"strings"
)

func GetIDType(id data.ID) string {
	if id == 0 {
		return "None"
	}
	if id <= 1000000000 {
		return "paper"
	}
	if id <= 2000000000 {
		return "author"
	}
	if id <= 2100000000 {
		return "field"
	}
	if id <= 2110000000 {
		return "affiliation"
	}
	if id <= 2120000000 {
		return "journal"
	}
	if id <= 2130000000 {
		return "conference"
	}
	if id <= 2140000000 {
		return "conference_instance"
	}
	if id <= 2141000000 {
		return "country"
	}
	return "undefined"
}

func ParseStringIDToSortedList(s string) ([]data.ID, error) {
	vs := strings.Split(s, ",")
	idList := make([]data.ID, 0, len(vs))
	for _, id := range vs {
		if id == "" {
			continue
		}
		v, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return []data.ID{}, err
		}
		idList = append(idList, data.ID(v))
	}
	sort.Slice(idList, func(i, j int) bool {
		return idList[i] < idList[j]
	})
	return idList, nil
}

func ParseStringIDToSet(s string) (map[data.ID]bool, error) {
	vs := strings.Split(s, ",")
	idSet := make(map[data.ID]bool)
	for _, id := range vs {
		if id == "" {
			continue
		}
		v, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return idSet, err
		}
		idSet[data.ID(v)] = true
	}
	return idSet, nil
}

func ShadowCount(count int) string {
	if count >= 1000 {
		kCnt := count / 1000
		return strconv.FormatInt(int64(kCnt), 10) + "k+"
	} else if count >= 100 {
		hCnt := count / 100
		return strconv.FormatInt(int64(hCnt), 10) + "00+"
	} else {
		return "<100"
	}
}

func CalcHIndexBySortedCitationList(sortedCitationList []int) int {
	for i, cit := range sortedCitationList {
		if i+1 > cit {
			return i
		}
	}
	return len(sortedCitationList)
}

func CalcHIndexByCitationList(citationList []int) int {
	sort.Slice(citationList, func(i, j int) bool {
		return citationList[i] > citationList[j]
	})
	return CalcHIndexBySortedCitationList(citationList)
}
