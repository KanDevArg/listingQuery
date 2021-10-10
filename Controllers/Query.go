package Controllers

import (
	"github.com/gin-gonic/gin"
	"listingQuery/Models"
	"net/http"
)

func GetAllHits(c *gin.Context) {
	var hits []Models.Hit
	err:= Models.GetAllHits(&hits)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, hits)
	}
}


