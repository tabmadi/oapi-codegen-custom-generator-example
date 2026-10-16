// Package generator provides functionality for generating code from OpenAPI specifications.
package generator

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"text/template"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/util"

	"github.com/tabmadi/oapi-codegen-custom-generator-example/internal/config"
)

const (
	readOnly  = "read only"
	writeOnly = "write only"
)

// Generator handles the code generation process based on OpenAPI specification files.
type Generator struct {
	filePath string
}

// NewGenerator creates a new Generator instance with the provided configuration.
func NewGenerator(cfg *config.Config) *Generator {
	return &Generator{
		filePath: cfg.FilePath,
	}
}

// OperationData holds data for template generation.
type OperationData struct {
	OperationID string
	Method      string
	Path        string
	LogMessage  string
}

// Generate loads the OpenAPI specification and performs the code generation process.
func (g *Generator) Generate() {
	swagger, err := util.LoadSwagger(g.filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load spec: %v", err))
	}

	tmpl, err := template.ParseFiles("templates/functions.go.tmpl")
	if err != nil {
		panic(fmt.Sprintf("Failed to parse template file: %v", err))
	}

	var operations []OperationData

	paths := swagger.Paths.Map()

	// Map order is random, so the paths are sorted to keep the output stable between runs.
	for _, path := range slices.Sorted(maps.Keys(paths)) {
		pathItem := paths[path]

		if pathItem.Get != nil {
			operations = append(operations, OperationData{
				OperationID: pathItem.Get.OperationID,
				Method:      "GET",
				Path:        path,
				LogMessage:  readOnly,
			})
		}

		if pathItem.Post != nil {
			operations = append(operations, OperationData{
				OperationID: pathItem.Post.OperationID,
				Method:      "POST",
				Path:        path,
				LogMessage:  writeOnly,
			})
		}

		if pathItem.Put != nil {
			operations = append(operations, OperationData{
				OperationID: pathItem.Put.OperationID,
				Method:      "PUT",
				Path:        path,
				LogMessage:  writeOnly,
			})
		}

		if pathItem.Delete != nil {
			operations = append(operations, OperationData{
				OperationID: pathItem.Delete.OperationID,
				Method:      "DELETE",
				Path:        path,
				LogMessage:  writeOnly,
			})
		}
	}

	err = tmpl.Execute(os.Stdout, operations)
	if err != nil {
		panic(fmt.Sprintf("Failed to execute template: %v", err))
	}
}
