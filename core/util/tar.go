package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateTarballGz(sourceDir, targetFile string) error {
	absDir, err := filepath.Abs(sourceDir)
	if err != nil {
		return fmt.Errorf("absolute path failed: %w", err)
	}

	absTarget, err := filepath.Abs(targetFile)
	if err != nil {
		return fmt.Errorf("absolute target path failed: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(absTarget), 0755); err != nil {
		return fmt.Errorf("creating target parent directory failed: %w", err)
	}

	cmd := exec.Command("tar",
		"-C", absDir,
		"-cvzf", absTarget,
		".",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("tar command failed: %w", err)
	}

	return nil
}
