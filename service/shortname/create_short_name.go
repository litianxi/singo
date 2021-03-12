package shortname

import (
	"singo/model"
	"singo/serializer"
	"time"
)

// CreateShortNameService 短域名创建服务
type CreateShortNameService struct {
	ShortName  string    `form:"short_name" json:"short_name" binding:"required,min=5,max=100"`
	Class      string    `form:"class" json:"class" binding:"required"`
	LongName   string    `form:"long_name" json:"long_name" binding:"required"`
	ExpireTime time.Time `form:"expire_time" json:"expire_time" binding:"required"`
}

// Create 短域名创建函数
func (service *CreateShortNameService) Create() serializer.Response {
	shortname := model.ShortName{
		ShortName:  service.ShortName,
		Class:      service.Class,
		LongName:   service.LongName,
		ExpireTime: service.ExpireTime,
	}

	if err := model.DB.Create(&shortname).Error; err != nil {
		return serializer.ParamErr("创建失败", err)
	}

	return serializer.BuildShortNameResponse(shortname)
}
