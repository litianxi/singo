package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// ShortName 用户模型
type ShortName struct {
	gorm.Model
	ShortName  string `gorm:"size:100"`
	Extra      string `gorm:"size:1000"`
	LongName   string `gorm:"size:500"`
	ExpireTime time.Time
	Class      string  `gorm:"size:1000"`
	Logs       string  `gorm:"size:1000"`
	Managers   []*User `gorm:"many2many:shortname_manager"`
	Leaders    []*User `gorm:"many2many:shortname_leaders"`
}

// Users 用户模型
type Users struct {
	gorm.Model
	Owner    string
	OwnerId  string
	Dept     string
	Leader   string
	LeaderId string
}
