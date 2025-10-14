package core

type Component struct {
	Name     string `yaml:"name"`
	Provides string `yaml:"provides"`
	Type     string `yaml:"type"`
	Filename string `yaml:"_"`
}

type ComponentMap = map[string]*Component

func (c *Component) LoadYaml(fn string) error {
	err := LoadYaml(fn, c)
	c.Filename = fn
	return err
}
