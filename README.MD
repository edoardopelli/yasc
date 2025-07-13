# Swagger JSON to YAML Converter (Go)

A lightweight CLI tool written in Go that converts a Swagger/OpenAPI specification from JSON format to YAML.

## Prerequisites

* **Go** (version 1.16 or higher) installed and available in your `PATH`.
* **Git** (optional, to clone the repository).

## Installation

1. **Clone or download** the repository (if you haven’t already):

   ```bash
   git clone https://github.com/your-username/swagger-json-to-yaml.git
   cd swagger-json-to-yaml
   ```

2. **Fetch the YAML dependency**

   If you are using GOPATH mode (no `go.mod` in this project), run:

   ```bash
   go mod tidy
   ```

   This will install the `yaml.v3` package into your `$GOPATH`.

3. **Build the executable**

   ```bash
   go build -o swagger-convert swagger-convert.go
   ```

   This produces a binary named `swagger-convert` in the current directory.

## Usage

### Convert a file from JSON to YAML

```bash
./swagger-convert -i swagger.json -o swagger.yaml
```

* `-i swagger.json` : Path to the input Swagger JSON file.
* `-o swagger.yaml` : Path for the output Swagger YAML file.

### Using standard input and output

You can also pipe data without specifying files:

```bash
cat swagger.json | ./swagger-convert > swagger.yaml
```

* Omitting `-i` reads JSON from **stdin**.
* Omitting `-o` writes YAML to **stdout**.

## Flags

* `-i string`  : Input file path (default reads from stdin)
* `-o string`  : Output file path (default writes to stdout)

## Example

```bash
# Convert swagger.json to swagger.yaml
./swagger-convert -i api/swagger.json -o api/swagger.yaml

# Or via pipe:
cat api/swagger.json | ./swagger-convert > api/swagger.yaml
```

## Notes

* This tool does **not** rely on Go modules (`go.mod`/`go.sum`). It assumes you have `gopkg.in/yaml.v3` installed in your GOPATH.
* If you prefer modules, initialize one yourself:

  ```bash
  go mod init github.com/your-username/swagger-json-to-yaml
  go get gopkg.in/yaml.v3@v3.0.1
  go build -o swagger-convert
  ```

---

*Happy converting!*

