package api

import (
	"github.com/gin-gonic/gin"
	"singo/service/shortname"
)

// Commnet
func CreateShortName(c *gin.Context) {
	service := shortname.CreateShortNameService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShortNameShow 短域名Show
func ShortNameShow(c *gin.Context) {
	service := shortname.ShowShortNameService{}

	//if err := c.ShouldBind(&service); err == nil {
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
	//} else {
	//c.JSON(200, ErrorResponse(err))
	//}
}

// ListShortName 短域名 列表
func ListShortName(c *gin.Context) {
	service := shortname.ListShortNameService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteShortName 短域名删除
func DeleteShortName(c *gin.Context) {
	service := shortname.DeleteShortNameService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)

}

// UpdateShortName 短域名更新
// Commnet
func UpdateShortName(c *gin.Context) {
	service := shortname.UpdateShortNameService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
