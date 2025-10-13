// SPDX-License-Identifier: AGPL-3.0-or-later
package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ExecCmd(prefix string, cmdline []string, wd string) error {
	log.Printf("[%s] EXEC: %s\n", prefix, cmdline)
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = wd
	return cmd.Run()
}

func ExecOut(cmdline []string) string {
	out, err := exec.Command(cmdline[0], cmdline[1:]...).Output()
	if err != nil {
		fmt.Println("Exec error:", err)
	}
	return strings.TrimSpace(string(out))
}

func ExecRetcode(cmdline []string, wd string) int {
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Dir = wd
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
