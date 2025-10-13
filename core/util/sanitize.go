// SPDX-License-Identifier: AGPL-3.0-or-later
package util

import (
	"regexp"
	"strings"
)

func SanitizeFilename(name string) string {
	// Replace '/' and any NUL byte (\x00)
	invalid := regexp.MustCompile(`[/\x00]`)
	sanitized := invalid.ReplaceAllString(name, "_")

	// Trim whitespace at start/end (optional, stylistic)
	sanitized = strings.TrimSpace(sanitized)

	// Avoid empty names
	if sanitized == "" {
		sanitized = "_"
	}

	return sanitized
}
