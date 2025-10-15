package util

import (
	"fmt"
	"log"
	"strings"
	"os"
	"os/exec"
)

func ExecCmd(cmdline []string) error {
	log.Printf("EXEC: %s\n", cmdline)
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ExecOut(cmdline []string) string {
	out, err := exec.Command(cmdline[0], cmdline[1:]...).Output()
	if err != nil {
		fmt.Println("Exec error:", err)
	}
	return strings.TrimSpace(string(out))
}

func ExecRetcode(cmdline []string) int {
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		} else {
			fmt.Println("Failed to run command:", err)
			return 127
		}
	}
	return 0
}
