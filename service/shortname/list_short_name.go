package shortname

import (
	"singo/model"
	"singo/serializer"
)

// ListShortNameService 列出短域名
type ListShortNameService struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

// List 短域名
func (service *ListShortNameService) List() serializer.Response {

	var shortnames []model.ShortName

	page := service.Page
	page_size := service.PageSize
	total := 0
	//
	//if service.Limit == 0 {
	//	service.Limit = 3
	//}

	//if err := model.DB.Model(model.ShortName{}).Count(&total).Error; err != nil {
	//	return serializer.Response{
	//		Code:  50000,
	//		Msg:   "数据库连接错误",
	//		Error: err.Error(),
	//	}
	//}
	//
	//if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&shortnames).Error; err != nil {
	//	return serializer.Response{
	//		Code:  50000,
	//		Msg:   "数据库连接错误",
	//		Error: err.Error(),
	//	}
	//}

	if err := model.DB.Model(model.ShortName{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	if err := model.DB.Scopes(serializer.Paginate(page, page_size)).Find(&shortnames).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildShortNames(shortnames), uint(total))
}
