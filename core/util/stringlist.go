// SPDX-License-Identifier: AGPL-3.0-or-later
package util

import (
	"gopkg.in/yaml.v3"
)

type StringList []string

func (s *StringList) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		var single string
		if err := value.Decode(&single); err != nil {
			return err
		}
		*s = []string{single}
	case yaml.SequenceNode:
		var list []string
		if err := value.Decode(&list); err != nil {
			return err
		}
		*s = list
	default:
		*s = nil
	}
	return nil
}
