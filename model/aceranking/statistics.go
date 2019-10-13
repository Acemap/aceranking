package aceranking

import (
	"reflect"
	"sort"
)

type Statistics struct {
	Count         int     `json:"count"`
	Citation      int     `json:"citation"`
	CitationShare float64 `json:"cit_share"`
	HIndex        int     `json:"hindex"`
	AceScore      int     `json:"ace_score"`
}

// Order by. 1: count, 2: citation, 3: citation_share, 4: H-index, 5: AceScore
func Sort(slice interface{}, orderBy int) {
	v := reflect.ValueOf(slice)
	switch orderBy {
	case 1:
		sort.Slice(slice, func(i, j int) bool {
			vi := v.Index(i).FieldByName("Count").Interface().(int)
			vj := v.Index(j).FieldByName("Count").Interface().(int)
			return vi > vj
		})
	case 2:
		sort.Slice(slice, func(i, j int) bool {
			vi := v.Index(i).FieldByName("Citation").Interface().(int)
			vj := v.Index(j).FieldByName("Citation").Interface().(int)
			return vi > vj
		})
	case 3:
		sort.Slice(slice, func(i, j int) bool {
			vi := v.Index(i).FieldByName("CitationShare").Interface().(float64)
			vj := v.Index(j).FieldByName("CitationShare").Interface().(float64)
			return vi > vj
		})
	case 4:
		sort.Slice(slice, func(i, j int) bool {
			vi := v.Index(i).FieldByName("HIndex").Interface().(int)
			vj := v.Index(j).FieldByName("HIndex").Interface().(int)
			return vi > vj
		})
	case 5:
		sort.Slice(slice, func(i, j int) bool {
			vi := v.Index(i).FieldByName("AceScore").Interface().(int)
			vj := v.Index(j).FieldByName("AceScore").Interface().(int)
			return vi > vj
		})
	}
}
