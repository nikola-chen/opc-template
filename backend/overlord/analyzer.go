package overlord

import "strings"

type ErrorType string

const (
	NoError      ErrorType = "NO_ERROR"
	CompileError ErrorType = "COMPILE_ERROR"
	RuntimeError ErrorType = "RUNTIME_ERROR"
	UnknownError ErrorType = "UNKNOWN_ERROR"
)

type Analysis struct {
	Type            ErrorType
	Message         string
	AutoFixable     bool
	IsSchemaRelated bool
}

func Analyze(logs string) Analysis {
	if logs == "" {
		return Analysis{Type: NoError, AutoFixable: false, IsSchemaRelated: false}
	}

	isSchema := strings.Contains(logs, "schema validation failed") || strings.Contains(logs, "field mismatch")

	if strings.Contains(logs, "undefined") {
		return Analysis{
			Type:            CompileError,
			Message:         "Undefined symbol",
			AutoFixable:     true,
			IsSchemaRelated: isSchema,
		}
	}

	if strings.Contains(logs, "panic") {
		return Analysis{
			Type:            RuntimeError,
			Message:         "Runtime panic",
			AutoFixable:     false,
			IsSchemaRelated: isSchema,
		}
	}

	return Analysis{
		Type:            UnknownError,
		Message:         logs,
		AutoFixable:     false,
		IsSchemaRelated: isSchema,
	}
}
