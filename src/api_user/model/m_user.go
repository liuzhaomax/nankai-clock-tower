package model

import (
	"context"
	"github.com/google/wire"
	"github.com/liuzhaomax/go-maxms/internal/core"
	"gorm.io/gorm"
)

var ModelUserSet = wire.NewSet(wire.Struct(new(ModelUser), "*"))

type ModelUser struct {
	DB *gorm.DB
}

func (m *ModelUser) QueryUserByWechatOpenid(ctx context.Context, openid string, user *User) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Where("wechat_openid = ?", openid).First(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *ModelUser) CreateUser(ctx context.Context, user *User) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *ModelUser) UpdateUser(ctx context.Context, user *User) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Save(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *ModelUser) UpdateUserAvatar(ctx context.Context, user *User) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).
		Model(&User{}).
		Where("user_id = ?", user.UserId).
		Select("WechatAvatar").
		Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *ModelUser) UpdateUserNickName(ctx context.Context, user *User) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).
		Model(&User{}).
		Where("user_id = ?", user.UserId).
		Select("WechatNickname").
		Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *ModelUser) QueryUserByUserId(ctx context.Context, user *User) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).
		Preload("Groups").
		Preload("UserGroups").
		Where("user_id = ?", user.UserId).
		First(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
