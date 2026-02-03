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

	// Find latest file by modification time
	var latestEntry os.DirEntry
	var latestTime int64 = 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if info.ModTime().Unix() > latestTime {
			latestTime = info.ModTime().Unix()
			latestEntry = entry
		}
	}

	if latestEntry == nil {
		return "", nil
	}

	data, err := os.ReadFile(filepath.Join(logDir, latestEntry.Name()))
	if err != nil {
		return "", err
	}

	return string(data), nil
}
