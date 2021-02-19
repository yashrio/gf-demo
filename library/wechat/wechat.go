package wechat

import (
	"errors"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
)

var wc *wechat.Wechat
var miniCfg *config.Config

// 小程序登录授权结果
type Code2Session struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}


// 初始化微信配置
func InitWechat() error {
	wc = wechat.NewWechat()
	redisConfig, err := gredis.ConfigFromStr(g.Cfg().GetString(("redis.default")))
	if err != nil {
		return errors.New("初始化微信配置失败")
	}
	redisOpts := &cache.RedisOpts{
		Host:        redisConfig.Host,
		Password:    redisConfig.Pass,
		Database:    8,
		MaxActive:   redisConfig.MaxActive,
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: gconv.Int(redisConfig.IdleTimeout.Seconds()),
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)

	// 小程序配置信息
	miniCfg = &config.Config {
		AppID: g.Cfg().GetString("wechat.mini.AppID"),
		AppSecret: g.Cfg().GetString("wechat.mini.AppSecret"),
	}

	return nil
}

// 小程序登录
func MiniCode2Session(jsCode string) (Code2Session, error) {
	mini := wc.GetMiniProgram(miniCfg)
	a := mini.GetAuth()
	if result, err := a.Code2Session(jsCode); err != nil {
		return Code2Session{}, err
	} else {
		return Code2Session{
			OpenID: result.OpenID,
			UnionID: result.UnionID,
			SessionKey: result.SessionKey,
		}, nil
	}
}

// 小程序解密用户信息
func MiniGetUserInfo(sessionKey, encryptedData, iv string) (*encryptor.PlainData, error) {
	mini := wc.GetMiniProgram(miniCfg)
	a := mini.GetEncryptor()
	if result, err := a.Decrypt(sessionKey, encryptedData, iv); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

