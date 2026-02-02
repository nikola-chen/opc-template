package overlord

import (
	"errors"
	"opc-template/backend/overlord/ai"
	"opc-template/backend/overlord/patch"
)

// ApplyFix is where AI or rule-based patching happens
func ApplyFix(a Analysis) error {
	logs, _ := CollectLogs()

	req := ai.FixRequest{
		Logs:      logs,
		ErrorType: string(a.Type),
		CodeScope: "backend",
	}

	router := ai.Router{
		Primary:   NewClaudeClient(), // 你可实现
		Secondary: NewGPTClient(),    // 可选
	}

	resp, err := router.Fix(req)
	if err != nil {
		return err
	}

	if resp.Patch == "" {
		return errors.New("model returned empty patch")
	}

	if !patch.IsValidUnifiedDiff(resp.Patch) {
		return errors.New("invalid patch from model")
	}

	return patch.Apply(resp.Patch)
}

// Clients Factory

func NewClaudeClient() ai.Client {
	return &MockClient{name: "Claude-3.5-Sonnet"}
}

func NewGPTClient() ai.Client {
	return &MockClient{name: "GPT-4o"}
}

type MockClient struct {
	name string
}

func (m *MockClient) Convert(req ai.FixRequest) (ai.FixResponse, error) {
	// Legacy method if used anywhere, but interface uses Fix now
	return m.Fix(req)
}

func (m *MockClient) Fix(req ai.FixRequest) (ai.FixResponse, error) {
	// Simple mock that returns a basic patch if the error is "Undefined symbol"
	if req.ErrorType == "COMPILE_ERROR" {
		// Assuming we can guess what to fix. For now, return empty or dummy.
		// In a real scenario, this would call the LLM API.
		// We will return an error so it falls back or Requires Human,
		// effectively simulating "AI didn't know how to fix it" unless we have a specific test case.
		return ai.FixResponse{
			Patch:  "",
			Reason: "Mock client: not implemented real LLM call yet",
		}, nil
	}
	return ai.FixResponse{}, nil
}

func (m *MockClient) Name() string {
	return m.name
}
