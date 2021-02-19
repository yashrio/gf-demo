package define

// 用户分页查询条件
type UserQuery struct {
	NickName	string	// 昵称
	Page		int		`v:"required#请输入页码" binding:"required"` 	// 页码
	Size		int		`v:"required#请输入每页条数" binding:"required"`	// 每页条数
}

type UserQueryResult struct {
	*PageListRes
}

// 微信用户信息
type WechatUserInfo struct {
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
}

// 用户信息
type ProfileResult struct {
	NickName string // 昵称
	Avatar   string // 头像
}