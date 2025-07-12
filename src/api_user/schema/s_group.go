package schema

import (
	"github.com/liuzhaomax/go-maxms/src/api_user/model"
)

type Group struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type GroupReq struct {
	Name string `json:"name"`
}

type GroupRes struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func MapUserGroup2UserGroupEntity(user *model.User, group *model.Group, userGroup *model.UserGroup) {
	userGroup.UserID = user.ID
	userGroup.GroupID = group.ID
}
