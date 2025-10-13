// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

type BuildConfig struct {
	SolutionFile    string
	RootDir         string
	WorkDir         string
	SolutionDefines map[string]string
	ProjectDefines  map[string]string
}
