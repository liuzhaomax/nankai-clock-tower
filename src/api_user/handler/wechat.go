package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/go-maxms/internal/core"
	"github.com/liuzhaomax/go-maxms/src/api_user/schema"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getWechatOpenid(code string) (*schema.WechatAuthRes, error) {
	cfg := core.GetConfig()
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		cfg.App.Wechat.AppId, cfg.App.Wechat.AppSecret, code)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	authResp := &schema.WechatAuthRes{}
	if err := json.Unmarshal(body, authResp); err != nil {
		return nil, err
	}

	if authResp.ErrCode != 0 {
		return nil, fmt.Errorf("获取微信openid失败: %d: %s", authResp.ErrCode, authResp.ErrMsg)
	}

	return authResp, nil
}

func getAvatarUrl(c *gin.Context, fileHeader *multipart.FileHeader) (string, error) {
	userId := c.Request.Header.Get(core.UserId)
	savePath := getAvatarSavePath(fileHeader.Filename, userId)
	err := c.SaveUploadedFile(fileHeader, savePath)
	if err != nil {
		return core.EmptyString, core.FormatError(core.InternalServerError, "文件保存失败", err)
	}

	cfg := core.GetConfig()
	url := fmt.Sprintf("%s://%s/%s", cfg.Server.Protocol, cfg.App.Domain, savePath)
	return url, nil
}

func getAvatarSavePath(originalFilename string, userId string) string {
	ext := filepath.Ext(originalFilename)
	newFilename := fmt.Sprintf("%s%s", core.ShortUUID(), ext)

	saveDir := filepath.Join(".", "www", "avatars", userId)
	_, err := os.Stat(saveDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(saveDir, os.ModePerm)
		if err != nil {
			return ""
		}
	}

	path := filepath.Join(saveDir, newFilename)
	endpoint := strings.ReplaceAll(path, `\`, `/`)

	return endpoint
}
