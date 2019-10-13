package aceranking

import (
	"acemap/controller"
	"acemap/data"
	"acemap/model/aceranking"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authorForm struct {
	Type        string  `form:"type"`
	VenueIDs    string  `form:"venue_ids"`
	Area        string  `form:"area"`
	StartYear   int     `form:"start_year"`
	EndYear     int     `form:"end_year"`
	FirstAuthor int     `form:"first_author"`
	AuthorID    data.ID `form:"author_id"`
}

func Author(c *gin.Context) {
	var param authorForm
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
	authorIndex := data.AuthorIDToIndex[param.AuthorID]
	var authorInfo aceranking.AuthorInfo
	if param.FirstAuthor == 0 {
		authorInfo = aceranking.GetAuthorInfoOnAllAuthors(param.Type, param.Area, venueMask, param.StartYear, param.EndYear, authorIndex)
	} else if param.FirstAuthor == 1 {
		authorInfo = aceranking.GetAuthorInfoOnFirstAuthorWeak(param.Type, param.Area, venueMask, param.StartYear, param.EndYear, authorIndex)
	} else {
		authorInfo = aceranking.GetAuthorInfoOnFirstAuthorStrong(param.Type, param.Area, venueMask, param.StartYear, param.EndYear, authorIndex)
	}

	c.JSON(http.StatusOK, controller.RespondData(authorInfo))
}
