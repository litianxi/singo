package shortname

import (
	"singo/model"
	"singo/serializer"
)

// ShowShortNameService 列出短域名
type ShowShortNameService struct {
	//ShortName  string    `form:"short_name" json:"short_name" binding:"required,min=5,max=100"`
	//Class      string    `form:"class" json:"class" binding:"required"`
	//LongName   string    `form:"long_name" json:"long_name" binding:"required"`
	//ExpireTime time.Time `form:"expire_time" json:"expire_time" binding:"required"`
}

// Create 短域名创建函数
func (service *ShowShortNameService) Show(id string) serializer.Response {

	var shortname model.ShortName

	//shortname := model.ShortName{
	//ShortName:  service.ShortName,
	//Class:      service.Class,
	//LongName:   service.LongName,
	//ExpireTime: service.ExpireTime,
	//}

	if err := model.DB.First(&shortname, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Error: err.Error(),
			Msg:   "短域名不存在",
		}

		//return serializer.ParamErr("查询失败", err)
	}

	return serializer.BuildShortNameResponse(shortname)
}
