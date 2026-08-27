# 灰度发布系统（gray-release）

一个基于纯 Go 标准库（`net/http`，零第三方依赖）的灰度发布后端，提供服务版本管理、灰度发布单、流量规则、白名单、分步放量与变更记录。内置灰度流量判定（白名单优先 + 按优先级匹配规则），支持发布单状态机、分步放量与快照导入导出，并提供原生前端控制台。

## 运行

```bash
cd origin
go run ./cmd/server
# 默认监听 :8080，可用 PORT / ADDR 覆盖
# 可选环境变量：AUTH_TOKEN（启用 Bearer 鉴权）、RATE_LIMIT（每 IP 每分钟限流，默认 200）
```

访问前端控制台：`http://localhost:8080/`（需在 `origin/` 目录下启动）。

## 分层结构

```
origin/
├── cmd/server/main.go        # 入口 + 前端挂载 + 优雅关闭
├── frontend/index.html       # 原生前端控制台（零构建）
├── internal/
│   ├── app/ config/ model/ store/ service/ handler/
└── pkg/ httpx/ idgen/ logger/ semver/ strutil/
```

## 核心概念

- **Version**：服务版本（制品地址、校验和、大小），同服务内版本号唯一。
- **Release**：灰度发布单，状态机 `draft → creating → releasing → completed`，可 `rolled_back` / `cancelled`。
- **TrafficRule**：流量规则，类型 `percentage` / `header` / `cookie` / `user`，按 `priority` 升序匹配。
- **Whitelist**：白名单用户，灰度判定时优先放行。
- **ReleaseStep**：放量步骤（目标百分比），逐步推进。
- **RolloutRecord**：放量历史（起始/目标百分比）。
- **ChangeLog**：发布单变更记录。

## API 一览（核心）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST/GET | /api/versions | 创建 / 列表（service_id 筛选） |
| GET/PUT/DELETE | /api/versions/{id} | 版本详情 / 更新 / 删除 |
| POST/GET | /api/releases | 创建 / 列表（service_id/status/keyword） |
| GET/PUT/DELETE | /api/releases/{id} | 发布单详情 / 更新 / 删除 |
| PATCH | /api/releases/{id}/status | 状态机流转 |
| POST/GET | /api/traffic-rules | 流量规则创建 / 列表 |
| PUT | /api/traffic-rules/{id} | 更新规则 |
| PATCH | /api/traffic-rules/{id}/toggle | 启用 / 停用 |
| DELETE | /api/traffic-rules/{id} | 删除规则 |
| POST/GET/DELETE | /api/whitelists | 白名单增 / 查 / 删 |
| POST/GET | /api/releases/{id}/steps | 添加 / 查看放量步骤 |
| POST | /api/releases/{id}/advance | 推进放量步骤 |
| GET | /api/releases/{id}/progress | 当前放量进度 |
| GET | /api/releases/{id}/records | 放量记录 |
| GET | /api/releases/{id}/overview | 放量概览 |
| POST | /api/releases/{id}/evaluate | 灰度判定 |
| GET | /api/change-logs | 变更记录（分页） |
| GET | /api/stats/releases | 发布单统计 |
| GET | /api/stats/active-releases | 进行中的发布单 |
| GET | /api/export · POST /api/import | 快照导出 / 导入 |
| GET | /healthz · /readyz | 存活 / 就绪探针 |

## 灰度判定逻辑

1. 白名单命中 → 直接放行；
2. 按优先级遍历已启用的流量规则：`user`（用户 ID 精确匹配）、`header` / `cookie`（键值匹配）、`percentage`（用户 ID 经 FNV 哈希落入 0-99 桶，小于百分比即命中）。

## 统一响应

```json
{ "code": 0, "message": "ok", "data": { } }
```
