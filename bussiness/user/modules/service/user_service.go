package service

import(
	"fmt"
	"time"
	"znda/bussiness/user/model"
	"znda/bussiness/user/utils"
	"github.com/feekk/zddgo/errors"
	"github.com/feekk/zddgo/ztime"
)

type UserService struct{
	userModel model.UserModel
}
func(s UserService) RegisterUser(ctx context.Context, role int, phone, channel string) (uid int64, err error){
	creator := utils.GetRoleIdCreator(role)
	if creator == nil {
		err = errors.Errorf("UIDCreators not found. role:(%d)", role)
		return
	}
	uid := creator.Id()
	now := time.Now()
	model := model.UserModel{
		Uid : uid,
		Role: role,
		Phone: phone,
		RegisterTime: ztime.JsonTime(now),
		CreateTime: ztime.JsonTime(now),
		UpdateTime: ztime.JsonTime(now),
	}
	s.userModel.Create(ctx, &model)
}

