// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

import (
	"fmt"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/workflow/depgraph"
)

func doDepGraph(prj *model.Project) {
	if err := depgraph.DepGraph(prj); err != nil {
		panic(fmt.Sprintf("depgraph failed: %s\n", err))
	}
}
