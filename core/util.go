package core

import (
    "os"
    "log"
    "gopkg.in/yaml.v3"
)

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
