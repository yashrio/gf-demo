package model

const (
	ContextKey = "ContextKey"
)

type Context struct {
	User	*ContextUser
}

// 上下文用户新
type ContextUser struct {
	Id 			int64
	NickName	string
}
