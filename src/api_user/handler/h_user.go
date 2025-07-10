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
	"io"
	"mime/multipart"
)

func (h *HandlerUser) PatchAvatar(c *gin.Context) (any, error) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		h.Logger.Error(core.FormatError(core.ParseIssue, "读取头像失败", err))
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		h.Logger.Error(core.FormatError(core.IOException, "头像文件打开失败", err))
		return nil, err
	}
	defer func(file *multipart.File) {
		err := (*file).Close()
		if err != nil {
			h.Logger.Error(core.FormatError(core.IOException, "头像文件关闭失败", err))
		}
	}(&file)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		h.Logger.Error(core.FormatError(core.ParseIssue, "头像二进制流转[]byte失败", err))
		return nil, err
	}

	err = h.Tx.ExecTrans(c, func(ctx context.Context) error {
		useId := c.Request.Header.Get(core.UserId)
		user := &model.User{
			UserId:       useId,
			WechatAvatar: fileBytes,
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
