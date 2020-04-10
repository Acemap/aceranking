package service

import (
	"aceranking/model"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ShadowCount(count int) string {
	if count >= 1000 {
		kCnt := count / 1000
		return fmt.Sprintf("%dk+", kCnt)
	} else if count >= 100 {
		hCnt := count / 100
		return fmt.Sprintf("%d00+", hCnt)
	} else {
		return "<100"
	}
}

func SplitVenueIDs(VenueIDs string) model.IDList {
	var venueList model.IDList
	vs := strings.Split(VenueIDs, ",")
	for _, id := range vs {
		pair := strings.Split(id, "_")
		if v, err := strconv.ParseUint(pair[0], 10, 32); err != nil {
			continue
		} else {
			venueList = append(venueList, model.ID(v))
		}
	}
	return venueList
}

func CutFieldAnalysis(f []model.FieldAnalysisType) []model.FieldAnalysisType {
	if len(f) > 10 {
		cnt := 0
		for i := 10; i < len(f); i++ {
			cnt += f[i].Count
		}
		f = append(f[:10], model.FieldAnalysisType{
			FieldName: "Others",
			Count:     cnt,
		})
	}
	return f
}

func CheckAffSet(paperInfo *model.Paper, index int, affIDSet map[model.ID]bool) bool {
	// Check if index-th author have an affiliation in affIDSet.
	for _, idx := range paperInfo.Author.AffIndex[index] {
		if affIDSet[paperInfo.Affiliation.IDs[idx]] {
			return true
		}
	}
	return false
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
