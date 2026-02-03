package ai

type FixRequest struct {
	Logs      string
	ErrorType string
	CodeScope string
}

type FixResponse struct {
	Patch  string // unified diff
	Reason string
}

type GenerateRequest struct {
	Prompt string
}

type Client interface {
	Fix(req FixRequest) (FixResponse, error)
	Generate(req GenerateRequest) (string, error)
	Name() string
}
