package model

import "sort"

type Statistics struct {
	Count         int     `json:"count"`
	Citation      int     `json:"citation"`
	T2Citation    int     `json:"t2_citation"`
	CitationShare float64 `json:"cit_share"`
	HIndex        int     `json:"hindex"`
	AceScore      int     `json:"ace_score"`
}

func SortAffiliationStatistics(stat []AffiliationStatistics, orderby int) {
	switch orderby {
	case 1:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].Count > stat[j].Count
		})
	case 2:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].Citation > stat[j].Citation
		})
	case 3:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].CitationShare > stat[j].CitationShare
		})
	case 4:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].HIndex > stat[j].HIndex
		})
	case 5:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].AceScore > stat[j].AceScore
		})
	case 6:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].T2Citation > stat[j].T2Citation
		})
	}
}

func SortAuthorStatistics(stat []AuthorStatistics, orderby int) {
	switch orderby {
	case 1:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].Count > stat[j].Count
		})
	case 2:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].Citation > stat[j].Citation
		})
	case 3:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].CitationShare > stat[j].CitationShare
		})
	case 4:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].HIndex > stat[j].HIndex
		})
	case 5:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].AceScore > stat[j].AceScore
		})
	case 6:
		sort.Slice(stat, func(i, j int) bool {
			return stat[i].T2Citation > stat[j].T2Citation
		})
	}
}

// Order by. 1: count, 2: citation, 3: citation_share, 4: H-index, 5: AceScore
func Sort(slice interface{}, orderBy int) {
	switch stat := slice.(type) {
	case []AffiliationStatistics:
		SortAffiliationStatistics(stat, orderBy)
	case []AuthorStatistics:
		SortAuthorStatistics(stat, orderBy)
	}
}
