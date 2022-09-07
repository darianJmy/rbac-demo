package db

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"rbac-demo/pkg/db/sysUser"
	"rbac-demo/pkg/models"
	"time"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func Connection(data []byte) (*gorm.DB, error) {
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	} else {
		dbConfig, _ := db.DB()
		dbConfig.SetMaxOpenConns(0)
		dbConfig.SetMaxIdleConns(10)
		dbConfig.SetConnMaxLifetime(time.Hour)
	}

	if !db.Migrator().HasTable(&models.SysUser{}) {
		if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&models.SysUser{}); err != nil {
			return nil, err
		}
	}
	if !db.Migrator().HasTable(&models.SysRole{}) {
		if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&models.SysRole{}); err != nil {
			return nil, err
		}
	}
	if !db.Migrator().HasTable(&models.SysMenu{}) {
		if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&models.SysMenu{}); err != nil {
			return nil, err
		}
	}
	if !db.Migrator().HasTable(&models.SysUserRole{}) {
		if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&models.SysUserRole{}); err != nil {
			return nil, err
		}
	}
	if !db.Migrator().HasTable(&models.SysRoleMenu{}) {
		if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			Migrator().CreateTable(&models.SysRoleMenu{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}

type ShareDaoFactory interface {
	SysUser() sysUser.SysUserInterface
}

type shareDaoFactory struct {
	db *gorm.DB
}

func (f *shareDaoFactory) SysUser() sysUser.SysUserInterface {
	return sysUser.NewSysUser(f.db)
}

func NewDaoFactory(db *gorm.DB) ShareDaoFactory {
	return &shareDaoFactory{
		db: db,
	}
}
