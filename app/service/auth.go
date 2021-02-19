package service

import (
	"fmt"
	"gf-demo/app/dao"
	"gf-demo/app/model"
	"gf-demo/app/service/define"
	"gf-demo/library/wechat"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

var Auth = new(authService)

type authService struct {}

//// 手机号登录
//func (s *authService) MobileLogin(a *model.AuthApiMobileLogin) error {
//	if !s.CheckMobile(a.Mobile) {
//		// 注册
//	}
//	if userInfo, err := s.GetUserInfoByMobile(a.Mobile); err != nil {
//
//	}
//
//	return nil
//}

// 登录
func (s *authService) Login(r *ghttp.Request) (string, interface{}) {
	var keys string
	data := r.GetFormMapStrStr()
	stype := data["stype"]
	mobile := data["mobile"]
	switch stype {
	case "mobile": {
		if !s.CheckMobile(mobile) {
			// 注册
			userInfo := new(model.UserInfo)
			userInfo.NickName = ""
			dao.UserInfo.Save(userInfo)
			userLogin := new(model.UserLogin)
			userLogin.Id = userInfo.Id
			userLogin.UserName = mobile
			dao.UserLogin.Save(userLogin)
		}

		if userInfo, err := s.GetUserInfoByMobile(mobile); err == nil {
			keys = gconv.String(userInfo.Id)
			return keys, &model.ContextUser{
				NickName: userInfo.NickName,
				Id: userInfo.Id,
			}
		} else {
			r.Response.WriteJsonExit(gtoken.Error("登录失败"))
		}
	}
	case "wechat": {
		// 微信小程序登录
		jsCode 	:= data["code"]
		info	:= data["info"]
		fmt.Println(info)
		var userInfo = new(define.WechatUserInfo)
		if info != "" {
			if j, err := gjson.DecodeToJson(info); err != nil {
				fmt.Printf("用户信息解析失败: %v", err.Error())
			} else {
				if err := j.Struct(&userInfo); err == nil {
				}
			}
		}
		if session, err := wechat.MiniCode2Session(jsCode); err !=nil {
			r.Response.WriteJsonExit(gtoken.Error("授权失败"))
		} else {
			fmt.Println(session)
			// 注册
			if !s.CheckWechatOpenId(session.OpenID) {
				// 用户基础信息
				userInfo := &model.UserInfo{
					Id: 22,
					NickName: userInfo.NickName,
					AvatarUrl: userInfo.AvatarURL,
					Country: userInfo.Country,
					Province: userInfo.Province,
					City: userInfo.City,
					Gender: gconv.String(userInfo.Gender),
					IdNumber: "123213213",
				}
				dao.UserInfo.Save(userInfo)

				// 保存社交平台账号
				userSocial := &model.UserSocial{
					SocialType: "wechat",
					OpenId: session.OpenID,
					UnionId: session.UnionID,
					SessionKey: session.SessionKey,
					UserId: userInfo.Id,
				}
				dao.UserSocial.Save(userSocial)
			}
			if userInfo, err := s.GetUserInfoByWechatOpenId(session.OpenID); err == nil {
				keys = gconv.String(userInfo.Id)
				go updateSocialSessionKey(userInfo.Id, session.SessionKey)
				return keys, &model.ContextUser{
					NickName: userInfo.NickName,
					Id: userInfo.Id,
				}
			} else {
				r.Response.WriteJsonExit(gtoken.Error("登录失败"))
			}
		}
	}
	}

	return keys, nil
}


// 登录返回方法
func (s *authService) LoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
	} else {
		token := respData.GetString("token")
		r.Response.WriteJson(gtoken.Succ(g.Map{
			"token": token,
		}))
	}
}

//gtoken验证后返回
func (s *authService) AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if r.Method == "OPTIONS" || respData.Success() {
		r.Middleware.Next()
	} else {
		respData.Msg = "用户信息验证失败"
		response := r.Response
		options := response.DefaultCORSOptions()
		response.CORS(options)
		response.WriteJson(respData)
		r.ExitAll()
	}
}


//退出登陆
func (s *authService) LoginOut(r *ghttp.Request) bool {
	//删除在线用户状态
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" && parts[1] != "" {
			//删除在线用户状态操作
			//user_online.Model.Delete("token", parts[1])
		}
	}
	authHeader = r.GetString("token")
	if authHeader != "" {
		//删除在线用户状态操作
		//user_online.Model.Delete("token", authHeader)
	}
	return true
}

// 判断手机号是否已注册
func (s *authService) CheckMobile(mobile string) bool {
	if i, err := dao.UserLogin.FindCount("user_name", mobile); err != nil {
		return false
	} else {
		return i > 0
	}
}

// 根据手机号查询用户信息
func (s *authService) GetUserInfoByMobile(mobile string) (*model.UserInfo, error) {
	if user, err := dao.UserLogin.FindOne("user_name", mobile); err == nil {
		return dao.UserInfo.FindOne("id", user.Id)
	} else {
		return nil, err
	}
}

// 判断微信openid是否已注册
func (s *authService) CheckWechatOpenId(openId string) bool {
	if i, err := dao.UserSocial.FindCount("social_type=? and open_id=?", "wechat", openId); err != nil {
		return false
	} else {
		return i > 0
	}
}

// 根据微信平台openid获取用户信息
func (s *authService) GetUserInfoByWechatOpenId(openId string) (*model.UserInfo, error) {
	if user, err := dao.UserSocial.FindOne("social_type=? and open_id=?", "wechat", openId); err == nil {
		return dao.UserInfo.FindOne("id", user.UserId)
	} else {
		return nil, err
	}
}

// 更新用户社交平台sessionKey
func updateSocialSessionKey(userId int64, sessionKey string) {
	if social, err := dao.UserSocial.FindOne("user_id", userId); err == nil {
		social.SessionKey = sessionKey
		dao.UserSocial.Data("session_key", sessionKey).Where("user_id", userId).Update()
	}
}


