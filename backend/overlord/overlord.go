package overlord

import (
	"opc-template/backend/overlord/ai"
)

type Result int

const (
	Success Result = iota
	NeedHuman
)

// Overlord は Self-Healing 的总控调度器（The Loop）
// 职责：
// - 收集运行日志
// - 进行错误分析
// - 决策修复策略（G3 Schema-first 优先，G2 Code-fix 兜底）
// - 驱动对应修复流程
type Overlord struct {
	MaxRetry int
	AIClient ai.Client
}

// Run 执行一次完整的自愈闭环
func (o *Overlord) Run() Result {
	// 1. 收集最新运行日志
	logs, err := CollectLogs()
	if err != nil {
		// 无日志或读取失败，无法判断，交由人工
		return NeedHuman
	}

	// 2. 分析错误类型，生成策略信号
	analysis := Analyze(logs)

	// 3. 策略决策（优先 Schema-first）
	switch {
	case analysis.IsSchemaRelated:
		// G3：Schema-first 修复路径
		return o.SchemaFixFlow()

	case analysis.AutoFixable:
		// G2：Code-fix 修复路径
		return o.CodeFixFlow()

	default:
		// 无法自动判断或风险过高
		return NeedHuman
	}
}

// CodeFixFlow G2: Code-fix 修复流程
func (o *Overlord) CodeFixFlow() Result {
	logs, err := CollectLogs()
	if err != nil {
		return NeedHuman
	}

	// Limit analysis to simple errors
	analysis := Analyze(logs)
	if !analysis.AutoFixable {
		return NeedHuman
	}

	// Apply fix
	if err := o.ApplyFix(analysis); err != nil {
		return NeedHuman
	}

	// Retry run to verify fix
	if err := RetryRun(); err == nil {
		return Success
	}

	return NeedHuman
}
