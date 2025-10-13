// SPDX-License-Identifier: AGPL-3.0-or-later
package util

import (
	"fmt"
	"strings"
)

type MultiFlag map[string]string

func (m *MultiFlag) String() string {
	return fmt.Sprint("%+v", *m)
}

func (m *MultiFlag) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("broken -define value: \"%s\"\n", value)
	}

	(*m)[parts[0]] = parts[1]
	return nil
}
