package main

import (
	"embed"
	"fmt"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/simple"
	demo "github.com/bitwormhole/go-wormhole-core/demo/demo-for-config"
	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

//go:embed src/main/resources
var resources embed.FS

func main() {

	config := &simple.AppConfig{}
	fsys := fs.Default()
	roots := fsys.Roots()

	config.SetResources(&resources, "src/main/resources")
	demo.Config(config)

	context, _ := application.Run(config)
	code := application.Exit(context)
	fmt.Println("exited, code=", code)
	fmt.Println("  file.roots=", roots)

}
