package dao

import (
	"aceranking/cache"
	"aceranking/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

//===========================================================

var paperProjection = bson.M{
	"title":             1,
	"year":              1,
	"author":            1,
	"first_author":      1,
	"affiliation":       1,
	"first_affiliation": 1,
	//"reference_list":      1,
	"analysis":          1,
	"conference_series": 1,
	"journal":           1,
}

//===========================================================

func FindPapers(filter interface{}) []model.Paper {
	amPaper := MongoClient.Database("acemap").Collection("paper")

	cursor, err := amPaper.Find(context.TODO(), filter, options.Find().SetProjection(paperProjection).SetBatchSize(150000))
	if err != nil {
		log.Println(err)
		return nil
	}

	var papers []model.Paper
	_ = cursor.All(context.TODO(), &papers)

	return papers
}

//===========================================================

func GetAffiliationListByFilter(f *model.Filter) model.IDList {
	var affiliationList model.IDList
	if f.SpecificAffID == 0 {
		affiliationList = cache.Cache(GetAffiliationListByArea, f.Area).(model.IDList)
	} else {
		affiliationList = model.IDList{f.SpecificAffID}
	}
	return affiliationList
}

//===========================================================

func GetVenuePaperCount(venueIDs model.IDList, f *model.Filter) map[model.ID]int {
	cntMap := make(map[model.ID]int)
	affiliationList := GetAffiliationListByFilter(f)

	// For cache
	var jids, cids model.IDList
	for _, vid := range venueIDs {
		cacheKey := fmt.Sprintf("GetVenuePaperCount/%v/%v", f, vid)
		if c, ok := cache.Get(cacheKey); ok {
			cntMap[vid] = c.(int)
		} else {
			if vid.GetIDType() == "conference" {
				cids = append(cids, vid)
			} else {
				jids = append(jids, vid)
			}
		}
	}

	wg := sync.WaitGroup{}
	ch := make(chan model.VenuePaperCount, 100)

	taskForOneType := func(vids model.IDList) {
		if len(vids) == 0 {
			wg.Done()
			return
		}

		var affKey string
		if f.FirstAuthor == 0 {
			affKey = "affiliation.id"
		} else {
			affKey = "first_affiliation.id"
		}

		var venueKey string
		if vids[0].GetIDType() == "conference" {
			venueKey = "conference_series.id"
		} else {
			venueKey = "journal.id"
		}

		pipeline := bson.A{
			bson.M{
				"$match": bson.D{
					{venueKey, bson.M{"$in": vids}},
					{"year", bson.M{"$lte": f.EndYear, "$gte": f.StartYear}},
					{affKey, bson.M{"$in": affiliationList}},
				},
			},
			bson.M{
				"$group": bson.D{
					{"_id", "$" + venueKey},
					{"count", bson.M{"$sum": 1}},
				},
			},
		}
		amPaper := MongoClient.Database("acemap").Collection("paper")
		cursor, err := amPaper.Aggregate(context.TODO(), pipeline, nil)
		if err != nil {
			wg.Done()
			return
		}

		for cursor.Next(context.TODO()) {
			var c model.VenuePaperCount
			_ = cursor.Decode(&c)
			ch <- c
		}
		wg.Done()
	}

	wg.Add(2)
	go taskForOneType(cids)
	go taskForOneType(jids)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		cntMap[c.VenueID] = c.Count
		cacheKey := fmt.Sprintf("GetVenuePaperCount/%v/%v", f, c.VenueID)
		cache.Set(cacheKey, c.Count)
	}
	return cntMap
}

//===========================================================

func GetPapersByFilter(venueIDs model.IDList, f *model.Filter) []model.Paper {
	affiliationList := GetAffiliationListByFilter(f)
	ch := make(chan []model.Paper)
	wg := sync.WaitGroup{}

	taskForOneType := func(vids model.IDList) {
		if len(vids) == 0 {
			wg.Done()
			return
		}

		var affKey string
		if f.FirstAuthor == 0 {
			affKey = "affiliation.id"
		} else {
			affKey = "first_affiliation.id"
		}

		var venueKey string
		if vids[0].GetIDType() == "conference" {
			venueKey = "conference_series.id"
		} else {
			venueKey = "journal.id"
		}

		filter := bson.D{
			{venueKey, bson.M{"$in": vids}},
			{"year", bson.M{"$lte": f.EndYear, "$gte": f.StartYear}},
			{affKey, bson.M{"$in": affiliationList}},
		}
		papers := FindPapers(filter)
		ch <- papers
		wg.Done()
	}

	var jids, cids model.IDList
	for _, vid := range venueIDs {
		if vid.GetIDType() == "conference" {
			cids = append(cids, vid)
		} else {
			jids = append(jids, vid)
		}
	}

	wg.Add(2)
	go taskForOneType(cids)
	go taskForOneType(jids)

	go func() {
		wg.Wait()
		close(ch)
	}()

	var papers []model.Paper
	for p := range ch {
		papers = append(papers, p...)
	}
	return papers
}

//===========================================================

func GetRefCount(venueIDs model.IDList) map[model.ID]int {
	cntMap := make(map[model.ID]int, 100)

	var jids, cids model.IDList
	for _, vid := range venueIDs {
		if vid.GetIDType() == "conference" {
			cids = append(cids, vid)
		} else {
			jids = append(jids, vid)
		}
	}

	wg := sync.WaitGroup{}
	ch := make(chan model.RefCount)

	taskForOneType := func(vids model.IDList) {
		if len(vids) == 0 {
			wg.Done()
			return
		}

		var venueKey string
		if vids[0].GetIDType() == "conference" {
			venueKey = "conference_series.id"
		} else {
			venueKey = "journal.id"
		}

		pipeline := bson.A{
			bson.M{
				"$match": bson.D{
					{venueKey, bson.M{"$in": vids}},
				},
			},
			bson.M{
				"$unwind": "$reference_list",
			},
			bson.M{
				"$group": bson.D{
					{"_id", "$reference_list"},
					{"count", bson.M{"$sum": 1}},
				},
			},
			bson.M{
				"$match": bson.D{
					{"count", bson.M{"$in": vids}},
				},
			},
		}
		amPaper := MongoClient.Database("acemap").Collection("paper")
		cursor, err := amPaper.Aggregate(context.TODO(), pipeline, nil)
		if err != nil {
			wg.Done()
			return
		}

		for cursor.Next(context.TODO()) {
			var c model.RefCount
			_ = cursor.Decode(&c)
			ch <- c
		}
		wg.Done()
	}

	wg.Add(2)
	go taskForOneType(cids)
	go taskForOneType(jids)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		cntMap[c.ReferenceID] = c.Count
	}
	return cntMap

}
