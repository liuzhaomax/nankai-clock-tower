package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId           string `gorm:"index:idx_user_id;unique;varchar(50);not null"`
	WechatOpenid     string `gorm:"index:idx_wechat_openid;unique;varchar(50);not null"`
	WechatUnionid    string `gorm:"varchar(50);not null"`
	WechatSessionKey string `gorm:"varchar(200);not null"`
	WechatAvatar     string `gorm:"varchar(200)"`
	WechatNickname   string `gorm:"varchar(50)"`
}
