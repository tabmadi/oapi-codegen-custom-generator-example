# oapi-codegen Custom Generator Example

A small Go program that reads an OpenAPI spec with the [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) loader and writes Go code from a `text/template` of your own. It shows how to generate code that the built-in oapi-codegen outputs do not cover.

## What it shows

oapi-codegen generates models, servers, and clients. A project often needs more from the same spec: route tables, permission maps, mocks, or docs. This example uses the oapi-codegen loader and writes its own output in four steps:

1. **Load**: `util.LoadSwagger` parses `api/openapi.yml` into the kin-openapi model.
2. **Extract**: the generator walks the paths in sorted order and collects each `GET`, `POST`, `PUT`, and `DELETE` operation.
3. **Render**: `templates/functions.go.tmpl` turns each operation into a Go function.
4. **Print**: the result goes to standard output.

For the sample spec, the output starts like this:

```go
// listUsers handles GET /users
func listUsers() {
	log.Println("read only")
}

// createUser handles POST /users
func createUser() {
	log.Println("write only")
}
```

## Run it

1. Install [mise](https://mise.jdx.dev/getting-started.html).
2. Install the tools, then run the generator:

   ```bash
   git clone git@github.com:tabmadi/oapi-codegen-custom-generator-example.git
   cd oapi-codegen-custom-generator-example
   mise install
   mise run setup
   mise run generate
   ```

3. To read another spec, set `FILE_PATH`:

   ```bash
   FILE_PATH=path/to/openapi.yml mise run generate > generated.go
   ```

| Variable | Description | Default |
| --- | --- | --- |
| `FILE_PATH` | The OpenAPI spec that the generator reads | `./api/openapi.yml` |

## How it works

| File | Holds |
| --- | --- |
| [`cmd/generate/main.go`](cmd/generate/main.go) | Loads the config and runs the generator |
| [`internal/generator/generator.go`](internal/generator/generator.go) | Loads the spec, collects the operations, and renders the template |
| [`templates/functions.go.tmpl`](templates/functions.go.tmpl) | The output shape. Change it to change what the generator writes |
| [`api/openapi.yml`](api/openapi.yml) | The sample spec: a users CRUD API |
| [`internal/config/config.go`](internal/config/config.go) | Reads `FILE_PATH` from the environment and `.env` |

To generate something else, change the template first. To give the template more data, add a field to `OperationData` and fill it in the loop. The kin-openapi model also holds the parameters, the request bodies, the responses, and any `x-` extension.

## What is deliberately missing

- **Every HTTP method**: the loop reads `GET`, `POST`, `PUT`, and `DELETE`. `PATCH`, `HEAD`, and `OPTIONS` follow the same pattern.
- **Formatted output**: the result is not passed through `go/format`.
- **A file output**: the generator prints to standard output. Redirect it to a file, or call it from a `//go:generate` line.
- **Tests**: a golden-file test that compares the output with a committed file is the usual pattern.

## Tasks

| Task | Does |
| --- | --- |
| `mise run setup` | Downloads the modules and installs the git hooks |
| `mise run generate` | Runs the generator and prints the code |
| `mise run build` | Builds a static binary at `bin/generate` |
| `mise run lint` | Runs `lint:go`, `lint:md`, and `lint:secrets` |
| `mise run format` | Applies every formatter |
| `mise run check` | Runs `lint` and `test`. Run it before you finish a change |
| `mise run ci` | Runs exactly what CI runs. The `pre-push` hook runs it too |

## License

Apache License 2.0. See [LICENSE](LICENSE).
