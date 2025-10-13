package main

import (
	"log"

	"github.com/metux/mpbt/core"
)

func main() {
	list := make([]core.Component, 0)
	err := core.LoadComponents("../cf/xlibre/components", &list)
	if err != nil {
		log.Fatalf("error opening components directory: %s\n", err)
	}

	for _,c := range list {
		log.Printf("Component: %+v\n", c)
	}
}
