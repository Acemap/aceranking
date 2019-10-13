package aceranking

import (
	"acemap/controller"
	"acemap/model/aceranking"
	"github.com/gin-gonic/gin"
	"net/http"
)

type venueForm struct {
	Type      string `form:"type"`
	Area      string `form:"area"`
	StartYear int    `form:"start_year"`
	EndYear   int    `form:"end_year"`
}

func VenueHandler(c *gin.Context) {
	var param venueForm
	var err error

	err = c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusOK, controller.RespondError(err.Error()))
		return
	}

	vc, err := aceranking.GetVenueCategory(param.Type, param.Area, param.StartYear, param.EndYear)
	if err != nil {
		c.JSON(http.StatusOK, controller.RespondError("Error: "+err.Error()))
	}

	c.JSON(http.StatusOK, controller.RespondData(vc))
}
