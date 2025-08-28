# AI待办事项清单 (AI To-Do List)

一个基于Go语言构建的现代化待办事项管理API，采用分层架构设计。

## 🚀 项目概述

这是一个完整的全栈学习项目，旨在通过实践掌握：
- **后端开发**：Go + Gin + GORM
- **前端开发**：React（计划中）
- **基础设施**：Docker + AWS + CI/CD（计划中）
- **AI集成**：vLLM（计划中）

## ✨ 功能特性

- ✅ 完整的CRUD操作（创建、读取、更新、删除待办事项）
- ✅ RESTful API设计
- ✅ 分层架构（Models, Services, Controllers, Routes）
- ✅ PostgreSQL数据库 + GORM ORM
- ✅ 健康检查端点

## 🛠️ 技术栈

- **语言**: Go 1.25.0
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: PostgreSQL
- **容器化**: Docker

## 🚀 快速开始

### 前置要求
- Go 1.25.0+
- Docker
- PostgreSQL

### 1. 克隆项目
```bash
git clone https://github.com/Mateo0219/todo-backend.git
cd todo-backend
```

### 2. 启动数据库
```bash
docker run --name my-postgres \
  -e POSTGRES_PASSWORD=mysecretpassword \
  -e POSTGRES_DB=postgres \
  -p 5432:5432 \
  -d postgres
```

### 3. 运行项目
```bash
go mod tidy
go run main.go
```

服务器将在 `http://localhost:8080` 启动

## 📚 API接口

### 基础端点
- `GET /health` - 健康检查

### Todo API
- `GET /api/v1/todos` - 获取所有待办事项
- `POST /api/v1/todos` - 创建新待办事项
- `GET /api/v1/todos/:id` - 获取单个待办事项
- `PUT /api/v1/todos/:id` - 更新待办事项
- `DELETE /api/v1/todos/:id` - 删除待办事项

### 使用示例

#### 创建待办事项
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"学习Go分层架构","status":"pending"}'
```

#### 获取所有待办事项
```bash
curl http://localhost:8080/api/v1/todos
```

## 🏗️ 项目结构

```
todo-backend/
├── models/          # 数据模型层
├── services/        # 业务逻辑层
├── controllers/     # 控制器层
├── routes/          # 路由层
├── config/          # 配置层
├── middleware/      # 中间件层
├── utils/           # 工具函数层
└── main.go          # 主程序入口
```

## 📈 开发进度

- ✅ 基础CRUD API
- ✅ 分层架构重构
- ✅ 数据库集成
- 🔄 中间件集成（进行中）
- 📋 前端React应用（计划中）
- 📋 AI功能集成（计划中）

## 👨‍💻 作者

**全栈工程师学习者**

- GitHub: [@Mateo0219](https://github.com/Mateo0219)
- 项目链接: [https://github.com/Mateo0219/todo-backend](https://github.com/Mateo0219/todo-backend)

---

⭐ 如果这个项目对你有帮助，请给它一个星标！
