package main

import (
	"fmt"
	"strings"
)

func parseDefine(s string) (name, value string, ok bool) {
	const prefix = "" // -D doesn't work with flags package
	if !strings.HasPrefix(s, prefix) {
		return "", "", false
	}

	remainder := strings.TrimPrefix(s, prefix)

	parts := strings.SplitN(remainder, "=", 2)
	if len(parts) != 2 {
		return "", "", false
	}

	return parts[0], parts[1], true
}

type MultiFlag map[string]string

func (m *MultiFlag) String() string {
	return fmt.Sprint("%+v", *m)
}

func (m *MultiFlag) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		fmt.Printf("broken -define value: \"%s\"\n", value)
		return fmt.Errorf("broken -define value: \"%s\"\n", value)
	}

	(*m)[parts[0]] = parts[1]
	return nil
}
