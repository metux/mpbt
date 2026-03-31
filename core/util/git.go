// SPDX-License-Identifier: AGPL-3.0-or-later
package util

import (
	"fmt"
)

type GitRepo struct {
	Dir  string
	Name string
}

func (g GitRepo) IsCheckedOut() bool {
	return ExecRetcode([]string{"git", "rev-parse", "HEAD"}, g.Dir) == 0
}

func (g GitRepo) Init() error {
	return ExecCmd(g.Name, []string{"git", "init", g.Dir}, "")
}

func (g GitRepo) SetRemoteUrl(remote string, url string) error {
	return ExecCmd(g.Name, []string{"git", "config", "remote." + remote + ".url", url}, g.Dir)
}

func (g GitRepo) Fetch(depth int, remote string, force bool, refs ...string) error {
	c1 := []string{"git", "fetch"}
	if depth > 0 {
		c1 = append(c1, fmt.Sprintf("--depth=%d", depth))
	}

	if force {
		c1 = append(c1, "--force")
	}

	c1 = append(c1, remote)
	c1 = append(c1, refs...)

	return ExecCmd(g.Name, c1, g.Dir)
}

func (g GitRepo) ConfigFetch(remote string, refs ...string) error {
	for _, r := range refs {
		c1 := []string{"git", "config", "--add", "remote." + remote + ".fetch", "+" + r}
		ExecCmd(g.Name, c1, g.Dir)
	}
	return nil
}

func (g GitRepo) ConfigSet(name string, value string) error {
	return ExecCmd(g.Name, []string{"git", "config", name, value}, g.Dir)
}

func (g GitRepo) SimpleCheckout(refname string) error {
	return ExecCmd(g.Name, []string{"git", "checkout", refname}, g.Dir)
}

func (g GitRepo) GetCurrentRev() string {
	return ExecOut([]string{"git", "rev-parse", "HEAD"}, g.Dir)
}
