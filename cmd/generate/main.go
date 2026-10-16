// Package main is the entry point for the application.
package main

import (
	"github.com/tabmadi/oapi-codegen-custom-generator-example/internal/config"
	"github.com/tabmadi/oapi-codegen-custom-generator-example/internal/generator"
)

func main() {
	cfg := config.Load()

	gen := generator.NewGenerator(cfg)
	gen.Generate()
}
