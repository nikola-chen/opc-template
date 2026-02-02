package overlord

import (
	"opc-template/backend/overlord/ai"
	"os"
)

func LoadSchema() (string, error) {
	data, err := os.ReadFile("design/schema.json")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (o *Overlord) SchemaFixFlow() Result {
	schema, err := LoadSchema()
	if err != nil {
		return NeedHuman
	}

	logs, _ := CollectLogs()

	patch, err := ai.GenerateSchemaPatch(schema, logs)
	if err != nil || len(patch) == 0 {
		return NeedHuman
	}

	// v0.1：先不真正 apply，只证明流程闭环
	if err := RunGenerate(); err != nil {
		return NeedHuman
	}

	if err := RetryRun(); err == nil {
		return Success
	}

	return NeedHuman
}
