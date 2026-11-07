// Thin entry: cobra commands live in internal/cli.
package main

import (
	"os"

	"github.com/itsryanarmani/clis/search-console/internal/cli"
)

func main() {
	os.Exit(cli.Main(os.Args[1:], os.Stdout, os.Stderr, os.Getenv))
}
