package handler

import (
	"encoding/json"
	"fmt"
	"github.com/liuzhaomax/go-maxms/internal/core"
	"github.com/liuzhaomax/go-maxms/src/api_user/schema"
	"io"
	"net/http"
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
		return nil, fmt.Errorf("获取微信openid失败: %s: %s", authResp.ErrCode, authResp.ErrMsg)
	}

	return authResp, nil
}
