package aceranking

import (
	"acemap/data"
	"strconv"
	"strings"
)

func GetVenueMask(VenueIDs string) ([]bool, error) {
	vs := strings.Split(VenueIDs, ",")
	venueMask := make([]bool, data.NumOfVenues)
	for _, id := range vs {
		if id != "" {
			pair := strings.Split(id, "_")
			v, err := strconv.ParseUint(pair[0], 10, 32)
			if err != nil {
				return nil, err
			}
			venueMask[data.VenueIDToIndex[data.ID(v)]] = true
		}
	}
	return venueMask, nil
}

func GetPapers(venueMask []bool, startYear, endYear int) []int {
	//cacheKey := fmt.Sprintf("GetPapers/%v/%d/%d", venueMask, startYear, endYear)
	//if value, ok := model.Cache.Get(cacheKey); ok {
	//	papers := value.([]int)
	//	return papers
	//}
	papers := make([]int, 0)
	for venueIndex, ok := range venueMask {
		if ok {
			for _, paperIndex := range data.VenuePaperMap[venueIndex] {
				year := data.PaperInfoList[paperIndex].Year
				if year <= endYear && year >= startYear {
					papers = append(papers, paperIndex)
				}
			}
		}
	}
	//model.Cache.SetDefault(cacheKey, papers)
	return papers
}

func GetRefCount(venueMask []bool, startYear, endYear int) []int {
	//cacheKey := fmt.Sprintf("GetRefCount/%v/%d/%d", venueMask, startYear, endYear)
	//if value, ok := model.Cache.Get(cacheKey); ok {
	//	refCount := value.([]int)
	//	return refCount
	//}

	refCount := make([]int, data.NumOfPapers)
	papers := GetPapers(venueMask, startYear, endYear)
	for _, paperIndex := range papers {
		for _, ref := range data.PaperInfoList[paperIndex].RefList {
			refCount[ref]++
		}
	}
	//model.Cache.SetDefault(cacheKey, refCount)
	return refCount
}

func CutFieldAnalysis(f []FieldAnalysisType) []FieldAnalysisType {
	if len(f) > 10 {
		cnt := 0
		for i := 10; i < len(f); i++ {
			cnt += f[i].Count
		}
		f = append(f[:10], FieldAnalysisType{
			FieldName: "Others",
			Count:     cnt,
		})
	}
	return f
}
