# Bo (Bo Lang) - The Bo Programming Language (WIP)

> IT IS USED FOR EDUCATIONAL PURPOSES AND TO DEMONSTRATE THEORETICAL CONCEPTS.

What Does Toy Language Mean?
Toy language refers to any computer programming language that is not considered to be suitable or capable for building general purpose and high-end software and applications. It can be any programming language that lacks the advanced features, capabilities, programming constructs and paradigms of high level language. Toy language may also be termed esoteric programming language.

Techopedia Explains Toy Language
Toy language was primarily created as a means of programming language research and education, proof of concept for a computer science or programming theory and to create a prototype for a new programming language. Typically, toy language has all the capabilities to perform simple to complex mathematical and programming computations. However, it has an incapability in terms of lesser or no library programs support, missing programming constructs such as pointers and arrays, which limits it in creating general-use programs and applications. Pascal, Treelang and Logo are popular examples of toy language.

> Quote from [Techopedia](https://www.techopedia.com/definition/22609/toy-language)

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

## Resources

- [Writing An Interpreter In Go](https://interpreterbook.com/): A book by Thorsten Ball that walks through building an interpreter in Go.
- [The Go Programming Language](https://www.gopl.io/): A book by Alan A. A. Donovan and Brian W. Kernighan that covers the Go programming language.
- [Writing A Compiler In Go](https://compilerbook.com/): A book by Thorsten Ball that walks through building a compiler in Go.
- [The Dragon Book](https://en.wikipedia.org/wiki/Compilers:_Principles,_Techniques,_and_Tools): A book by Alfred V. Aho, Monica S. Lam, Ravi Sethi, and Jeffrey D. Ullman that covers compiler design.
- [Crafting Interpreters](https://craftinginterpreters.com/): A book by Bob Nystrom that covers building interpreters and compilers.
