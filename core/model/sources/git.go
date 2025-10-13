package sources

import (
	"github.com/metux/mpbt/core/util"
)

type Git struct {
	// remote URL - added as `origin`
	Url string

	// reference to checkout
	Ref string

	// fetch depth
	Depth int

	// List of remote refs to fetch
	Fetch util.StringList
}
