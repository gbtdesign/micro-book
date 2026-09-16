# micro-book

Go 微服务项目，涵盖从语法基础到完整微服务架构的全栈实践。

## 项目结构

```
micro-book/
├── syntax/          # Go 语言基础语法示例
├── gin/             # Gin Web 框架示例
├── gorm/            # GORM ORM 示例
├── grpc/            # gRPC 服务通信示例
├── mongo/           # MongoDB CRUD 示例
├── sarama/          # Kafka 消息队列示例（Sarama）
├── opentelemetry/   # OpenTelemetry 链路追踪示例
├── cronjob/         # 定时任务示例
├── wire/            # Wire 依赖注入示例
├── context/         # Context 使用示例
├── webook/          # 主项目 - WeBook 微服务应用
└── webook-fe/       # WeBook 前端（Next.js）
```

## 核心项目：WeBook

WeBook 是一个基于微服务架构的在线内容平台，包含以下特性：

### 技术栈

| 分类 | 技术 |
|------|------|
| 语言 | Go 1.25 |
| Web 框架 | Gin |
| ORM | GORM |
| 微服务通信 | gRPC |
| 服务注册与发现 | etcd / go-zero / Kratos |
| 消息队列 | Kafka（Sarama） |
| 数据库 | MySQL 8.0 |
| 缓存 | Redis |
| 文档数据库 | MongoDB |
| 依赖注入 | Wire |
| 配置管理 | Viper（支持远程 etcd 配置） |
| 日志 | Zap |
| 链路追踪 | OpenTelemetry + Zipkin |
| 监控 | Prometheus + Grafana |
| 容器化 | Docker + Kubernetes |
| 认证 | JWT / Session |
| 限流 | 自定义限流器 |
| 定时任务 | robfig/cron |
| 前端 | Next.js + Tailwind CSS |

### 架构设计

```
webook/
├── api/             # API 定义（Proto & HTTP）
├── config/          # 配置文件
├── ioc/             # IoC 容器（依赖注入初始化）
├── internal/        # 内部业务逻辑
│   ├── domain/      # 领域模型
│   ├── errs/        # 错误定义
│   ├── events/      # 事件消费者
│   ├── job/         # 定时任务
│   ├── repository/  # 数据访问层
│   │   ├── cache/   # Redis 缓存层
│   │   └── dao/     # 数据库 DAO 层
│   ├── service/     # 业务服务层
│   └── web/         # HTTP Handler 层
├── interactive/     # 互动模块（独立 gRPC 服务）
├── migrator/        # 数据迁移工具
├── pkg/             # 公共工具包
│   ├── ginx/        # Gin 扩展
│   ├── gormx/       # GORM 扩展
│   ├── grpcx/       # gRPC 扩展
│   ├── logger/      # 日志封装
│   ├── migrator/    # 迁移框架
│   ├── ratelimit/   # 限流器
│   ├── redisx/      # Redis 扩展
│   ├── saramax/     # Kafka 扩展
│   └── zapx/        # Zap 扩展
└── script/          # 数据库初始化脚本
```

## 快速开始

### 环境要求

- Go 1.25+
- Docker & Docker Compose
- Node.js（前端项目）

### 启动基础设施

```bash
cd webook
docker-compose up -d
```

这将启动以下服务：

| 服务 | 端口 |
|------|------|
| MySQL | 13316 |
| Redis | 6379 |
| etcd | 12379 |
| MongoDB | 27017 |
| Kafka | 9094 |
| Prometheus | 9090 |
| Grafana | 3000 |
| Zipkin | 9411 |

### 启动后端服务

```bash
cd webook
go run . --config=config/dev.yaml
```

服务将在 `:8080` 端口启动，Prometheus metrics 暴露在 `:8081`。

### 启动前端

```bash
cd webook-fe
npm install
npm run dev
```

## 开发工具

### 生成 Mock

```bash
make mock
```

### 生成 gRPC 代码

```bash
make grpc
```

## Kubernetes 部署

项目包含完整的 K8s 部署配置：

- `k8s-mysql-*.yaml` - MySQL 持久化存储与服务
- `k8s-redis-*.yaml` - Redis 部署与服务
- `k8s-webook-*.yaml` - WeBook 应用部署与服务
- `k8s-ingress-nginx.yaml` - Ingress 配置

## 学习模块说明

| 目录 | 内容 |
|------|------|
| `syntax/` | Go 基础语法：类型、控制流、函数、泛型、并发 |
| `gin/` | Gin 框架入门 |
| `gorm/` | GORM 数据库操作 |
| `grpc/` | gRPC 服务端/客户端、负载均衡、限流、故障转移 |
| `sarama/` | Kafka 生产者与消费者 |
| `mongo/` | MongoDB CRUD 操作 |
| `opentelemetry/` | 分布式链路追踪 |
| `cronjob/` | 定时任务调度 |
| `wire/` | Google Wire 依赖注入 |
| `context/` | Go Context 使用 |
