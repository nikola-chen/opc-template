# OPC CLI Reference

本文件是 **OPC Template 中 CLI 行为的权威定义**。
所有实现、Agent、自动化流程，**必须遵循本文件描述的语义**。

---

## CLI 总览

```bash
opc <command> [options]
```

OPC CLI 是 **软件生产流水线的控制面板**，而不是普通脚手架命令。

---

## 命令一览

| Command    | 作用               | 是否可重复 |
| ---------- | ---------------- | ----- |
| `design`   | 生成 / 更新设计 Schema | ✅     |
| `generate` | 根据 Schema 生成代码   | ✅     |
| `run`      | 启动系统             | ✅     |
| `heal`     | 自愈修复（The Loop）   | ✅     |
| `explain`  | 解释命令行为           | ✅     |
| `help`     | 帮助信息             | ✅     |

---

## `opc design`

### 作用

* 创建或更新 `design/schema.json`
* 不生成代码

### 输入

* 人类 Prompt
* 设计文件（可选）

### 输出

* `design/schema.json`

### 失败情况

* Schema 无法解析 → exit 1

---

## `opc generate`

### 作用

* 根据 Schema 生成工程代码

### 输入

* `design/schema.json`

### 输出

* `frontend/`
* `backend/`
* `infra/`

### 注意

* 该命令 **可反复执行**
* 会覆盖 AI-owned 目录

---

## `opc run`

### 作用

* 启动系统（前端 / 后端 / 基础设施）

### 输出

* 日志写入 `runtime/logs/`

### 失败情况

* 启动失败 → exit 1

---

## `opc heal`（The Loop）

### 作用

* 执行自愈闭环

### 行为流程

1. 读取 `runtime/logs`
2. 分析错误
3. 修改生成代码
4. 再次运行

### 终止条件

* 成功启动 → exit 0
* 超过修复次数 → exit 2（需要人工）

---

## `opc explain`

### 作用

* **不执行任何动作**
* 解释每个命令将做什么

### 适用场景

* 新用户理解流程
* AI Agent 规划执行顺序

---

## `opc help`

### 作用

* 显示简要帮助

---

## Exit Code 规范（非常重要）

| Code | 含义            |
| ---- | ------------- |
| 0    | 成功            |
| 1    | 可修复错误（AI 可介入） |
| 2    |               |
