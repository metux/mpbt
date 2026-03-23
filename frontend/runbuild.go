// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

func RunBuild(cf BuildConfig) {
	prj := LoadProject(cf)

	doFetch(&prj, false)
	doBuild(&prj)
}
