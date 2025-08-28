package models

import (
	"gorm.io/gorm"
	"time"
)

// Todo 表示一个待办事项
type Todo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Title     string         `json:"title" gorm:"not null"`
	Status    string         `json:"status" gorm:"default:'pending'"`
}

// TableName 指定数据库表名
func (Todo) TableName() string {
	return "todos"
}

// BeforeCreate GORM钩子，在创建记录前执行
func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	if t.Status == "" {
		t.Status = "pending"
	}
	return nil
}

// IsValid 验证Todo数据是否有效
func (t *Todo) IsValid() bool {
	return t.Title != "" && len(t.Title) <= 255
}
