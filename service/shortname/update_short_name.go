package shortname

import (
	"singo/model"
	"singo/serializer"
	"time"
)

// UpdateShortNameService 更新短域名
type UpdateShortNameService struct {
	ShortName  string    `form:"short_name" json:"short_name" `
	Class      string    `form:"class" json:"class" `
	LongName   string    `form:"long_name" json:"long_name" `
	ExpireTime time.Time `form:"expire_time" json:"expire_time"`
}

// UpdateShortNameService 更新函数
func (service *UpdateShortNameService) Update(id string) serializer.Response {

	var shortname model.ShortName

	if err := model.DB.First(&shortname, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Error: err.Error(),
			Msg:   "短域名不存在",
		}
	}

	if err := model.DB.Model(&shortname).Updates(service).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Error: err.Error(),
			Msg:   "短域名更新失败",
		}
	}

	return serializer.BuildShortNameResponse(shortname)
}
