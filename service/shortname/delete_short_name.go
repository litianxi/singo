package shortname

import (
	"singo/model"
	"singo/serializer"
)

// DeleteShortNameService 删除短域名
type DeleteShortNameService struct {
}

// Delete 短域名删除函数
func (service *DeleteShortNameService) Delete(id string) serializer.Response {

	var shortname model.ShortName

	if err := model.DB.First(&shortname, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Error: err.Error(),
			Msg:   "短域名不存在",
		}
	}

	if err := model.DB.Delete(&shortname).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Error: err.Error(),
			Msg:   "短域名删除失败",
		}
	}

	return serializer.Response{}
}
