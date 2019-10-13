package aceranking

import (
	"acemap/data"
	"acemap/model/dao"
	"sort"
)

func GetAuthorInfoOnAllAuthors(typ, area string, venueMask []bool, startYear, endYear, authorIndex int) AuthorInfo {
	affiliationMask := dao.GetAffiliationMaskByArea(area)
	allPapers := GetPapers(venueMask, startYear, endYear)
	paperList := make([]AuthorPaperInfo, 0)
	fieldCountMap := make(map[string]int)

	for _, paperIndex := range allPapers {
		paperInfo := data.PaperInfoList[paperIndex]
		for _, authorAffPair := range paperInfo.AuthorAffList {
			if affiliationMask[authorAffPair.AffiliationIndex] && authorIndex == authorAffPair.AuthorIndex {
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

	authorInfo := AuthorInfo{
		PaperList:     paperList[:end],
		FieldAnalysis: fieldAnalysis,
	}

	return authorInfo
}

func GetAuthorInfoOnFirstAuthorWeak(typ, area string, venueMask []bool, startYear, endYear, authorIndex int) AuthorInfo {
	affiliationMask := dao.GetAffiliationMaskByArea(area)
	allPapers := GetPapers(venueMask, startYear, endYear)
	paperList := make([]AuthorPaperInfo, 0)
	fieldCountMap := make(map[string]int)

	for _, paperIndex := range allPapers {
		paperInfo := data.PaperInfoList[paperIndex]
		firstAff := paperInfo.AuthorAffList[0].AffiliationIndex
		for _, authorAffPair := range paperInfo.AuthorAffList {
			if affiliationMask[authorAffPair.AffiliationIndex] && authorIndex == authorAffPair.AuthorIndex && firstAff == authorAffPair.AffiliationIndex {
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

	authorInfo := AuthorInfo{
		PaperList:     paperList[:end],
		FieldAnalysis: fieldAnalysis,
	}

	return authorInfo
}

func GetAuthorInfoOnFirstAuthorStrong(typ, area string, venueMask []bool, startYear, endYear, authorIndex int) AuthorInfo {
	affiliationMask := dao.GetAffiliationMaskByArea(area)
	allPapers := GetPapers(venueMask, startYear, endYear)
	paperList := make([]AuthorPaperInfo, 0)
	fieldCountMap := make(map[string]int)

	for _, paperIndex := range allPapers {
		paperInfo := data.PaperInfoList[paperIndex]
		firstAuthor := paperInfo.AuthorAffList[0].AuthorIndex
		if authorIndex != firstAuthor {
			continue
		}
		for _, authorAffPair := range paperInfo.AuthorAffList {
			if affiliationMask[authorAffPair.AffiliationIndex] && authorIndex == authorAffPair.AuthorIndex {
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

	authorInfo := AuthorInfo{
		PaperList:     paperList[:end],
		FieldAnalysis: fieldAnalysis,
	}

	return authorInfo
}
