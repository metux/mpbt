// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/workflow/build"
)

func doBuild(prj *model.Project) {
	log.Printf("building ...\n")
	if err := build.Build(prj); err != nil {
		panic(fmt.Sprintf("build failed: %s\n", err))
	}
}
