package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ModelUserSet = wire.NewSet(wire.Struct(new(ModelUser), "*"))

type ModelUser struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	UserId           string      `gorm:"index:idx_user_id;unique;size:50;not null"`
	WechatOpenid     string      `gorm:"index:idx_wechat_openid;unique;size:50;not null"`
	WechatUnionid    string      `gorm:"size:50;not null"`
	WechatSessionKey string      `gorm:"size:200;not null"`
	WechatAvatar     string      `gorm:"size:200"`
	WechatNickname   string      `gorm:"size:50"`
	Groups           []Group     `gorm:"many2many:user_group;"`
	UserGroups       []UserGroup `gorm:"foreignKey:UserID"`
}

type Group struct {
	gorm.Model
	Name  string `gorm:"index:idx_name;unique;size:50;not null"`
	Users []User `gorm:"many2many:user_group;"`
}

type UserGroup struct {
	UserID  uint `gorm:"primaryKey"`
	GroupID uint `gorm:"primaryKey"`
	Score   int  `gorm:"default:0"`
}
