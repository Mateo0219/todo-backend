package config

import (
	"log"
	"todo-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase() error {
	dbConfig := GetDatabaseConfig()
	
	var err error
	DB, err = gorm.Open(postgres.Open(dbConfig.GetDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("数据库连接成功")
	
	// 自动迁移数据库表结构
	if err := DB.AutoMigrate(&models.Todo{}); err != nil {
		return err
	}
	
	log.Println("数据库表结构迁移完成")
	return nil
}

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return DB
}

// CloseDatabase 关闭数据库连接
func CloseDatabase() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
