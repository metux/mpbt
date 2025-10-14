package core

type Component struct {
	Name     string `yaml:"name"`
	Provides StringList `yaml:"provides"`
	Type     string `yaml:"type"`
	Filename string `yaml:"_"`
	BuildDepend StringList `yaml:"build-depends"`
	Depend StringList `yaml:"depends"`
}

type ComponentMap = map[string]*Component

func (c *Component) LoadYaml(fn string) error {
	err := LoadYaml(fn, c)
	c.Filename = fn
	return err
}
