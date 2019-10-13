package aceranking

import (
	"acemap/data"
	"acemap/model"
	"acemap/model/dao"
	"fmt"
)

type AuthorAffInfo struct {
	AffiliationID   data.ID `json:"affiliation_id"`
	Abbreviation    string  `json:"abbr"`
	AffiliationName string  `json:"name"`
	Count           int     `json:"count"`
}

type AuthorStatistics struct {
	AuthorID      data.ID             `json:"author_id"`
	AuthorName    string              `json:"author_name"`
	CountAnalysis []CountAnalysisType `json:"count_analysis"`
	AffCount      map[int]int         `json:"-"`
	AffList       []AuthorAffInfo     `json:"affiliation_list"`
	Statistics
}

func GetAuthorStatOnAllAuthors(venueMask []bool, area string, startYear, endYear int) []AuthorStatistics {
	cacheKey := fmt.Sprintf("GetAuthorStatOnAllAuthors/%v/%s/%d/%d", venueMask, area, startYear, endYear)
	if value, ok := model.Cache.Get(cacheKey); ok {
		authorList := value.([]AuthorStatistics)
		return authorList
	}

	affiliationMask := dao.GetAffiliationMaskByArea(area)
	//statistics := make(map[int]*AuthorStatistics)
	//citationList := make(map[int][]int)
	statistics := make([]*AuthorStatistics, data.NumOfAuthors)
	citationList := make([][]int, data.NumOfAuthors)

	papers := GetPapers(venueMask, startYear, endYear)
	for _, paperIndex := range papers {
		paperInfo := data.PaperInfoList[paperIndex]
		mark := make(map[int]bool)
		for _, authorAffPair := range paperInfo.AuthorAffList {
			authorIndex := authorAffPair.AuthorIndex
			affIndex := authorAffPair.AffiliationIndex
			if !affiliationMask[affIndex] {
				continue
			}
			if !mark[authorIndex] {
				mark[authorIndex] = true
				if statistics[authorIndex] == nil {
					statistics[authorIndex] = &AuthorStatistics{
						AuthorID:      data.AuthorList[authorIndex],
						CountAnalysis: NewCountAnalysis(startYear, endYear),
						AffCount:      make(map[int]int),
					}
					citationList[authorIndex] = make([]int, 0)
				}
				statistics[authorIndex].Citation += paperInfo.Citation
				statistics[authorIndex].CitationShare += authorAffPair.CitationShare
				statistics[authorIndex].CountAnalysis[paperInfo.Year-startYear].Count++
				citationList[authorIndex] = append(citationList[authorIndex], paperInfo.Citation)
			}
			statistics[authorIndex].AffCount[affIndex]++
		}
	}

	refCount := GetRefCount(venueMask, startYear, endYear)
	for ref, count := range refCount {
		if count == 0 || venueMask[data.PaperInfoList[ref].VenueIndex] {
			continue
		}
		mark := make(map[int]bool)
		for _, authorAffPair := range data.PaperInfoList[ref].AuthorAffList {
			authorIndex := authorAffPair.AuthorIndex
			affIndex := authorAffPair.AffiliationIndex
			if !affiliationMask[affIndex] {
				continue
			}
			if !mark[authorIndex] {
				mark[authorIndex] = true
				if statistics[authorIndex] == nil {
					statistics[authorIndex] = &AuthorStatistics{
						AuthorID:      data.AuthorList[authorIndex],
						CountAnalysis: NewCountAnalysis(startYear, endYear),
					}
				}
				statistics[authorIndex].AceScore += count
			}
		}
	}

	for i, list := range citationList {
		if list != nil {
			statistics[i].Count = len(list)
			statistics[i].AceScore += statistics[i].Citation
			statistics[i].HIndex = model.CalcHIndexByCitationList(list)
		}
	}

	authorList := make([]AuthorStatistics, 0)
	for _, v := range statistics {
		if v != nil {
			authorList = append(authorList, *v)
		}
	}

	model.Cache.SetDefault(cacheKey, authorList)
	return authorList
}

func GetAuthorStatOnFirstAuthorWeak(venueMask []bool, area string, startYear, endYear int) []AuthorStatistics {
	cacheKey := fmt.Sprintf("GetAuthorStatOnAllAuthors/%v/%s/%d/%d", venueMask, area, startYear, endYear)
	if value, ok := model.Cache.Get(cacheKey); ok {
		authorList := value.([]AuthorStatistics)
		return authorList
	}

	affiliationMask := dao.GetAffiliationMaskByArea(area)
	//statistics := make(map[int]*AuthorStatistics)
	//citationList := make(map[int][]int)
	statistics := make([]*AuthorStatistics, data.NumOfAuthors)
	citationList := make([][]int, data.NumOfAuthors)

	papers := GetPapers(venueMask, startYear, endYear)
	for _, paperIndex := range papers {
		paperInfo := data.PaperInfoList[paperIndex]
		mark := make(map[int]bool)
		firstAff := paperInfo.AuthorAffList[0].AffiliationIndex
		for _, authorAffPair := range paperInfo.AuthorAffList {
			authorIndex := authorAffPair.AuthorIndex
			affIndex := authorAffPair.AffiliationIndex
			if !affiliationMask[affIndex] {
				continue
			}
			if !mark[authorIndex] {
				mark[authorIndex] = true
				if statistics[authorIndex] == nil {
					statistics[authorIndex] = &AuthorStatistics{
						AuthorID:      data.AuthorList[authorIndex],
						CountAnalysis: NewCountAnalysis(startYear, endYear),
						AffCount:      make(map[int]int),
					}
					citationList[authorIndex] = make([]int, 0)
				}
				statistics[authorIndex].CitationShare += authorAffPair.CitationShare
				if affIndex == firstAff {
					statistics[authorIndex].Citation += paperInfo.Citation
					statistics[authorIndex].CountAnalysis[paperInfo.Year-startYear].Count++
					citationList[authorIndex] = append(citationList[authorIndex], paperInfo.Citation)
				}
			}
			statistics[authorIndex].AffCount[affIndex]++
		}
	}

	refCount := GetRefCount(venueMask, startYear, endYear)
	for ref, count := range refCount {
		if count == 0 || venueMask[data.PaperInfoList[ref].VenueIndex] {
			continue
		}
		firstAff := data.PaperInfoList[ref].AuthorAffList[0].AffiliationIndex
		for _, authorAffPair := range data.PaperInfoList[ref].AuthorAffList {
			authorIndex := authorAffPair.AuthorIndex
			affIndex := authorAffPair.AffiliationIndex
			if !affiliationMask[affIndex] || affIndex != firstAff {
				continue
			}
			if statistics[authorIndex] == nil {
				statistics[authorIndex] = &AuthorStatistics{
					AuthorID:      data.AuthorList[authorIndex],
					CountAnalysis: NewCountAnalysis(startYear, endYear),
				}
			}
			statistics[authorIndex].AceScore += count
		}
	}

	for i, list := range citationList {
		if list != nil {
			statistics[i].Count = len(list)
			statistics[i].AceScore += statistics[i].Citation
			statistics[i].HIndex = model.CalcHIndexByCitationList(list)
		}
	}

	authorList := make([]AuthorStatistics, 0)
	for _, v := range statistics {
		if v != nil {
			authorList = append(authorList, *v)
		}
	}

	model.Cache.SetDefault(cacheKey, authorList)
	return authorList
}

func GetAuthorStatOnFirstAuthorStrong(venueMask []bool, area string, startYear, endYear int) []AuthorStatistics {
	cacheKey := fmt.Sprintf("GetAuthorStatOnAllAuthors/%v/%s/%d/%d", venueMask, area, startYear, endYear)
	if value, ok := model.Cache.Get(cacheKey); ok {
		authorList := value.([]AuthorStatistics)
		return authorList
	}

	affiliationMask := dao.GetAffiliationMaskByArea(area)
	//statistics := make(map[int]*AuthorStatistics)
	//citationList := make(map[int][]int)
	statistics := make([]*AuthorStatistics, data.NumOfAuthors)
	citationList := make([][]int, data.NumOfAuthors)

	papers := GetPapers(venueMask, startYear, endYear)
	for _, paperIndex := range papers {
		paperInfo := data.PaperInfoList[paperIndex]
		mark := make(map[int]bool)
		firstAuthor := paperInfo.AuthorAffList[0].AuthorIndex
		for _, authorAffPair := range paperInfo.AuthorAffList {
			authorIndex := authorAffPair.AuthorIndex
			affIndex := authorAffPair.AffiliationIndex
			if !affiliationMask[affIndex] {
				continue
			}
			if !mark[authorIndex] {
				mark[authorIndex] = true
				if statistics[authorIndex] == nil {
					statistics[authorIndex] = &AuthorStatistics{
						AuthorID:      data.AuthorList[authorIndex],
						CountAnalysis: NewCountAnalysis(startYear, endYear),
						AffCount:      make(map[int]int),
					}
					citationList[authorIndex] = make([]int, 0)
				}
				statistics[authorIndex].CitationShare += authorAffPair.CitationShare
				if authorIndex == firstAuthor {
					statistics[authorIndex].Citation += paperInfo.Citation
					statistics[authorIndex].CountAnalysis[paperInfo.Year-startYear].Count++
					citationList[authorIndex] = append(citationList[authorIndex], paperInfo.Citation)
				}
			}
			statistics[authorIndex].AffCount[affIndex]++
		}
	}

	refCount := GetRefCount(venueMask, startYear, endYear)
	for ref, count := range refCount {
		if count == 0 || venueMask[data.PaperInfoList[ref].VenueIndex] {
			continue
		}
		firstAuthor := data.PaperInfoList[ref].AuthorAffList[0].AuthorIndex
		for _, authorAffPair := range data.PaperInfoList[ref].AuthorAffList {
			authorIndex := authorAffPair.AuthorIndex
			affIndex := authorAffPair.AffiliationIndex
			if !affiliationMask[affIndex] || firstAuthor != authorIndex {
				continue
			}
			if statistics[authorIndex] == nil {
				statistics[authorIndex] = &AuthorStatistics{
					AuthorID:      data.AuthorList[authorIndex],
					CountAnalysis: NewCountAnalysis(startYear, endYear),
				}
			}
			statistics[authorIndex].AceScore += count
			break
		}
	}

	for i, list := range citationList {
		if list != nil {
			statistics[i].Count = len(list)
			statistics[i].AceScore += statistics[i].Citation
			statistics[i].HIndex = model.CalcHIndexByCitationList(list)
		}
	}

	authorList := make([]AuthorStatistics, 0)
	for _, v := range statistics {
		if v != nil {
			authorList = append(authorList, *v)
		}
	}

	model.Cache.SetDefault(cacheKey, authorList)
	return authorList
}
