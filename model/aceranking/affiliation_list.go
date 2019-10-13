package aceranking

import (
	"acemap/data"
	"acemap/model"
	"acemap/model/dao"
	"fmt"
)

type AffiliationStatistics struct {
	AffiliationID   data.ID `json:"affiliation_id"`
	AffiliationName string  `json:"name"`
	Abbreviation    string  `json:"abbr"`
	Statistics
}

func GetAffStatOnAllAuthors(venueMask []bool, startYear, endYear, affIndex int, refCount []int) AffiliationStatistics {
	cacheKey := fmt.Sprintf("GetAffStatOnAllAuthors/%v/%d/%d/%d", venueMask, startYear, endYear, affIndex)
	if value, ok := model.Cache.Get(cacheKey); ok {
		statistics := value.(AffiliationStatistics)
		return statistics
	}

	affID := data.AffiliationList[affIndex]
	statistics := AffiliationStatistics{
		AffiliationID:   affID,
		AffiliationName: dao.GetAffiliationNameByID(affID),
		Abbreviation:    dao.GetAffiliationAbbrByID(affID),
	}
	citationList := make([]int, 0)
	//refCount := GetRefCount(venueMask, startYear, endYear)

	for _, paperIndex := range data.AffiliationPaperList[affIndex] {
		if !venueMask[data.PaperInfoList[paperIndex].VenueIndex] {
			statistics.AceScore += refCount[paperIndex]
		} else {
			paperInfo := data.PaperInfoList[paperIndex]
			if paperInfo.Year > endYear || paperInfo.Year < startYear {
				continue
			}
			for _, authorAffPair := range paperInfo.AuthorAffList {
				aff := authorAffPair.AffiliationIndex
				if aff == affIndex {
					statistics.CitationShare += authorAffPair.CitationShare
				}
			}
			statistics.Citation += paperInfo.Citation
			citationList = append(citationList, paperInfo.Citation)
		}
	}
	statistics.AceScore += statistics.Citation
	statistics.Count = len(citationList)
	statistics.HIndex = model.CalcHIndexByCitationList(citationList)

	model.Cache.SetDefault(cacheKey, statistics)
	return statistics
}

func GetAffStatOnFirstAuthor(venueMask []bool, startYear, endYear, affIndex int, refCount []int) AffiliationStatistics {
	cacheKey := fmt.Sprintf("GetAffStatOnFirstAuthor/%v/%d/%d/%d", venueMask, startYear, endYear, affIndex)
	if value, ok := model.Cache.Get(cacheKey); ok {
		statistics := value.(AffiliationStatistics)
		return statistics
	}

	affID := data.AffiliationList[affIndex]
	statistics := AffiliationStatistics{
		AffiliationID:   affID,
		AffiliationName: dao.GetAffiliationNameByID(affID),
		Abbreviation:    dao.GetAffiliationAbbrByID(affID),
	}
	citationList := make([]int, 0)
	//refCount := GetRefCount(venueMask, startYear, endYear)

	for _, paperIndex := range data.AffiliationPaperList[affIndex] {
		ok := isAffiliationFirst(affIndex, paperIndex)
		if !venueMask[data.PaperInfoList[paperIndex].VenueIndex] {
			if ok {
				statistics.AceScore += refCount[paperIndex]
			}
		} else {
			paperInfo := data.PaperInfoList[paperIndex]
			if paperInfo.Year > endYear || paperInfo.Year < startYear {
				continue
			}
			for _, authorAffPair := range paperInfo.AuthorAffList {
				aff := authorAffPair.AffiliationIndex
				if aff == affIndex {
					statistics.CitationShare += authorAffPair.CitationShare
				}
			}
			if ok {
				statistics.Citation += paperInfo.Citation
				citationList = append(citationList, paperInfo.Citation)
			}
		}
	}
	statistics.AceScore += statistics.Citation
	statistics.Count = len(citationList)
	statistics.HIndex = model.CalcHIndexByCitationList(citationList)

	model.Cache.SetDefault(cacheKey, statistics)
	return statistics
}

func isAffiliationFirst(affIndex, paperIndex int) bool {
	paperInfo := data.PaperInfoList[paperIndex]
	firstAuthor := paperInfo.AuthorAffList[0].AuthorIndex
	for _, authorAffPair := range paperInfo.AuthorAffList {
		aff := authorAffPair.AffiliationIndex
		author := authorAffPair.AuthorIndex
		if author != firstAuthor {
			break
		}
		if aff == affIndex {
			return true
		}
	}
	return false
}
