package overlord

import (
	"os"
	"os/exec"
	"path/filepath"
)

func RunGenerate() error {
	cmd := exec.Command("make", "generate")
	return cmd.Run()
}

func RetryRun() error {
	cmd := exec.Command("make", "run")
	return cmd.Run()
}

func CollectLogs() (string, error) {
	logDir := "runtime/logs"
	entries, err := os.ReadDir(logDir)
	if err != nil {
		return "", err
	}

	if len(entries) == 0 {
		return "", nil
	}

	latest := entries[len(entries)-1]
	data, err := os.ReadFile(filepath.Join(logDir, latest.Name()))
	if err != nil {
		return "", err
	}

	return string(data), nil
}
