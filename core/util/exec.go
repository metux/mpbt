package util

import (
	"log"
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
