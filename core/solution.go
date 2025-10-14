package core

type Solution struct {
	ComponentMapping map[string]string `yaml:"component-mapping"`
	Filename string `yaml:"_"`
	Build StringList `yaml:"build"`
}

func (c *Solution) LoadYaml(fn string) error {
	err := LoadYaml(fn, c)
	c.Filename = fn
	return err
}

func (c *Solution) GetMapped(name string) string {
	if c.ComponentMapping == nil {
		return name
	}
	if val, ok := c.ComponentMapping[name]; ok {
		return val
	}
	return name
}
