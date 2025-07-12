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

// DeleteUserGroup 硬删除用户群组关联
func (m *ModelUser) DeleteUserGroup(ctx context.Context, userGroup *UserGroup) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Unscoped().Delete(userGroup)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *ModelUser) QueryGroupWithUsersByName(ctx context.Context, group *Group) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).
		Preload("Users").
		Where("name = ?", group.Name).
		First(group)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// DeleteGroup 硬删除群组
func (m *ModelUser) DeleteGroup(ctx context.Context, group *Group) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Unscoped().Delete(group)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
