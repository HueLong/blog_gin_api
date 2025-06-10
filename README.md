# Blog Gin API

基于 Golang 和 Gin 框架的高性能、高可用 RESTful API 服务。

## 项目特点

- 高性能：使用 Gin 框架，支持高并发
- 高可用：完善的错误处理和日志系统
- 可复用：模块化设计，易于扩展
- 可读性强：清晰的代码结构和注释

## 项目结构

```
.
├── cmd/                    # 主程序入口
├── config/                 # 配置文件
├── internal/              # 内部包
│   ├── handler/          # HTTP 处理器
│   ├── middleware/       # 中间件
│   ├── model/           # 数据模型
│   ├── repository/      # 数据访问层
│   ├── service/         # 业务逻辑层
│   └── pkg/             # 内部工具包
├── pkg/                  # 公共工具包
├── scripts/             # 脚本文件
└── test/                # 测试文件
```

## 快速开始

1. 克隆项目
```bash
git clone https://github.com/yourusername/blog_gin_api.git
```

2. 安装依赖
```bash
go mod download
```

3. 运行项目
```bash
go run cmd/main.go
```

## 配置说明

配置文件位于 `config/config.yaml`，包含以下配置项：

- 服务器配置
- 数据库配置
- 日志配置
- 其他配置项

## API 文档

API 文档使用 Swagger 生成，访问地址：`http://localhost:8080/swagger/index.html`

## 测试

运行测试：
```bash
go test ./...
```

## 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License 