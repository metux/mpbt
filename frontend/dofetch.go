// SPDX-License-Identifier: AGPL-3.0-or-later
package frontend

import (
	"fmt"
	"log"

	"github.com/metux/mpbt/core/model"
	"github.com/metux/mpbt/core/workflow/fetch"
)

func doFetch(prj *model.Project) {
	log.Printf("fetching sources ...\n")
	if err := fetch.FetchSource(prj); err != nil {
		panic(fmt.Sprintf("fetch failed: %s\n", err))
	}
}
