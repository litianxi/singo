package serializer

import (
	"singo/model"
	"time"
)

// User 用户序列化器
type ShortName struct {
	ID         uint      `json:"id"`
	ShortName  string    `json:"short_name"`
	LongName   string    `json:"long_name"`
	ExpireTime time.Time `json:"expire_time"`
	Class      string    `json:"class"`
	//Logs []*Log
	Owner       string `json:"owner"`
	OwnerEmail  string `json:"owner_email"`
	Dept        string `json:"dept"`
	Leader      string `json:"leader"`
	LeaderEmail string `json:"leader_email"`
}

// BuildShortName 序列化用户
func BuildShortName(shortname model.ShortName) ShortName {
	return ShortName{
		ID:         shortname.ID,
		ShortName:  shortname.ShortName,
		LongName:   shortname.LongName,
		ExpireTime: shortname.ExpireTime,
		Class:      shortname.Class,
		//Logs []*Log
		Owner:       shortname.Owner,
		OwnerEmail:  shortname.OwnerEmail,
		Dept:        shortname.Dept,
		Leader:      shortname.Leader,
		LeaderEmail: shortname.LeaderEmail,
	}
}

// BuildShortNameResponse 序列化用户响应
func BuildShortNameResponse(shortname model.ShortName) Response {
	return Response{
		Data: BuildShortName(shortname),
	}
}

func BuildShortNames(items []model.ShortName) (shortnames []ShortName) {
	for _, item := range items {
		shortname := BuildShortName(item)
		shortnames = append(shortnames, shortname)
	}
	return
}
