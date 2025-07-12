package schema

import (
	"github.com/liuzhaomax/go-maxms/src/api_user/model"
)

type NickNameReq struct {
	NickName string `json:"nickName"`
}

type NickNameRes struct {
	NickName string `json:"nickName"`
}

type AvatarRes struct {
	Avatar string `json:"avatar"`
}

type UserRes struct {
	UserID   string   `json:"userId"`
	Avatar   string   `json:"avatar"`
	NickName string   `json:"nickName"`
	Groups   *[]Group `json:"groups"`
}

func MapUserEntity2UserRes(user *model.User, userRes *UserRes) {
	userRes.UserID = user.UserId
	userRes.Avatar = user.WechatAvatar
	userRes.NickName = user.WechatNickname
	userRes.Groups = &[]Group{}

	for _, group := range user.Groups {
		tmpGroup := Group{
			Id:   group.ID,
			Name: group.Name,
		}
		for _, userGroup := range user.UserGroups {
			if group.ID == userGroup.GroupID {
				tmpGroup.Score = userGroup.Score
				break
			}
		}
		*userRes.Groups = append(*userRes.Groups, tmpGroup)
	}
}
