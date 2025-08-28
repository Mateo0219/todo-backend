package controllers

import (
	"net/http"
	"strconv"
	"todo-backend/models"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

// TodoController 处理Todo相关的HTTP请求
type TodoController struct {
	todoService services.TodoService
}

// NewTodoController 创建新的TodoController实例
func NewTodoController(todoService services.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

// CreateTodo 创建新的待办事项
func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求数据无效",
			"details": err.Error(),
		})
		return
	}

	if err := c.todoService.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建待办事项失败",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "待办事项创建成功",
		"data":    todo,
	})
}

// GetAllTodos 获取所有待办事项
func (c *TodoController) GetAllTodos(ctx *gin.Context) {
	todos, err := c.todoService.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取待办事项列表失败",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": todos,
		"count": len(todos),
	})
}

// GetTodoByID 根据ID获取待办事项
func (c *TodoController) GetTodoByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID格式",
		})
		return
	}

	todo, err := c.todoService.GetTodoByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "待办事项未找到",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

// UpdateTodo 更新待办事项
func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID格式",
		})
		return
	}

	var updates map[string]interface{}
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求数据无效",
			"details": err.Error(),
		})
		return
	}

	if err := c.todoService.UpdateTodo(uint(id), updates); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "更新待办事项失败",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "待办事项更新成功",
	})
}

// DeleteTodo 删除待办事项
func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID格式",
		})
		return
	}

	if err := c.todoService.DeleteTodo(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除待办事项失败",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "待办事项删除成功",
	})
}
