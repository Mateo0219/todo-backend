package services

import (
	"errors"
	"todo-backend/models"
	"gorm.io/gorm"
)

// TodoService 定义Todo相关的业务逻辑接口
type TodoService interface {
	CreateTodo(todo *models.Todo) error
	GetAllTodos() ([]models.Todo, error)
	GetTodoByID(id uint) (*models.Todo, error)
	UpdateTodo(id uint, updates map[string]interface{}) error
	DeleteTodo(id uint) error
}

// todoService 实现TodoService接口
type todoService struct {
	db *gorm.DB
}

// NewTodoService 创建新的TodoService实例
func NewTodoService(db *gorm.DB) TodoService {
	return &todoService{db: db}
}

// CreateTodo 创建新的待办事项
func (s *todoService) CreateTodo(todo *models.Todo) error {
	if !todo.IsValid() {
		return errors.New("待办事项数据无效")
	}
	
	return s.db.Create(todo).Error
}

// GetAllTodos 获取所有待办事项
func (s *todoService) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := s.db.Find(&todos).Error
	return todos, err
}

// GetTodoByID 根据ID获取待办事项
func (s *todoService) GetTodoByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := s.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// UpdateTodo 更新待办事项
func (s *todoService) UpdateTodo(id uint, updates map[string]interface{}) error {
	var todo models.Todo
	if err := s.db.First(&todo, id).Error; err != nil {
		return err
	}
	
	return s.db.Model(&todo).Updates(updates).Error
}

// DeleteTodo 删除待办事项
func (s *todoService) DeleteTodo(id uint) error {
	var todo models.Todo
	if err := s.db.First(&todo, id).Error; err != nil {
		return err
	}
	
	return s.db.Delete(&todo).Error
}
