package overlord

import (
	"errors"
	"fmt"
	"opc-template/backend/overlord/ai"
	"opc-template/backend/overlord/patch"
)

// ApplyFix is where AI or rule-based patching happens
func (o *Overlord) ApplyFix(a Analysis) error {
	logs, _ := CollectLogs()

	if o.AIClient == nil {
		return errors.New("no AI client configured")
	}

	req := ai.FixRequest{
		Logs:      logs,
		ErrorType: string(a.Type),
		CodeScope: "backend",
	}

	resp, err := o.AIClient.Fix(req)
	if err != nil {
		return fmt.Errorf("AI fix failed: %w", err)
	}

	if resp.Patch == "" {
		return errors.New("model returned empty patch")
	}

	if !patch.IsValidUnifiedDiff(resp.Patch) {
		return errors.New("invalid patch from model")
	}

	return patch.Apply(resp.Patch)
}
