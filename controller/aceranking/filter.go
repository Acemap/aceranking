package aceranking

import (
	"acemap/controller"
	"acemap/model/aceranking"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterHandler(c *gin.Context) {
	f := aceranking.GetFilter()
	c.JSON(http.StatusOK, controller.RespondData(f))
}
