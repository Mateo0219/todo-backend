package routes

import (
	"todo-backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupTodoRoutes 设置Todo相关的路由
func SetupTodoRoutes(router *gin.Engine, todoController *controllers.TodoController) {
	// 创建API v1路由组
	v1 := router.Group("/api/v1")
	{
		// Todo相关路由
		todos := v1.Group("/todos")
		{
			todos.POST("", todoController.CreateTodo)           // 创建待办事项
			todos.GET("", todoController.GetAllTodos)          // 获取所有待办事项
			todos.GET("/:id", todoController.GetTodoByID)      // 获取单个待办事项
			todos.PUT("/:id", todoController.UpdateTodo)       // 更新待办事项
			todos.DELETE("/:id", todoController.DeleteTodo)    // 删除待办事项
		}
	}
}

// SetupHealthCheck 设置健康检查路由
func SetupHealthCheck(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "服务运行正常",
		})
	})
}
