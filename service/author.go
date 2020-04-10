package service

import (
	"aceranking/cache"
	"aceranking/dao"
	"aceranking/model"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
)

func Author(req *model.AuthorReq) (*model.AuthorResp, error) {
	venueList := SplitVenueIDs(req.VenueIDs)
	venueSet := venueList.ToSet()
	affList := dao.GetAffiliationListByFilter(&req.Filter)
	affSet := affList.ToSet()

	var authorKey string
	if req.FirstAuthor == 0 {
		authorKey = "author.id"
	} else {
		authorKey = "first_author.id"
	}

	filter := bson.D{
		{authorKey, req.AuthorID},
		{"year", bson.M{"$lte": req.EndYear, "$gte": req.StartYear}},
	}
	papers := dao.FindPapers(filter)

	// 5 papers to shown in the pages.
	paperList := make([]model.AuthorPaperInfo, 0)
	fieldCountMap := make(map[string]int)

	for _, paperInfo := range papers {
		venueID := paperInfo.ConferenceSeries.ID
		venueName := paperInfo.ConferenceSeries.Abbr
		if !venueSet[venueID] {
			venueID = paperInfo.Journal.ID
			venueName = paperInfo.Journal.Name
			if !venueSet[venueID] {
				continue
			}
		}

		authorIndex := 0
		for i, aid := range paperInfo.Author.IDs {
			if aid == req.AuthorID {
				authorIndex = i
				break
			}
		}
		if !CheckAffSet(&paperInfo, authorIndex, affSet) {
			continue
		}

		field := cache.Cache(dao.GetFieldByVenueID, req.Type, venueID).(string)
		fieldCountMap[field]++
		paperList = append(paperList, model.AuthorPaperInfo{
			PaperID:  paperInfo.PaperID,
			Title:    paperInfo.Title,
			Year:     paperInfo.Year,
			Venue:    venueName,
			Citation: paperInfo.Analysis.CitationCount,
		})
	}

	sort.Slice(paperList, func(i, j int) bool {
		return paperList[i].Citation > paperList[j].Citation
	})

	// Number of rows to show AT MOST.
	end := 5
	if len(paperList) < end {
		end = len(paperList)
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

	resp := &model.AuthorResp{
		PaperList:     paperList[:end],
		FieldAnalysis: fieldAnalysis,
	}

	return resp, nil
}
