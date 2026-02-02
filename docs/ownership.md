# File Ownership Contract

为了让 AI 能长期维护项目，必须明确 **谁拥有哪一部分**。

---

## Human-owned（人拥有）

你可以、也应该修改：

* `design/schema.json`
* `docs/`
* 明确标注的业务决策区

---

## AI-owned（AI 拥有）

不要直接手改：

* `frontend/`
* `backend/`
* `infra/`
* `runtime/`

这些目录里的代码：

* 可以被删除
* 可以被重建
* 不保证稳定

---

## Why？

因为：

> **AI 只能维护“它生成的东西”**

---

## 推荐修改流程

❌ 错误：

```text
直接改 backend 代码
```

✅ 正确：

```text
改 Schema → make generate
```
