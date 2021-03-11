package api

import "github.com/gin-gonic/gin"

// Commnet
func CreateShortName(c *gin.Context) {
	service := service.Service{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


