package service

import (
	"aceranking/cache"
	"aceranking/dao"
	"aceranking/model"
	"sort"
)

func AuthorList(req *model.AuthorListReq) (*model.AuthorListResp, error) {
	orderby := req.OrderBy
	req.OrderBy = 0
	resp := cache.Cache(_authorList, req).(*model.AuthorListResp)
	authorList := resp.AuthorList
	model.Sort(authorList, orderby)
	maxRow := 50
	if len(authorList) < maxRow {
		maxRow = len(authorList)
	}
	authorList = authorList[:maxRow]
	resp.AuthorList = authorList
	return resp, nil
}

func _authorList(req *model.AuthorListReq) (*model.AuthorListResp, error) {
	venueList := SplitVenueIDs(req.VenueIDs)
	//venueSet := venueList.ToSet()
	affList := dao.GetAffiliationListByFilter(&req.Filter)
	affSet := affList.ToSet()

	papers := dao.GetPapersByFilter(venueList, &req.Filter)

	authorStat := make(map[model.ID]*model.AuthorStatistics)
	citationList := make(map[model.ID][]int)
	countAnalysis := model.NewCountAnalysis(req.StartYear, req.EndYear)
	fieldCountMap := make(map[string]int)
	affName := make(map[model.ID]string)
	affAbbr := make(map[model.ID]string)
	for _, paperInfo := range papers {
		if req.SpecificAffID != 0 {
			countAnalysis[paperInfo.Year-req.StartYear].Count++
			venueID := paperInfo.Journal.ID
			if venueID == 0 {
				venueID = paperInfo.ConferenceSeries.ID
			}
			field := cache.Cache(dao.GetFieldByVenueID, req.Type, venueID).(string)
			fieldCountMap[field]++
		}

		if req.FirstAuthor == 0 {
			for i, authorID := range paperInfo.Author.IDs {
				if !CheckAffSet(&paperInfo, i, affSet) {
					continue
				}
				if authorStat[authorID] == nil {
					authorStat[authorID] = &model.AuthorStatistics{
						AuthorID:      authorID,
						AuthorName:    paperInfo.Author.Names[i],
						CountAnalysis: model.NewCountAnalysis(req.StartYear, req.EndYear),
						AffCount:      make(map[model.ID]int),
					}
					citationList[authorID] = make([]int, 0)
				}
				stat := authorStat[authorID]
				stat.CitationShare += float64(paperInfo.Analysis.CitationCount) / float64(len(paperInfo.Author.IDs))
				stat.Citation += paperInfo.Analysis.CitationCount
				stat.T2Citation += paperInfo.Analysis.T2CitationCount
				stat.CountAnalysis[paperInfo.Year-req.StartYear].Count++
				citationList[authorID] = append(citationList[authorID], paperInfo.Analysis.CitationCount)
				for _, idx := range paperInfo.Author.AffIndex[i] {
					affID := paperInfo.Affiliation.IDs[idx]
					stat.AffCount[affID]++
					affName[affID] = paperInfo.Affiliation.Names[idx]
					affAbbr[affID] = paperInfo.Affiliation.Abbrs[idx]
				}
			}
		} else {
			authorID := paperInfo.FirstAuthor.ID
			if authorStat[authorID] == nil {
				authorStat[authorID] = &model.AuthorStatistics{
					AuthorID:      authorID,
					AuthorName:    paperInfo.FirstAuthor.Name,
					CountAnalysis: model.NewCountAnalysis(req.StartYear, req.EndYear),
					AffCount:      make(map[model.ID]int),
				}
				citationList[authorID] = make([]int, 0)
			}
			stat := authorStat[authorID]
			stat.CitationShare += float64(paperInfo.Analysis.CitationCount)
			stat.Citation += paperInfo.Analysis.CitationCount
			stat.T2Citation += paperInfo.Analysis.T2CitationCount
			stat.CountAnalysis[paperInfo.Year-req.StartYear].Count++
			citationList[authorID] = append(citationList[authorID], paperInfo.Analysis.CitationCount)
			stat.AffCount[paperInfo.FirstAffiliation.ID]++
			affName[paperInfo.FirstAffiliation.ID] = paperInfo.FirstAffiliation.Name
			affAbbr[paperInfo.FirstAffiliation.ID] = paperInfo.FirstAffiliation.Abbr
		}
	}

	for authorID, list := range citationList {
		authorStat[authorID].Count = len(list)
		authorStat[authorID].AceScore = int(1.23*float64(authorStat[authorID].Citation)) + 333
		authorStat[authorID].HIndex = CalcHIndexByCitationList(list)
	}

	maxRow := 50
	authorIDSet := make(map[model.ID]bool)

	allAuthorList := make([]model.AuthorStatistics, 0, len(authorStat))
	for _, v := range authorStat {
		allAuthorList = append(allAuthorList, *v)
	}
	if len(allAuthorList) < maxRow {
		maxRow = len(allAuthorList)
	}

	for i := 1; i <= 6; i++ {
		if i == 5 {
			continue
		}
		model.Sort(allAuthorList, i)
		for j := 0; j < maxRow; j++ {
			authorIDSet[allAuthorList[j].AuthorID] = true
		}
	}

	authorList := make([]model.AuthorStatistics, 0)
	for _, v := range allAuthorList {
		if authorIDSet[v.AuthorID] {
			authorList = append(authorList, v)
		}
	}

	//t1 := time.Now()
	//refCnt := dao.GetRefCount(venueList)
	//fmt.Println(time.Since(t1), len(refCnt))

	//calcAceScoreForOneVenue := func(authorStat *AuthorStatistics) {
	//	authorID := authorStat.AuthorID
	//	var filter bson.M
	//	if firstAuthor == 2 {
	//		filter = bson.M{"first_author.id": authorID}
	//	} else {
	//		filter = bson.M{"author.id": authorID}
	//	}
	//	papers := nosql.FindPapers(filter)
	//
	//	for _, paperInfo := range papers {
	//		if paperInfo.ConferenceSeries.ID == venueID || paperInfo.Journal.ID == venueID {
	//			continue
	//		}
	//		// ignore case: first_author = 1
	//		authorStat.AceScore += refCount[paperInfo.PaperID]
	//	}
	//}

	for i := range authorList {
		//calcAceScoreForOneVenue(&authorList[i])

		authorList[i].AffList = make([]model.AuthorAffInfo, 0)
		for affID, count := range authorList[i].AffCount {
			authorList[i].AffList = append(authorList[i].AffList, model.AuthorAffInfo{
				AffiliationID:   affID,
				Abbreviation:    affAbbr[affID],
				AffiliationName: affName[affID],
				Count:           count,
			})
		}
		sort.Slice(authorList[i].AffList, func(j, k int) bool {
			return authorList[i].AffList[j].Count > authorList[i].AffList[k].Count
		})
		if l := len(authorList[i].AffList); l == 0 {
			continue
		} else if l >= 2 && authorList[i].AffList[1].Count >= 10 {
			authorList[i].AffList = authorList[i].AffList[:2]
		} else {
			authorList[i].AffList = authorList[i].AffList[:1]
		}
	}

	fieldAnalysis := make([]model.FieldAnalysisType, 0, len(fieldCountMap))
	for field, cnt := range fieldCountMap {
		fieldAnalysis = append(fieldAnalysis, model.FieldAnalysisType{
			FieldName: field,
			Count:     cnt,
		})
	}
	sort.Slice(fieldAnalysis, func(i, j int) bool {
		return fieldAnalysis[i].Count > fieldAnalysis[j].Count
	})
	fieldAnalysis = CutFieldAnalysis(fieldAnalysis)

	resp := &model.AuthorListResp{
		AuthorList:    authorList,
		CountAnalysis: countAnalysis,
		FieldAnalysis: fieldAnalysis,
		HasFigure:     false,
	}

	return resp, nil
}
