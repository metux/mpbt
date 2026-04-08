// SPDX-License-Identifier: AGPL-3.0-or-later
package util

import (
	"fmt"
	"log"
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

func (g GitRepo) Fetch(depth int, remote string, force bool, retries int, refs ...string) error {
	var err error

	for x := 0; x < retries; x++ {
		c1 := []string{"git", "fetch"}
		if depth > 0 {
			c1 = append(c1, fmt.Sprintf("--depth=%d", depth))
		}

		if force {
			c1 = append(c1, "--force")
		}

		c1 = append(c1, remote)
		c1 = append(c1, refs...)

		err = ExecCmd(g.Name, c1, g.Dir)
		if err == nil {
			return nil
		}

		log.Printf("attempt %d of %d failed. retrying\n", x, retries)
	}

	return err
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

func (g GitRepo) SimpleCheckout(refname string, localbranch string) error {
	c1 := []string{"git", "checkout", refname}
	if localbranch != "" {
		c1 = append(c1, "-b", localbranch)
	}

	return ExecCmd(g.Name, c1, g.Dir)
}

func (g GitRepo) GetCurrentRev() string {
	return ExecOut([]string{"git", "rev-parse", "HEAD"}, g.Dir)
}
