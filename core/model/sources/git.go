// SPDX-License-Identifier: AGPL-3.0-or-later
package sources

import (
	"github.com/metux/go-magicdict/api"
	"github.com/metux/mpbt/core/util"
)

type GitRemote struct {
	Name string

	// remote URL
	Url string

	// fetch depth
	Depth int

	// List of remote refs to fetch
	Fetch util.StringList
}

type Git struct {
	Remotes map[string]GitRemote

	Ref         string
	LocalBranch string

	PostCheckoutCmd util.StringList
	Config          map[api.Key]string
}

func LoadGitRemote(ent api.Entry, name string) GitRemote {
	return GitRemote{
		Name:  name,
		Url:   api.GetStr(ent, "url"),
		Depth: api.GetInt(ent, "depth", 0),
		Fetch: api.GetStrList(ent, "fetch"),
	}
}
