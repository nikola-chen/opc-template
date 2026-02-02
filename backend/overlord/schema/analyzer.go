package schema

import "strings"

type SchemaIssue struct {
	Path       string // JSONPath, e.g. $.entities.Order.fields.total
	Problem    string
	Suggestion string
}

func AnalyzeSchemaIssue(logs string, schema string) []SchemaIssue {
	// v0.1: 规则 + 关键词
	// v0.2: LLM 推理

	var issues []SchemaIssue

	if strings.Contains(logs, "column \"total\" does not exist") {
		issues = append(issues, SchemaIssue{
			Path:       "$.entities.Order.fields",
			Problem:    "Field 'total' missing",
			Suggestion: "Add field total:number",
		})
	}

	return issues
}
