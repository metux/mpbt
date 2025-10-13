package main

import (
	"log"

	"github.com/metux/mpbt/core"
)

func main() {
	db := core.ComponentsDB{}

	err := db.LoadComponents("../cf/xlibre/components")
	if err != nil {
		log.Fatalf("error opening components directory: %s\n", err)
	}

	for n,c := range db.Components {
		log.Printf("Component: %s => %+v\n", n, c)
	}

	for n,c := range db.Provides {
		log.Printf("Provides: %s => %+v\n", n, c)
	}

	sol := core.Solution{}
	sol.LoadYaml("../cf/xlibre/solutions/devuan.yaml")

	log.Printf("Solution: %+v\n", sol)
}
