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

func (h *HandlerUser) PostAvatar(c *gin.Context) (any, error) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		h.Logger.Error(core.FormatError(core.ParseIssue, "读取头像失败", err))
		return nil, err
	}

	avatarUrl, err := getAvatarUrl(c, fileHeader)
	if err != nil {
		h.Logger.Error(core.FormatError(core.IOException, "保存头像失败", err))
		return nil, err
	}

	err = h.Tx.ExecTrans(c, func(ctx context.Context) error {
		useId := c.Request.Header.Get(core.UserId)
		user := &model.User{
			UserId:       useId,
			WechatAvatar: avatarUrl,
		}
		err = h.Model.UpdateUserAvatar(ctx, user)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId存入头像失败", err))
			return err
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId存入头像失败", err))
			return err
		}
		return nil
	})
	if err != nil {
		return nil, code.DBOperationFailed
	}

	return nil, nil
}

func (h *HandlerUser) PatchNickName(c *gin.Context) (any, error) {
	nickNameReq := &schema.NickNameReq{}
	err := c.ShouldBind(nickNameReq)
	if err != nil {
		h.Logger.Error(core.FormatError(core.ParseIssue, "请求体无效", err))
		return nil, err
	}

	err = h.Tx.ExecTrans(c, func(ctx context.Context) error {
		useId := c.Request.Header.Get(core.UserId)
		user := &model.User{
			UserId:         useId,
			WechatNickname: nickNameReq.NickName,
		}
		err = h.Model.UpdateUserNickName(ctx, user)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId存入昵称失败", err))
			return err
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId存入昵称失败", err))
			return err
		}
		return nil
	})
	if err != nil {
		return nil, code.DBOperationFailed
	}

	return nil, nil
}

func (h *HandlerUser) GetUser(c *gin.Context) (*schema.UserRes, error) {
	useId := c.Request.Header.Get(core.UserId)
	user := &model.User{
		UserId: useId,
	}
	err := h.Tx.ExecTrans(c, func(ctx context.Context) error {
		err := h.Model.QueryUserByUserId(ctx, user)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId获取用户信息失败", err))
			return err
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据userId获取用户信息失败", err))
			return err
		}
		return nil
	})
	if err != nil {
		return nil, code.DBOperationFailed
	}
	// mapping
	userRes := &schema.UserRes{}
	schema.MapUserEntity2UserRes(user, userRes)
	return userRes, nil
}
