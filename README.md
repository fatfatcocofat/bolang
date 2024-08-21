# Bo (Bo Lang) - The Bo Programming Language (WIP)

> DISCLAIMER: This is DIY project and just for fun. It is not intended to be used in production.

Bo is a simple programming language that I am building for fun and learning purposes. It is a statically typed language with a syntax similar to Ruby and Go. The language is still in its early stages of development and is not yet ready for use. Bo is written in Go and uses the `antlr4` library for parsing.

## Example

```go
// In Bo, `main` is not required

// Importing modules (not yet implemented)
require <bo/fmt> // import standard library
require "path/to/bo-bo.bo" // import local file

// Variables
int x = 10
float y = 10.5
string name = "Bo"
bool isTrue = true

// Built-in functions
println("Hello, World!")
println("x:", x, "y:", y)
```

## Development

Bo is still in its early stages of development. The language is not yet ready for use. If you are interested in contributing, feel free to open an issue or submit a pull request :)

#### Generate parser

```bash
sh genparser.sh
```

## License

[MIT LICENSE](LICENSE)
