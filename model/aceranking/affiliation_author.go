package aceranking

import (
	"acemap/data"
	"acemap/model/dao"
	"sort"
)

type AuthorPaperInfo struct {
	PaperID  data.ID `json:"paper_id"`
	Title    string  `json:"title"`
	Year     int     `json:"year"`
	Venue    string  `json:"venue"`
	Citation int     `json:"citation"`
}

type AuthorInfo struct {
	PaperList     []AuthorPaperInfo   `json:"paper_list"`
	FieldAnalysis []FieldAnalysisType `json:"field_analysis"`
}

func GetAffAuthorInfoOnAllAuthors(typ string, venueMask []bool, startYear, endYear, affIndex, authorIndex int) AuthorInfo {
	paperList := make([]AuthorPaperInfo, 0)
	fieldCountMap := make(map[string]int)

	for _, paperIndex := range data.AffiliationPaperList[affIndex] {
		paperInfo := data.PaperInfoList[paperIndex]
		if venueMask[paperInfo.VenueIndex] && paperInfo.Year <= endYear && paperInfo.Year >= startYear {
			for _, authorAffPair := range paperInfo.AuthorAffList {
				if affIndex == authorAffPair.AffiliationIndex && authorIndex == authorAffPair.AuthorIndex {
					venueID := data.VenueList[paperInfo.VenueIndex]
					field := dao.GetFieldByVenueID(typ, venueID)
					fieldCountMap[field]++
					paperList = append(paperList, AuthorPaperInfo{
						PaperID:  data.PaperList[paperIndex],
						Title:    "",
						Year:     paperInfo.Year,
						Venue:    dao.GetVenueNameByID(venueID),
						Citation: paperInfo.Citation,
					})
					break
				}
			}
		}
	}

	sort.Slice(paperList, func(i, j int) bool {
		return paperList[i].Citation > paperList[j].Citation
	})

	// Number of rows to show AT MOST.
	end := 5
	if len(paperList) < end {
		end = len(paperList)
	}

	for i := 0; i < end; i++ {
		paperList[i].Title = dao.GetPaperTitleByID(paperList[i].PaperID)
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

	affAuthorInfo := AuthorInfo{
		PaperList:     paperList[:end],
		FieldAnalysis: fieldAnalysis,
	}

	return affAuthorInfo
}

func GetAffAuthorInfoOnFirstAuthorWeak(typ string, venueMask []bool, startYear, endYear, affIndex, authorIndex int) AuthorInfo {
	paperList := make([]AuthorPaperInfo, 0)
	fieldCountMap := make(map[string]int)

	for _, paperIndex := range data.AffiliationPaperList[affIndex] {
		paperInfo := data.PaperInfoList[paperIndex]
		if ok := isAffiliationFirst(affIndex, paperIndex); !ok {
			continue
		}
		if venueMask[paperInfo.VenueIndex] && paperInfo.Year <= endYear && paperInfo.Year >= startYear {
			for _, authorAffPair := range paperInfo.AuthorAffList {
				if affIndex == authorAffPair.AffiliationIndex && authorIndex == authorAffPair.AuthorIndex {
					venueID := data.VenueList[paperInfo.VenueIndex]
					field := dao.GetFieldByVenueID(typ, venueID)
					fieldCountMap[field]++
					paperList = append(paperList, AuthorPaperInfo{
						PaperID:  data.PaperList[paperIndex],
						Title:    "",
						Year:     paperInfo.Year,
						Venue:    dao.GetVenueNameByID(venueID),
						Citation: paperInfo.Citation,
					})
					break
				}
			}
		}
	}

	sort.Slice(paperList, func(i, j int) bool {
		return paperList[i].Citation > paperList[j].Citation
	})

	// Number of rows to show AT MOST.
	end := 5
	if len(paperList) < end {
		end = len(paperList)
	}

	for i := 0; i < end; i++ {
		paperList[i].Title = dao.GetPaperTitleByID(paperList[i].PaperID)
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

	affAuthorInfo := AuthorInfo{
		PaperList:     paperList[:end],
		FieldAnalysis: fieldAnalysis,
	}

	return affAuthorInfo
}

func GetAffAuthorInfoOnFirstAuthorStrong(typ string, venueMask []bool, startYear, endYear, affIndex, authorIndex int) AuthorInfo {
	paperList := make([]AuthorPaperInfo, 0)
	fieldCountMap := make(map[string]int)

	for _, paperIndex := range data.AffiliationPaperList[affIndex] {
		paperInfo := data.PaperInfoList[paperIndex]
		if ok := isAffiliationAuthorFirst(affIndex, authorIndex, paperIndex); !ok {
			continue
		}
		if venueMask[paperInfo.VenueIndex] && paperInfo.Year <= endYear && paperInfo.Year >= startYear {
			for _, authorAffPair := range paperInfo.AuthorAffList {
				if affIndex == authorAffPair.AffiliationIndex && authorIndex == authorAffPair.AuthorIndex {
					venueID := data.VenueList[paperInfo.VenueIndex]
					field := dao.GetFieldByVenueID(typ, venueID)
					fieldCountMap[field]++
					paperList = append(paperList, AuthorPaperInfo{
						PaperID:  data.PaperList[paperIndex],
						Title:    "",
						Year:     paperInfo.Year,
						Venue:    dao.GetVenueNameByID(venueID),
						Citation: paperInfo.Citation,
					})
					break
				}
			}
		}
	}

	sort.Slice(paperList, func(i, j int) bool {
		return paperList[i].Citation > paperList[j].Citation
	})

	// Number of rows to show AT MOST.
	end := 5
	if len(paperList) < end {
		end = len(paperList)
	}

	for i := 0; i < end; i++ {
		paperList[i].Title = dao.GetPaperTitleByID(paperList[i].PaperID)
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

	affAuthorInfo := AuthorInfo{
		PaperList:     paperList[:end],
		FieldAnalysis: fieldAnalysis,
	}

	return affAuthorInfo
}

func isAffiliationAuthorFirst(affIndex, authorIndex, paperIndex int) bool {
	paperInfo := data.PaperInfoList[paperIndex]
	firstAuthor := paperInfo.AuthorAffList[0].AuthorIndex
	if firstAuthor != authorIndex {
		return false
	}
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
