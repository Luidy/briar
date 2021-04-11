package service

import (
	"context"
	"ges/model"
)

type UserService interface {
	RegistUser(ctx context.Context, req *model.UserInfo) (res *model.UserInfo, err error)
}
