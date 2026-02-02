package ai

type SchemaPatchOp struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

func GenerateSchemaPatch(schema string, logs string) ([]SchemaPatchOp, error) {
	// v0.1: stub（后续接 LLM）
	// v0.2: 调用 BuildSchemaFixPrompt + LLM

	return []SchemaPatchOp{}, nil
}
