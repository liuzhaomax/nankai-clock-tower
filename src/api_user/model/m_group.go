package model

import (
	"context"
	"github.com/liuzhaomax/go-maxms/internal/core"
	"gorm.io/gorm"
)

func (m *ModelUser) CreateGroup(ctx context.Context, group *Group) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Create(group)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *ModelUser) QueryGroupByName(ctx context.Context, group *Group) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Where("name = ?", group.Name).First(group)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *ModelUser) CreateUserGroup(ctx context.Context, userGroup *UserGroup) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Create(userGroup)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
