package patch

import (
	"errors"
	"os/exec"
	"strings"
)

func Apply(diff string) error {
	if !IsValidUnifiedDiff(diff) {
		return errors.New("invalid diff format")
	}

	cmd := exec.Command("patch", "-p1")
	cmd.Stdin = strings.NewReader(diff)
	return cmd.Run()

}
