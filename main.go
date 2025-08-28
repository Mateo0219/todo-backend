package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"todo-backend/config"
	"todo-backend/controllers"
	"todo-backend/routes"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化数据库
	if err := config.InitDatabase(); err != nil {
		log.Fatal("数据库初始化失败:", err)
	}
	defer config.CloseDatabase()

	// 创建Gin路由器
	router := gin.Default()

	// 设置中间件
	setupMiddleware(router)

	// 创建服务层实例
	todoService := services.NewTodoService(config.GetDB())

	// 创建控制器实例
	todoController := controllers.NewTodoController(todoService)

	// 设置路由
	routes.SetupHealthCheck(router)
	routes.SetupTodoRoutes(router, todoController)

	// 启动服务器
	go startServer(router)

	// 等待中断信号
	waitForShutdown()

	log.Println("服务器已关闭")
}

// setupMiddleware 设置中间件
func setupMiddleware(router *gin.Engine) {
	// 这里可以添加日志、CORS、认证等中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}

// startServer 启动HTTP服务器
func startServer(router *gin.Engine) {
	log.Println("服务器启动中...")
	log.Println("API地址: http://localhost:8080")
	log.Println("健康检查: http://localhost:8080/health")
	log.Println("API文档: http://localhost:8080/api/v1/todos")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}

// waitForShutdown 等待关闭信号
func waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在关闭服务器...")

	// 给服务器一些时间来优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 这里可以添加优雅关闭逻辑
	select {
	case <-ctx.Done():
		log.Println("服务器强制关闭")
	}
}