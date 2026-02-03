package patch

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Apply(diff string) error {
	if !IsValidUnifiedDiff(diff) {
		return errors.New("invalid diff format")
	}

	cmd := exec.Command("patch", "-p1")
	cmd.Stdin = strings.NewReader(diff)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("patch failed: %w, stderr: %s", err, stderr.String())
	}

	return nil
}
