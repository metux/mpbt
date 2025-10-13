package core

import (
    "os"
    "log"
    "gopkg.in/yaml.v2"
)

type Component struct {
    Name string `yaml:"name"`
    Provides string `yaml:"provides"`
    Type string `yaml:"type"`
    Filename string `yaml:"_"`
}

func LoadYaml(fn string, obj interface{}) error {
    yamlFile, err := os.ReadFile(fn)
    if err != nil {
        return err
    }
    err = yaml.Unmarshal(yamlFile, obj)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
        return err
    }
    return nil
}

func LoadComponent(fn string) (Component, error) {
    component := Component{}
    err := LoadYaml(fn, &component)
    component.Filename = fn
    log.Printf("COMPONENT %+v\n", component)
    return component, err
}
