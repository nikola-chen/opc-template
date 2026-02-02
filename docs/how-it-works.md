# How OPC Template Works

OPC Template 的核心思想是：

> **把“软件开发”拆成 AI 能理解、能重复执行的阶段**

---

## 整体流程

```text
Human Idea
   ↓
Design Schema (design/schema.json)
   ↓
AI Generate Code
   ↓
Run System
   ↓
Error?
   ↓
AI Heal (The Loop)
```

---

## 为什么是 Schema-first？

* Schema 是 **稳定输入**
* 代码是 **可再生结果**
* Bug 修复 ≠ 手改代码，而是修正 Schema 或生成逻辑

---

## The Loop（自愈闭环）是什么？

当你运行：

```bash
make heal
```

系统会：

1. 收集 `runtime/logs`
2. 分析错误
3. 修改生成代码
4. 再次运行

这是一个 **机器可以反复执行的闭环**。

---

## 你（人）应该做什么？

* 决定业务逻辑
* 修改 Schema
* 判断 AI 修复是否合理

---

## AI 不该做什么？

* 改设计意图
* 改 Schema 的“业务语义”
* 无限循环修复失败（有阈值）
