package ai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client *openai.Client
	model  string
}

func NewOpenAIClient() (*OpenAIClient, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	config := openai.DefaultConfig(apiKey)
	// You can customize the base URL if needed, e.g., for aggressive caching proxies or enterprise endpoints
	if baseURL := os.Getenv("OPENAI_BASE_URL"); baseURL != "" {
		config.BaseURL = baseURL
	}

	client := openai.NewClientWithConfig(config)

	model := os.Getenv("OPC_MODEL")
	if model == "" {
		model = openai.GPT4o // Default to a smart model
	}

	return &OpenAIClient{
		client: client,
		model:  model,
	}, nil
}

func (c *OpenAIClient) Name() string {
	return "OpenAI/" + c.model
}

func (c *OpenAIClient) Fix(req FixRequest) (FixResponse, error) {
	prompt := BuildFixPrompt(req.Logs)
	if req.CodeScope != "" {
		prompt += fmt.Sprintf("\n\nContext Code:\n%s", req.CodeScope)
	}

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: c.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are an expert Go programmer and troubleshooting agent.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0.1, // Low temperature for deterministic fixes
		},
	)

	if err != nil {
		return FixResponse{}, err
	}

	if len(resp.Choices) == 0 {
		return FixResponse{}, fmt.Errorf("no response choices from AI")
	}

	content := resp.Choices[0].Message.Content

	// Basic parsing: assume the model returns just the patch or explanation + patch
	// In a real implementation, we might want more structured output (e.g. JSON mode)

	return FixResponse{
		Patch:  content,
		Reason: "AI generated fix based on logs",
	}, nil
}

func (c *OpenAIClient) Generate(req GenerateRequest) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: c.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are an expert full-stack developer.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: req.Prompt,
				},
			},
			Temperature: 0.7, // Higher temperature for creative generation
		},
	)

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices from AI")
	}

	return resp.Choices[0].Message.Content, nil
}

// Ensure OpenAIClient implements Client
var _ Client = (*OpenAIClient)(nil)
