package aceranking

import (
	"acemap/controller"
	"acemap/data"
	"acemap/model/aceranking"
	"acemap/model/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

type authorListForm struct {
	VenueIDs    string `form:"venue_ids"`
	Area        string `form:"area"`
	StartYear   int    `form:"start_year"`
	EndYear     int    `form:"end_year"`
	FirstAuthor int    `form:"first_author"`
	OrderBy     int    `form:"order_by"`
}

func AuthorList(c *gin.Context) {
	var param authorListForm
	var err error
	err = c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusOK, controller.RespondError(err.Error()))
		return
	}

	venueMask, err := aceranking.GetVenueMask(param.VenueIDs)
	if err != nil {
		c.JSON(http.StatusOK, controller.RespondError(err.Error()))
		return
	}
	//t := time.Now()
	var statistics []aceranking.AuthorStatistics
	if param.FirstAuthor == 0 {
		statistics = aceranking.GetAuthorStatOnAllAuthors(venueMask, param.Area, param.StartYear, param.EndYear)
	} else if param.FirstAuthor == 1 {
		statistics = aceranking.GetAuthorStatOnFirstAuthorWeak(venueMask, param.Area, param.StartYear, param.EndYear)
	} else {
		statistics = aceranking.GetAuthorStatOnFirstAuthorStrong(venueMask, param.Area, param.StartYear, param.EndYear)
	}

	//fmt.Println(len(statistics))
	//aceranking.Sort(statistics, param.OrderBy)

	//end := 50
	//if len(statistics) < end {
	//	end = len(statistics)
	//}

	// Number of rows to show AT MOST.
	end := 50
	if len(statistics) < end {
		end = len(statistics)
	}
	// Order by. 1: count; 2: citation, 3: citation_share, 4: H-index, 5: AceScore
	switch param.OrderBy {
	case 1:
		sort.Slice(statistics, func(i, j int) bool {
			return statistics[i].Count > statistics[j].Count
		})
		for end > 0 && statistics[end-1].Count == 0 {
			end--
		}
	case 2:
		sort.Slice(statistics, func(i, j int) bool {
			return statistics[i].Citation > statistics[j].Citation
		})
		for end > 0 && statistics[end-1].Citation == 0 {
			end--
		}
	case 3:
		sort.Slice(statistics, func(i, j int) bool {
			return statistics[i].CitationShare > statistics[j].CitationShare
		})
		for end > 0 && statistics[end-1].CitationShare == 0 {
			end--
		}
	case 4:
		sort.Slice(statistics, func(i, j int) bool {
			return statistics[i].HIndex > statistics[j].HIndex
		})
		for end > 0 && statistics[end-1].HIndex == 0 {
			end--
		}
	case 5:
		sort.Slice(statistics, func(i, j int) bool {
			return statistics[i].AceScore > statistics[j].AceScore
		})
		for end > 0 && statistics[end-1].AceScore == 0 {
			end--
		}
	}

	for u := 0; u < end; u++ {
		statistics[u].AuthorName = dao.GetAuthorNameByID(statistics[u].AuthorID)
		statistics[u].AffList = make([]aceranking.AuthorAffInfo, 0)
		for affIndex, count := range statistics[u].AffCount {
			affID := data.AffiliationList[affIndex]
			statistics[u].AffList = append(statistics[u].AffList, aceranking.AuthorAffInfo{
				AffiliationID:   affID,
				Abbreviation:    dao.GetAffiliationAbbrByID(affID),
				AffiliationName: dao.GetAffiliationNameByID(affID),
				Count:           count,
			})
		}
		sort.Slice(statistics[u].AffList, func(i, j int) bool {
			return statistics[u].AffList[i].Count > statistics[u].AffList[j].Count
		})
		l := len(statistics[u].AffList)
		if l == 0 {
			continue
		}
		end := 1
		if l >= 2 && statistics[u].AffList[1].Count >= 10 {
			end = 2
			if l >= 3 && statistics[u].AffList[2].Count >= 10 {
				end = 3
			}
		}
		statistics[u].AffList = statistics[u].AffList[:end]
	}

	c.JSON(http.StatusOK, controller.RespondData(statistics[:end]))
}
