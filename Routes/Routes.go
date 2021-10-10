package Routes

import (
	"github.com/gin-gonic/gin"
	"listingQuery/Controllers"
)

func SetupRoutes() *gin.Engine{
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("all", Controllers.GetAllHits)
	}
	return r
}