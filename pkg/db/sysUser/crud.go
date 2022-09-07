package sysUser

import (
	"context"
	"gorm.io/gorm"
	"rbac-demo/pkg/models"
	"time"
)

type SysUserInterface interface {
	Create(sysUser *models.SysUser) error
	Get(ctx context.Context, uid int) (*models.SysUser, error)
	Update(ctx context.Context, updates map[string]interface{})
	Delete(ctx context.Context, uid int) error
	GetByName(ctx context.Context, name string) (*models.SysUser, error)
}

type DB struct {
	db *gorm.DB
}

func NewSysUser(db *gorm.DB) SysUserInterface {
	return &DB{db}
}

func (g *DB) Create(sysUser *models.SysUser) error {
	now := time.Now()
	sysUser.CreateTime = now
	return g.db.Create(&sysUser).Error
}

func (g *DB) Get(ctx context.Context, uid int) (*models.SysUser, error) {
	var sysUser *models.SysUser
	if err := g.db.Where("user_id = ?", uid).First(&sysUser).Error; err != nil {
		return nil, err
	}
	return sysUser, nil
}

func (g *DB) Update(ctx context.Context, updates map[string]interface{}) {
	g.db.Model(&models.SysUser{}).Updates(updates)
}

func (g *DB) Delete(ctx context.Context, uid int) error {
	var sysUser *models.SysUser
	if err := g.db.Where("user_id = ?", uid).Delete(&sysUser).Error; err != nil {
		return err
	}
	return nil
}

func (g *DB) GetByName(ctx context.Context, name string) (*models.SysUser, error) {
	var sysUser *models.SysUser
	if err := g.db.Where("username = ?", name).First(&sysUser).Error; err != nil {
		return nil, err
	}
	return sysUser, nil
}
