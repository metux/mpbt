// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"flag"
	"fmt"
	"os"
)

func helppage() {
	fmt.Printf("Usage: %s -solution <fn> -root <dir> [command...] [-project-define <name>=<value>...] [-solution-define <name>=<value>...]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Printf("Available commands:\n")
	fmt.Printf("    build           pull sources (once) and run build\n")
	fmt.Printf("    fetch           pull sources only\n")
	fmt.Printf("    depgraph        generate dependency graph (.dot file)\n")
	fmt.Printf("    autopick        run autopick on repos that have it activated\n")
	os.Exit(1)
}
