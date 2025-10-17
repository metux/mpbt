package util

import (
	"fmt"
)

type GitRepo struct {
	Dir string
}

func (g GitRepo) IsCheckedOut() bool {
	return ExecRetcode([]string{"git", "rev-parse", "HEAD"}, g.Dir) == 0
}

func (g GitRepo) Init() error {
	return ExecCmd([]string{"git", "init", g.Dir}, "")
}

func (g GitRepo) SetRemoteUrl(remote string, url string) error {
	return ExecCmd([]string{"git", "config", "remote." + remote + ".url", url}, g.Dir)
}

func (g GitRepo) Fetch(depth int, remote string, refs ...string) error {
	c1 := []string{"git", "fetch"}
	if depth > 0 {
		c1 = append(c1, fmt.Sprintf("--depth=%d", depth))
	}

	c1 = append(c1, remote)
	c1 = append(c1, refs...)

	return ExecCmd(c1, g.Dir)
}

func (g GitRepo) SimpleCheckout(refname string) error {
	return ExecCmd([]string{"git", "checkout", refname}, g.Dir)
}
