package service

import (
	"aceranking/cache"
	"aceranking/dao"
	"aceranking/model"
)

func AffiliationList(req *model.AffiliationListReq) (*model.AffiliationListResp, error) {
	orderby := req.OrderBy
	req.OrderBy = 0
	resp := cache.Cache(_affiliationList, req).(*model.AffiliationListResp)
	affList := []model.AffiliationStatistics(*resp)
	model.Sort(affList, orderby)
	maxRow := 50
	if len(affList) < maxRow {
		maxRow = len(affList)
	}
	affList = affList[:maxRow]
	return resp, nil
}

func _affiliationList(req *model.AffiliationListReq) (*model.AffiliationListResp, error) {
	venueList := SplitVenueIDs(req.VenueIDs)
	//venueSet := venueList.ToSet()
	affIDList := dao.GetAffiliationListByFilter(&req.Filter)
	affSet := affIDList.ToSet()

	papers := dao.GetPapersByFilter(venueList, &req.Filter)

	affStat := make(map[model.ID]*model.AffiliationStatistics)
	citationList := make(map[model.ID][]int)
	for _, paperInfo := range papers {
		if req.FirstAuthor == 0 {
			for i, affID := range paperInfo.Affiliation.IDs {
				if !affSet[affID] {
					continue
				}
				if affStat[affID] == nil {
					affStat[affID] = &model.AffiliationStatistics{
						AffiliationID:   affID,
						AffiliationName: paperInfo.Affiliation.Names[i],
						Abbreviation:    paperInfo.Affiliation.Abbrs[i],
					}
					citationList[affID] = make([]int, 0)
				}
				stat := affStat[affID]
				stat.Citation += paperInfo.Analysis.CitationCount
				stat.T2Citation += paperInfo.Analysis.T2CitationCount
				citationList[affID] = append(citationList[affID], paperInfo.Analysis.CitationCount)
			}
			for _, idxList := range paperInfo.Author.AffIndex {
				for _, idx := range idxList {
					affID := paperInfo.Affiliation.IDs[idx]
					if affSet[affID] {
						affStat[affID].CitationShare += float64(paperInfo.Analysis.CitationCount) / float64(len(paperInfo.Author.IDs))
					}
				}
			}
		} else {
			affID := paperInfo.FirstAffiliation.ID
			if affStat[affID] == nil {
				affStat[affID] = &model.AffiliationStatistics{
					AffiliationID:   affID,
					AffiliationName: paperInfo.FirstAffiliation.Name,
					Abbreviation:    paperInfo.FirstAffiliation.Abbr,
				}
				citationList[affID] = make([]int, 0)
			}
			stat := affStat[affID]
			stat.Citation += paperInfo.Analysis.CitationCount
			stat.T2Citation += paperInfo.Analysis.T2CitationCount
			stat.CitationShare += float64(paperInfo.Analysis.CitationCount)
			citationList[affID] = append(citationList[affID], paperInfo.Analysis.CitationCount)
		}
	}

	for affID, list := range citationList {
		affStat[affID].Count = len(list)
		affStat[affID].AceScore = int(1.23*float64(affStat[affID].Citation)) + 333
		affStat[affID].HIndex = CalcHIndexByCitationList(list)
	}

	maxRow := 50
	affIDSet := make(map[model.ID]bool)

	allAffList := make([]model.AffiliationStatistics, 0, len(affStat))
	for _, v := range affStat {
		allAffList = append(allAffList, *v)
	}
	if len(allAffList) < maxRow {
		maxRow = len(allAffList)
	}

	for i := 1; i <= 6; i++ {
		if i == 5 {
			continue
		}
		model.Sort(allAffList, i)
		for j := 0; j < maxRow; j++ {
			affIDSet[allAffList[j].AffiliationID] = true
		}
	}

	affList := make(model.AffiliationListResp, 0)
	for _, v := range allAffList {
		if affIDSet[v.AffiliationID] {
			affList = append(affList, v)
		}
	}

	model.Sort(affList, req.OrderBy)
	affList = affList[:maxRow]

	return &affList, nil
}
