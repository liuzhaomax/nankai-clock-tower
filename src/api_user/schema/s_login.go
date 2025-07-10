package schema

import (
	"github.com/liuzhaomax/go-maxms/internal/core"
	"github.com/liuzhaomax/go-maxms/src/api_user/model"
)

type LoginReq struct {
	Code string `json:"code"`
}

type WechatAuthRes struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid,omitempty"`
	ErrCode    int    `json:"errcode,omitempty"`
	ErrMsg     string `json:"errmsg,omitempty"`
}

type TokenRes struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}

func MapWechatAuthRes2UserEntity(wechatAuthRes *WechatAuthRes, user *model.User) {
	user.WechatOpenid = wechatAuthRes.Openid
	user.WechatUnionid = wechatAuthRes.Unionid
	user.WechatSessionKey = wechatAuthRes.SessionKey
	if user.UserId == "" {
		user.UserId = core.ShortUUID()
	}
	if user.WechatNickname == "" {
		user.WechatNickname = "铁狼" + user.UserId[:6]
	}
}
