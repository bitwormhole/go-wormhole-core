package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/config"

	demo "github.com/bitwormhole/go-wormhole-core/demo/demo-for-config"
)

//go:embed src/main/resources
var resources embed.FS

func tryMain() error {

	config := &config.AppConfig{}
	// fsys := fs.Default()
	// roots := fsys.Roots()
	args := os.Args

	config.SetResources(&resources, "src/main/resources")
	demo.Config(config)

	context, err := application.Run(config, args)
	if err != nil {
		return err
	}

	code := application.Exit(context)
	fmt.Println("exited, code=", code)
	// fmt.Println("  file.roots=", roots)

	return nil
}

func main() {
	err := tryMain()
	if err != nil {
		panic(err)
	}
}
