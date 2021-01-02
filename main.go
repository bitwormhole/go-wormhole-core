package main

import (
	"fmt"

	"github.com/bitwormhole/gss/application"
	"github.com/bitwormhole/gss/demo"
	"github.com/bitwormhole/gss/io/file"
)

func main() {

	config := &application.Configuration{}
	fs := file.Default()
	roots := fs.Roots()

	demo.RegisterComponents(config)

	context, _ := application.Run(config)
	code := application.Exit(context)
	fmt.Println("exited, code=", code)
	fmt.Println("  file.roots=", roots)

}
