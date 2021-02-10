package main

import (
	"fmt"

	"github.com/bitwormhole/go-wormhole-core/application"
	demo "github.com/bitwormhole/go-wormhole-core/demo/demo-for-config"
	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

func main() {

	config := &application.SimpleConfiguration{}
	fsys := fs.Default()
	roots := fsys.Roots()

	demo.Config(config)

	context, _ := application.Run(config)
	code := application.Exit(context)
	fmt.Println("exited, code=", code)
	fmt.Println("  file.roots=", roots)

}
