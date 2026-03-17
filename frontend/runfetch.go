// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

func RunFetch(cf BuildConfig) {
	prj := LoadProject(cf)
	doFetch(&prj)
}
