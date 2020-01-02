package model

import(
	"context"
	"github.com/feekk/zddgo"
	"github.com/feekk/zddgo/ztime"
)
type UserModel struct{
	Id int64 `gorm:"column:id" json:"id"`
	Uid int64 `gorm:"column:uid" json:"uid"`
	Role int `gorm:"column:role" json:"role"`
	Phone string `gorm:"column:phone" json:"phone"`
	NickName string `gorm:"column:nick_name" json:"nick_name"`
	Email string `gorm:"column:email" json:"email"`
	RegisterTime ztime.JsonTime `gorm:"column:register_time" json:"register_time"`
	RegisterChannel int `gorm:"column:register_channel" json:"register_channel"`
	WxOpenId string `gorm:"column:wx_open_id" json:"wx_open_id"`
	WxRegisterTime ztime.JsonTime `gorm:"column:wx_register_time" json:"wx_register_time"`
	CreateTime ztime.JsonTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime ztime.JsonTime `gorm:"column:update_time" json:"update_time"`
	DeleteTime ztime.JsonTime `gorm:"column:delete_time" json:"delete_time"`
}
func(m UserModel) TableName() string{
	return "z_user"
}

func(m UserModel) Create(ctx context.Context, user *UserModel) (id int64, err error){
	db := zddgo.MysqlDef().Create(user)
	if err = db.Error; err != nil {
		return 
	}
	id = user.Id
	return
}

func(m UserModel) FindByUid(ctx context.Context, uid int64) (model *UserModel, err error){
	db := zddgo.MysqlDef().Where("uid = ?", uid).First(&model)
	if db.RecordNotFound(){
		return nil, nil
	}
	err = ret.Error
	return 
}