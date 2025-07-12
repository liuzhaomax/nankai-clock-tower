package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/go-maxms/internal/core"
	"github.com/liuzhaomax/go-maxms/src/api_user/code"
	"github.com/liuzhaomax/go-maxms/src/api_user/model"
	"github.com/liuzhaomax/go-maxms/src/api_user/schema"
	"gorm.io/gorm"
)

func (h *HandlerUser) PostGroup(c *gin.Context) (any, error) {
	groupReq := &schema.GroupReq{}
	err := c.ShouldBind(groupReq)
	if err != nil {
		h.Logger.Error(core.FormatError(core.ParseIssue, "请求体无效", err))
		return nil, code.RequestParsingErr
	}

	group := &model.Group{
		Name: groupReq.Name,
	}
	userId := c.Request.Header.Get(core.UserId)
	user := &model.User{
		UserId: userId,
	}
	userGroup := &model.UserGroup{}
	err = h.Tx.ExecTrans(c, func(ctx context.Context) error {
		// 创建group
		err = h.Model.CreateGroup(ctx, group)
		if err != nil && !errors.Is(err, gorm.ErrDuplicatedKey) {
			h.Logger.Error(core.FormatError(core.DBDenied, "创建群组失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			h.Logger.Error(core.FormatError(core.DBDenied, "创建群组失败", err))
			return code.GroupExisted
		}
		// 根据name找到group的ID
		err = h.Model.QueryGroupByName(ctx, group)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据name获取群组信息失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据name获取群组信息失败", err))
			return code.GroupNotExisted
		}
		// 根据userId找到user的ID
		err = h.Model.QueryUserByUserId(ctx, user)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId获取用户信息失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId获取用户信息失败", err))
			return code.UserNotExisted
		}
		// 拼接userGroup
		schema.MapUserGroup2UserGroupEntity(user, group, userGroup)
		// 创建关联信息userGroup
		err = h.Model.CreateUserGroup(ctx, userGroup)
		if err != nil && !errors.Is(err, gorm.ErrDuplicatedKey) {
			h.Logger.Error(core.FormatError(core.DBDenied, "关联用户与群组失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			h.Logger.Error(core.FormatError(core.DBDenied, "关联用户与群组失败", err))
			return code.UserGroupExisted
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	res := &schema.GroupRes{
		Id:    group.ID,
		Name:  group.Name,
		Score: 0,
	}
	return res, nil
}

func (h *HandlerUser) PatchJoinGroup(c *gin.Context) (any, error) {
	joinGroupReq := &schema.GroupReq{}
	err := c.ShouldBind(joinGroupReq)
	if err != nil {
		h.Logger.Error(core.FormatError(core.ParseIssue, "请求体无效", err))
		return nil, code.RequestParsingErr
	}

	group := &model.Group{
		Name: joinGroupReq.Name,
	}
	userId := c.Request.Header.Get(core.UserId)
	user := &model.User{
		UserId: userId,
	}
	userGroup := &model.UserGroup{}
	err = h.Tx.ExecTrans(c, func(ctx context.Context) error {
		// 根据name找到group的ID
		err = h.Model.QueryGroupByName(ctx, group)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据name获取群组信息失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据name获取群组信息失败", err))
			return code.GroupNotExisted
		}
		// 根据userId找到user的ID
		err = h.Model.QueryUserByUserId(ctx, user)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId获取用户信息失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId获取用户信息失败", err))
			return code.UserNotExisted
		}
		// 拼接userGroup
		schema.MapUserGroup2UserGroupEntity(user, group, userGroup)
		// 创建关联信息userGroup
		err = h.Model.CreateUserGroup(ctx, userGroup)
		if err != nil && !errors.Is(err, gorm.ErrDuplicatedKey) {
			h.Logger.Error(core.FormatError(core.DBDenied, "关联用户与群组失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			h.Logger.Error(core.FormatError(core.DBDenied, "关联用户与群组失败", err))
			return code.UserGroupExisted
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	res := &schema.GroupRes{
		Id:    group.ID,
		Name:  group.Name,
		Score: 0,
	}
	return res, nil
}

func (h *HandlerUser) DeleteQuitGroup(c *gin.Context) (any, error) {
	quitGroupReq := &schema.GroupReq{}
	err := c.ShouldBind(quitGroupReq)
	if err != nil {
		h.Logger.Error(core.FormatError(core.ParseIssue, "请求体无效", err))
		return nil, code.RequestParsingErr
	}

	group := &model.Group{
		Name: quitGroupReq.Name,
	}
	userId := c.Request.Header.Get(core.UserId)
	user := &model.User{
		UserId: userId,
	}
	userGroup := &model.UserGroup{}
	err = h.Tx.ExecTrans(c, func(ctx context.Context) error {
		// 根据name找到group的ID
		err = h.Model.QueryGroupWithUsersByName(ctx, group)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据name获取群组信息失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据name获取群组信息失败", err))
			return code.GroupNotExisted
		}
		// 根据userId找到user的ID
		err = h.Model.QueryUserByUserId(ctx, user)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId获取用户信息失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId获取用户信息失败", err))
			return code.UserNotExisted
		}
		// 拼接userGroup
		schema.MapUserGroup2UserGroupEntity(user, group, userGroup)
		// 创建关联信息userGroup
		err = h.Model.DeleteUserGroup(ctx, userGroup)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "退组失败", err))
			return code.DBFailed
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "退组失败", err))
			return code.QuitGroupFailed
		}
		// 如果该组没有用户，则硬删除该组：
		// 删除关联成功才能到这步，如果删除了关联，说明退组成功，那么如果原先组里只有1人，说明现在组里没人
		if len(group.Users) == 1 {
			err = h.Model.DeleteGroup(ctx, group)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				h.Logger.Error(core.FormatError(core.DBDenied, "删除群组失败", err))
				return code.DBFailed
			}
			if errors.Is(err, gorm.ErrRecordNotFound) {
				h.Logger.Error(core.FormatError(core.DBDenied, "删除群组失败", err))
				return code.DeleteGroupFailed
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
