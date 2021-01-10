package main

import (
	"fmt"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/demo"
	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

func main() {

	config := &application.Configuration{}
	fsys := fs.Default()
	roots := fsys.Roots()

	demo.RegisterComponents(config)

	context, _ := application.Run(config)
	code := application.Exit(context)
	fmt.Println("exited, code=", code)
	fmt.Println("  file.roots=", roots)

}
