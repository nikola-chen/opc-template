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

type Client interface {
	Fix(req FixRequest) (FixResponse, error)
	Name() string
}
