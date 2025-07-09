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
	"time"
)

func (h *HandlerUser) GetPuk(c *gin.Context) (string, error) {
	return core.GetConfig().App.PublicKeyStr, nil
}

func (h *HandlerUser) PostLogin(c *gin.Context) (*schema.TokenRes, error) {
	res := &schema.TokenRes{}
	loginReq := &schema.LoginReq{}
	err := c.ShouldBind(loginReq)
	if err != nil {
		h.Logger.Error(core.FormatError(core.ParseIssue, "请求体无效", err))
		return nil, err
	}
	decryptedCode, err := core.RSADecrypt(core.GetPrivateKey(), loginReq.Code)
	if err != nil {
		h.Logger.Error(core.FormatError(core.PermissionDenied, "请求体解码异常", err))
		return nil, err
	}
	wechatAuthRes, err := getWechatOpenid(decryptedCode)
	if err != nil {
		h.Logger.Error(core.FormatError(core.DownstreamDown, "请求微信code2session接口失败", err))
		return nil, err
	}
	// 更新数据库
	user := &model.User{}
	err = h.Tx.ExecTrans(c, func(ctx context.Context) error {
		err = h.Model.QueryUserByWechatOpenid(ctx, wechatAuthRes.Openid, user)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			h.Logger.Error(core.FormatError(core.DBDenied, "根据openid查找用户失败", err))
			return err
		}
		schema.MapWechatAuthRes2UserEntity(wechatAuthRes, user)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = h.Model.CreateUser(ctx, user)
			if err != nil {
				h.Logger.Error(core.FormatError(core.DBDenied, "创建用户失败", err))
				return err
			}
		} else {
			err = h.Model.UpdateUser(ctx, user)
			if err != nil {
				h.Logger.Error(core.FormatError(core.DBDenied, "更新用户失败", err))
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, code.DBOperationFailed
	}
	// 定义过期时长
	maxAge := 60 * 60 * 24 * 7 // 一周
	duration := time.Second * time.Duration(maxAge)
	// 生成Bearer jwt，使用userID签发
	j := core.NewJWT()
	token, err := j.GenerateToken(user.UserId, duration)
	if err != nil {
		h.Logger.Error(core.FormatError(core.PermissionDenied, "Token生成失败", err))
		return res, err
	}
	bearerToken := core.Bearer + token
	// 对Bearer jwt 进行RSA加密
	encryptedBearerToken, err := core.RSAEncrypt(core.GetPublicKey(), bearerToken)
	if err != nil {
		h.Logger.Error(core.FormatError(core.PermissionDenied, "Token加密失败", err))
		return res, err
	}
	// 拼接响应
	res.Token = encryptedBearerToken
	res.UserId = user.UserId
	return res, nil
}

func (h *HandlerUser) DeleteLogin(c *gin.Context) (any, error) {
	// maxAge := int(time.Millisecond)
	// domain := core.GetConfig().App.Domain
	// c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie(
	//     core.UserID,
	//     core.EmptyString,
	//     maxAge,
	//     "/",
	//     domain,
	//     true,
	//     true)
	return nil, nil
}
