# =========================================

# OPC · One Person Company

# AI 原生软件工厂 - Makefile

# =========================================

PROJECT_NAME := opc-template
ENV ?= dev

# ---------- 基础路径 ----------

DESIGN_DIR := design
FRONTEND_DIR := frontend
BACKEND_DIR := backend
INFRA_DIR := infra
RUNTIME_DIR := runtime

OPC_CONFIG := .opc/config.yaml

# ---------- Foundation ----------

OPC_FOUNDATION ?= ../opc-foundation

# ---------- 默认 ----------

.PHONY: help
help:
@echo ""
@echo "OPC · One Person Company"
@echo "----------------------------------"
@echo "make init        初始化项目"
@echo "make design      生成 / 更新 Design Schema"
@echo "make generate    AI 生成前后端代码"
@echo "make run         启动系统"
@echo "make test        运行测试"
@echo "make heal        自愈修复（The Loop）"
@echo "make clean       清理运行状态"
@echo ""

# ---------- 初始化 ----------

.PHONY: init
init:
@echo ">>> 初始化 OPC 项目"
mkdir -p $(DESIGN_DIR) $(RUNTIME_DIR)
cp -n $(OPC_FOUNDATION)/schema/base.schema.json $(DESIGN_DIR)/schema.json || true
@echo "✔ 初始化完成"

# ---------- 设计阶段 ----------

.PHONY: design
design:
@echo ">>> Design → Schema"
@go run $(OPC_FOUNDATION)/cli design 
--input prompt.txt 
--output $(DESIGN_DIR)/schema.json 
--config $(OPC_CONFIG)
@echo "✔ Schema 已生成"

# ---------- 代码生成 ----------

.PHONY: generate
generate:
@echo ">>> Schema → Code"
@go run $(OPC_FOUNDATION)/cli generate 
--schema $(DESIGN_DIR)/schema.json 
--frontend $(FRONTEND_DIR) 
--backend $(BACKEND_DIR) 
--infra $(INFRA_DIR) 
--runtime $(RUNTIME_DIR) 
--config $(OPC_CONFIG)
@echo "✔ 代码生成完成"

# ---------- 启动 ----------

.PHONY: run
run:
@echo ">>> 启动系统"
docker compose up -d
cd $(BACKEND_DIR) && go run cmd/server/main.go

# ---------- 测试 ----------

.PHONY: test
test:
@echo ">>> 运行测试"
cd $(BACKEND_DIR) && go test ./...
@echo "✔ 测试完成"

# ---------- 自愈闭环 ----------

.PHONY: heal
heal:
@echo ">>> The Loop · 自愈修复"
@go run $(OPC_FOUNDATION)/agents/overlord 
--logs $(RUNTIME_DIR)/logs 
--code $(BACKEND_DIR) 
--frontend $(FRONTEND_DIR) 
--schema $(DESIGN_DIR)/schema.json 
--config $(OPC_CONFIG)
@echo "✔ 自愈完成"

# ---------- 清理 ----------

.PHONY: clean
clean:
@echo ">>> 清理运行状态"
rm -rf $(RUNTIME_DIR)
docker compose down
@echo "✔ 清理完成"