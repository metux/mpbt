// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

import (
	"fmt"

	"github.com/metux/mpbt/core/workflow/autopick"
)

func RunAutoPick(cf BuildConfig) {
	prj := LoadProject(cf)

	if err := autopick.AutoPick(&prj); err != nil {
		panic(fmt.Sprintf("autopick failed: %s\n", err))
	}
}
