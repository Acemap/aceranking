package aceranking

import (
	"acemap/controller"
	"acemap/data"
	"acemap/model/aceranking"
	"github.com/gin-gonic/gin"
	"net/http"
)

type affiliationAuthorForm struct {
	Type          string  `form:"type"`
	VenueIDs      string  `form:"venue_ids"`
	StartYear     int     `form:"start_year"`
	EndYear       int     `form:"end_year"`
	FirstAuthor   int     `form:"first_author"`
	AffiliationID data.ID `form:"affiliation_id"`
	AuthorID      data.ID `form:"author_id"`
}

func AffiliationAuthor(c *gin.Context) {
	var param affiliationAuthorForm
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

	affIndex := data.AffiliationIDToIndex[param.AffiliationID]
	authorIndex := data.AuthorIDToIndex[param.AuthorID]
	var affAuthorInfo aceranking.AuthorInfo
	if param.FirstAuthor == 0 {
		affAuthorInfo = aceranking.GetAffAuthorInfoOnAllAuthors(param.Type, venueMask, param.StartYear, param.EndYear, affIndex, authorIndex)
	} else if param.FirstAuthor == 1 {
		affAuthorInfo = aceranking.GetAffAuthorInfoOnFirstAuthorWeak(param.Type, venueMask, param.StartYear, param.EndYear, affIndex, authorIndex)
	} else {
		affAuthorInfo = aceranking.GetAffAuthorInfoOnFirstAuthorStrong(param.Type, venueMask, param.StartYear, param.EndYear, affIndex, authorIndex)
	}

	c.JSON(http.StatusOK, controller.RespondData(affAuthorInfo))
}
