[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/nikola-chen/opc-template)](https://github.com/nikola-chen/opc-template/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nikola-chen/opc-template)](https://goreportcard.com/report/github.com/nikola-chen/opc-template)

# OPC Template

**One Person Company · AI 原生软件工厂模板** `v1.3.0`

> 一个用于构建 **“AI 可以长期维护的软件系统”** 的项目模板
> 面向：**独立开发者 / 一人公司 / 小型技术团队**

---

## 这不是普通模板

**OPC Template 不是脚手架，也不是一次性代码生成器。**

它提供的是一条 **AI 原生的软件生产流水线**：

> **Design → Schema → Code → Run → Error → Heal → Repeat**

在这条流水线上：

- 人类负责 **决策与设计**
- AI 负责 **生成、运行、修复**
- 项目可以长期演进，而不是“生成即废弃”

---

## 你能用它做什么？

- 构建并维护：
  - 微信小程序（可扩展 Web / 多端）
  - Go（Gin）+ DDD 风格后端
  - PostgreSQL / Redis / 异步架构

- 让 AI：
  - 根据设计生成代码
  - 运行系统
  - 读取报错日志
  - 在可控范围内自动修复代码（Self-Healing）

---

## 10 分钟快速开始（你将看到什么）

> 以下命令 **按顺序执行即可**

### 1️⃣ 初始化项目

```bash
make init
```

你会看到：

- 项目运行所需的基础结构就绪
- `.opc/` 与 `runtime/` 初始化完成

---

### 2️⃣ 生成设计 Schema

```bash
make design
```

你会看到：

- 生成或更新 `design/schema.json`

> **这是整个项目的“设计源头”**
> 后续一切代码，都会从这里派生。

---

### 3️⃣ 生成工程代码

```bash
make generate
```

你会看到：

- `frontend/`：前端工程代码
- `backend/`：Go 服务端代码
- `infra/`：数据库 / 中间件配置

这些目录 **可以被反复删除和重建**。

---

### 4️⃣ 启动系统

```bash
make run
```

你会看到：

- 系统启动
- 运行日志写入 `runtime/logs/`

---

### 5️⃣ 自愈修复（The Loop）

```bash
make heal
```

当系统启动失败时：

- AI 读取日志
- 分析错误
- 修改生成代码
- 再次尝试运行

> 注意：并非所有错误都可自动修复
> 当 CLI 返回 **exit code = 2** 时，表示需要人工介入做一次设计决策

---

## 项目结构与“所有权”（非常重要）

```text
opc-template/
├── design/        # ✅ 人可以改（核心设计）
├── frontend/      # ⚠️ AI 生成（尽量别手改）
├── backend/       # ⚠️ AI 生成
├── infra/         # ⚠️ AI 生成
├── runtime/       # 🤖 AI 使用（日志 / diff）
├── docs/          # ✅ 人读 / 人写
├── .opc/          # 🤖 AI 配置
└── Makefile       # 🤖 人 & AI 的统一入口
```

### 黄金法则

- ❌ 不要直接手改生成代码
- ✅ 修改 `design/schema.json`，然后重新 `make generate`

详细说明请阅读：
👉 `docs/ownership.md`

---

## CLI 使用说明（权威入口）

OPC 的所有能力都通过 CLI 暴露：

```bash
opc <command>
```

常用命令：

- `opc design`
- `opc generate`
- `opc run`
- `opc heal`
- `opc explain`
- `opc help`

📘 **CLI 的唯一权威说明（Single Source of Truth）：**
👉 `docs/cli.md`

---

### 关于 `make` 与 `opc` 的关系

- `opc`：**核心执行器**，定义所有稳定的命令语义
- `Makefile`：**流程编排器（Orchestrator）**，将多个 `opc` 命令串成完整流水线

> 对新用户：**优先使用 `make`**
> 对自动化 / Agent：**直接调用 `opc`**

---

## 我应该先读哪些文档？

推荐顺序：

1. **`docs/how-it-works.md`**
   理解整条 AI 软件流水线如何运作
2. **`docs/cli.md`**
   CLI 的行为与契约
3. **`docs/ownership.md`**
   文件所有权与修改边界

---

## 如果你卡住了怎么办？

1. 查看 `runtime/logs/`
2. 阅读 `docs/how-it-works.md`
3. 尝试：

   ```bash
   make heal
   ```

如果问题依然存在，说明：

> **需要人类做一次设计决策**

---

## 这个模板适合谁？

- 想把 AI 当“长期队友”的开发者
- 想降低维护成本的一人公司（OPC）
- 想实践 AI 原生工程方法的人

---

## 你现在拥有的是什么？

> 一个 **可被 AI 维护、可复制、可长期演进的软件起点**

---

## 下一步你可以做什么？

- 接入真正的 AI Foundation（多模型 / MCP / Agent）
- 实现完整的 Go 自愈监工（Overlord）
- 用本模板生成你的第一个真实业务项目

---

**Welcome to OPC.**
_Build once. Let AI maintain it._
