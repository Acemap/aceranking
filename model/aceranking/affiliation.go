package aceranking

import (
	"acemap/data"
	"acemap/model"
	"acemap/model/dao"
	"sort"
)

type AffAuthorStatistics struct {
	AuthorID      data.ID             `json:"author_id"`
	AuthorName    string              `json:"author_name"`
	CountAnalysis []CountAnalysisType `json:"count_analysis"`
	Statistics
}

type CountAnalysisType struct {
	Year  int `json:"year"`
	Count int `json:"count"`
}

func NewCountAnalysis(startYear, endYear int) []CountAnalysisType {
	c := make([]CountAnalysisType, endYear-startYear+1)
	for i := range c {
		c[i].Year = startYear + i
	}
	return c
}

type FieldAnalysisType struct {
	FieldName string `json:"name"`
	Count     int    `json:"count"`
}

type AffiliationDetails struct {
	AuthorList    []AffAuthorStatistics `json:"author_list"`
	CountAnalysis []CountAnalysisType   `json:"count_analysis"`
	FieldAnalysis []FieldAnalysisType   `json:"field_analysis"`
}

func GetAffDetailsOnAllAuthors(typ string, venueMask []bool, startYear, endYear, affIndex int) AffiliationDetails {
	statistics := make(map[int]*AffAuthorStatistics)
	citationList := make(map[int][]int)
	countAnalysis := NewCountAnalysis(startYear, endYear)
	fieldCountMap := make(map[string]int)
	refCount := GetRefCount(venueMask, startYear, endYear)

	for _, paperIndex := range data.AffiliationPaperList[affIndex] {
		paperInfo := data.PaperInfoList[paperIndex]
		if !venueMask[data.PaperInfoList[paperIndex].VenueIndex] {
			for _, authorAffPair := range paperInfo.AuthorAffList {
				if affIndex == authorAffPair.AffiliationIndex {
					authorIndex := authorAffPair.AuthorIndex
					if statistics[authorIndex] == nil {
						statistics[authorIndex] = &AffAuthorStatistics{
							AuthorID:      data.AuthorList[authorIndex],
							CountAnalysis: NewCountAnalysis(startYear, endYear),
						}
					}
					statistics[authorIndex].AceScore += refCount[paperIndex]
				}
			}
		} else {
			if paperInfo.Year > endYear || paperInfo.Year < startYear {
				continue
			}
			field := dao.GetFieldByVenueID(typ, data.VenueList[paperInfo.VenueIndex])
			countAnalysis[paperInfo.Year-startYear].Count++
			fieldCountMap[field]++

			for _, authorAffPair := range paperInfo.AuthorAffList {
				if affIndex != authorAffPair.AffiliationIndex {
					continue
				}
				authorIndex := authorAffPair.AuthorIndex

				if statistics[authorIndex] == nil {
					statistics[authorIndex] = &AffAuthorStatistics{
						AuthorID:      data.AuthorList[authorIndex],
						CountAnalysis: NewCountAnalysis(startYear, endYear),
					}
					citationList[authorIndex] = make([]int, 0)
				}

				statistics[authorIndex].Citation += paperInfo.Citation
				statistics[authorIndex].CitationShare += authorAffPair.CitationShare
				statistics[authorIndex].CountAnalysis[paperInfo.Year-startYear].Count++
				citationList[authorIndex] = append(citationList[authorIndex], paperInfo.Citation)
			}
		}
	}

	for i, list := range citationList {
		statistics[i].Count = len(list)
		statistics[i].AceScore += statistics[i].Citation
		statistics[i].HIndex = model.CalcHIndexByCitationList(list)
	}

	authorList := make([]AffAuthorStatistics, 0, len(statistics))
	for _, v := range statistics {
		authorList = append(authorList, *v)
	}

	fieldAnalysis := make([]FieldAnalysisType, 0, len(fieldCountMap))
	for field, cnt := range fieldCountMap {
		fieldAnalysis = append(fieldAnalysis, FieldAnalysisType{
			FieldName: field,
			Count:     cnt,
		})
	}
	sort.Slice(fieldAnalysis, func(i, j int) bool {
		return fieldAnalysis[i].Count > fieldAnalysis[j].Count
	})
	fieldAnalysis = CutFieldAnalysis(fieldAnalysis)

	details := AffiliationDetails{
		AuthorList:    authorList,
		CountAnalysis: countAnalysis,
		FieldAnalysis: fieldAnalysis,
	}

	return details
}

func GetAffDetailsOnFirstAuthorWeak(typ string, venueMask []bool, startYear, endYear, affIndex int) AffiliationDetails {
	statistics := make(map[int]*AffAuthorStatistics)
	citationList := make(map[int][]int)
	countAnalysis := NewCountAnalysis(startYear, endYear)
	fieldCountMap := make(map[string]int)
	refCount := GetRefCount(venueMask, startYear, endYear)

	for _, paperIndex := range data.AffiliationPaperList[affIndex] {
		paperInfo := data.PaperInfoList[paperIndex]
		ok := isAffiliationFirst(affIndex, paperIndex)
		if !venueMask[data.PaperInfoList[paperIndex].VenueIndex] {
			if !ok {
				continue
			}
			for _, authorAffPair := range paperInfo.AuthorAffList {
				if affIndex == authorAffPair.AffiliationIndex {
					authorIndex := authorAffPair.AuthorIndex
					if statistics[authorIndex] == nil {
						statistics[authorIndex] = &AffAuthorStatistics{
							AuthorID:      data.AuthorList[authorIndex],
							CountAnalysis: NewCountAnalysis(startYear, endYear),
						}
					}
					statistics[authorIndex].AceScore += refCount[paperIndex]
				}
			}
		} else {
			if paperInfo.Year > endYear || paperInfo.Year < startYear {
				continue
			}
			if ok {
				field := dao.GetFieldByVenueID(typ, data.VenueList[paperInfo.VenueIndex])
				countAnalysis[paperInfo.Year-startYear].Count++
				fieldCountMap[field]++
			}

			for _, authorAffPair := range paperInfo.AuthorAffList {
				if affIndex != authorAffPair.AffiliationIndex {
					continue
				}
				authorIndex := authorAffPair.AuthorIndex

				if statistics[authorIndex] == nil {
					statistics[authorIndex] = &AffAuthorStatistics{
						AuthorID:      data.AuthorList[authorIndex],
						CountAnalysis: NewCountAnalysis(startYear, endYear),
					}
					citationList[authorIndex] = make([]int, 0)
				}
				statistics[authorIndex].CitationShare += authorAffPair.CitationShare

				if ok {
					statistics[authorIndex].Citation += paperInfo.Citation
					statistics[authorIndex].CountAnalysis[paperInfo.Year-startYear].Count++
					citationList[authorIndex] = append(citationList[authorIndex], paperInfo.Citation)
				}
			}
		}
	}

	for i, list := range citationList {
		statistics[i].Count = len(list)
		statistics[i].AceScore += statistics[i].Citation
		statistics[i].HIndex = model.CalcHIndexByCitationList(list)
	}

	authorList := make([]AffAuthorStatistics, 0, len(statistics))
	for _, v := range statistics {
		authorList = append(authorList, *v)
	}

	fieldAnalysis := make([]FieldAnalysisType, 0, len(fieldCountMap))
	for field, cnt := range fieldCountMap {
		fieldAnalysis = append(fieldAnalysis, FieldAnalysisType{
			FieldName: field,
			Count:     cnt,
		})
	}
	sort.Slice(fieldAnalysis, func(i, j int) bool {
		return fieldAnalysis[i].Count > fieldAnalysis[j].Count
	})
	fieldAnalysis = CutFieldAnalysis(fieldAnalysis)

	details := AffiliationDetails{
		AuthorList:    authorList,
		CountAnalysis: countAnalysis,
		FieldAnalysis: fieldAnalysis,
	}

	return details
}

func GetAffDetailsOnFirstAuthorStrong(typ string, venueMask []bool, startYear, endYear, affIndex int) AffiliationDetails {
	statistics := make(map[int]*AffAuthorStatistics)
	citationList := make(map[int][]int)
	countAnalysis := NewCountAnalysis(startYear, endYear)
	fieldCountMap := make(map[string]int)
	refCount := GetRefCount(venueMask, startYear, endYear)

	for _, paperIndex := range data.AffiliationPaperList[affIndex] {
		paperInfo := data.PaperInfoList[paperIndex]
		firstAuthor := paperInfo.AuthorAffList[0].AuthorIndex
		ok := isAffiliationFirst(affIndex, paperIndex)
		if !venueMask[data.PaperInfoList[paperIndex].VenueIndex] {
			for _, authorAffPair := range paperInfo.AuthorAffList {
				if authorAffPair.AuthorIndex != firstAuthor {
					break
				}
				if affIndex == authorAffPair.AffiliationIndex {
					authorIndex := authorAffPair.AuthorIndex
					if statistics[authorIndex] == nil {
						statistics[authorIndex] = &AffAuthorStatistics{
							AuthorID:      data.AuthorList[authorIndex],
							CountAnalysis: NewCountAnalysis(startYear, endYear),
						}
					}
					statistics[authorIndex].AceScore += refCount[paperIndex]
				}
			}
		} else {
			if paperInfo.Year > endYear || paperInfo.Year < startYear {
				continue
			}
			if ok {
				field := dao.GetFieldByVenueID(typ, data.VenueList[paperInfo.VenueIndex])
				countAnalysis[paperInfo.Year-startYear].Count++
				fieldCountMap[field]++
			}

			for _, authorAffPair := range paperInfo.AuthorAffList {
				if affIndex != authorAffPair.AffiliationIndex {
					continue
				}
				authorIndex := authorAffPair.AuthorIndex
				if statistics[authorIndex] == nil {
					statistics[authorIndex] = &AffAuthorStatistics{
						AuthorID:      data.AuthorList[authorIndex],
						CountAnalysis: NewCountAnalysis(startYear, endYear),
					}
					citationList[authorIndex] = make([]int, 0)
				}
				statistics[authorIndex].CitationShare += authorAffPair.CitationShare

				if authorAffPair.AuthorIndex == firstAuthor {
					statistics[authorIndex].Citation += paperInfo.Citation
					statistics[authorIndex].CountAnalysis[paperInfo.Year-startYear].Count++
					citationList[authorIndex] = append(citationList[authorIndex], paperInfo.Citation)
				}
			}
		}
	}

	for i, list := range citationList {
		statistics[i].Count = len(list)
		statistics[i].AceScore += statistics[i].Citation
		statistics[i].HIndex = model.CalcHIndexByCitationList(list)
	}

	authorList := make([]AffAuthorStatistics, 0, len(statistics))
	for _, v := range statistics {
		authorList = append(authorList, *v)
	}

	fieldAnalysis := make([]FieldAnalysisType, 0, len(fieldCountMap))
	for field, cnt := range fieldCountMap {
		fieldAnalysis = append(fieldAnalysis, FieldAnalysisType{
			FieldName: field,
			Count:     cnt,
		})
	}
	sort.Slice(fieldAnalysis, func(i, j int) bool {
		return fieldAnalysis[i].Count > fieldAnalysis[j].Count
	})
	fieldAnalysis = CutFieldAnalysis(fieldAnalysis)

	details := AffiliationDetails{
		AuthorList:    authorList,
		CountAnalysis: countAnalysis,
		FieldAnalysis: fieldAnalysis,
	}

	return details
}
