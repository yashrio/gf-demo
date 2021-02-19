package service

import (
	"context"
	"errors"
	"gf-demo/app/dao"
	"gf-demo/app/model"
	"gf-demo/app/service/define"
	"gf-demo/library/guid"
	"github.com/gogf/gf/util/gconv"
)

var User = new(userService)

type userService struct {}

// 用户注册
func (s *userService) SignUp(r *model.UserServiceSignUpReg) error {
	if r.Nickname == "" {
		r.Nickname = r.Passport
	}

	// 保存
	r.Id = guid.Next()
	if _, err := dao.User.Save(r); err != nil {
		return err
	}
	return nil
}

// 查询用户列表
func (s *userService) GetUserList(q *define.UserQuery) (*define.PageListRes, error)  {
	m := dao.UserInfo.M
	if q.NickName != "" {
		m = m.Where("nick_name like ?", q.NickName)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	list, err := m.Page(q.Page, q.Size).All()
	if err != nil {
		return nil, errors.New("查询数据记录失败")
	}

	return &define.PageListRes{
		Total: total,
		List: list,
	}, nil
}

func (s *userService) Profile(ctx context.Context) (*define.ProfileResult,error) {
	if c := Context.Get(ctx); c != nil && c.User != nil {
		if user, err := dao.UserInfo.FindOne("id", c.User.Id); err != nil {

		} else {
			profile := new(define.ProfileResult)
			gconv.Struct(user, profile)
			return profile, nil
		}
		return nil, nil
	} else {
		return nil, errors.New("未登录")
	}
}
