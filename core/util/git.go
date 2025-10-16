package util

import (
	"fmt"
)

type GitRepo struct {
	Dir string
}

func (g GitRepo) IsCheckedOut() bool {
	return ExecRetcode([]string{"git", "-C", g.Dir, "rev-parse", "HEAD"}) == 0
}

func (g GitRepo) Init() error {
	return ExecCmd([]string{"git", "init", g.Dir})
}

func (g GitRepo) SetRemoteUrl(remote string, url string) error {
	return ExecCmd([]string{"git", "-C", g.Dir, "config", "remote." + remote + ".url", url})
}

func (g GitRepo) Fetch(depth int, remote string, refs... string) error {
	c1 := []string{"git", "-C", g.Dir, "fetch"}
	if depth > 0 {
		c1 = append(c1, fmt.Sprintf("--depth=%d", depth))
	}

	c1 = append(c1, remote)
	c1 = append(c1, refs...)

	return ExecCmd(c1)
}

func (g GitRepo) SimpleCheckout(refname string) error {
	return ExecCmd([]string{"git", "-C", g.Dir, "checkout", refname})
}
