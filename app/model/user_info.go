// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"gf-demo/app/model/internal"
)

// UserInfo is the golang structure for table user_info.
type UserInfo internal.UserInfo

// Fill with you ideas below.

// 登录成功信息
type LoginResult struct {
	NickName string
	Avatar string
	Token string
}
