// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

func RunDepGraph(cf BuildConfig) {
	prj := LoadProject(cf)
	doDepGraph(&prj)
}
