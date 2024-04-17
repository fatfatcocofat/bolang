# BO-LANG (Bo Programming Language)

> BoLang is a Toy Programming Language implemented in Go.

This project is inspired by the book [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball. The goal of this project is to learn more about programming languages, interpreters, and compilers. BoLang is a toy programming language that is simple and easy to understand. It is not intended to be used in production.

## Features

- [x] Arithmetic operations
- [x] Variables
- [x] Functions
- [x] Control structures (if, else, for, forever)
- [x] Comments
- [x] Error handling
- [x] Coming soon...

## Getting Started

```bash
# Clone the repository
git clone git@github.com:fatfatcocofat/bolang.git
cd bolang

# Build the interpreter
go build

# Run the interpreter
./bo

# Install the interpreter
go install
```

## Examples

```bash
# Run the interpreter
❯ bo # assuming you have the interpreter installed in your PATH environment variable

   _____    _____
  /\  __/\  ) ___ (
  ) )(_ ) )/ /\_/\ \
 / / __/ // /_/ (_\ \
 \ \  _\ \\ \ )_/ / /_
  ) )(__) )\ \/_\/ //_/\
  \/____\/  )_____( \_\/

Welcome to BoLang v0.1.0.
Type !help to see the help message.
>> print("Hello, World!")
Hello, World!
nil
>> let x = 10
nil
>> print(x)
10
nil
>> let add = fn(a, b) { return a + b }
nil
>> print(add(10, 20))
30
nil
>> if (x > 5) { print("x is greater than 5") } else { print("x is less than or equal to 5") }
x is greater than 5
nil
>> ...and so on
```

## Resources

- [Writing An Interpreter In Go](https://interpreterbook.com/): A book by Thorsten Ball that walks through building an interpreter in Go.
- [The Go Programming Language](https://www.gopl.io/): A book by Alan A. A. Donovan and Brian W. Kernighan that covers the Go programming language.
- [Writing A Compiler In Go](https://compilerbook.com/): A book by Thorsten Ball that walks through building a compiler in Go.
- [The Dragon Book](https://en.wikipedia.org/wiki/Compilers:_Principles,_Techniques,_and_Tools): A book by Alfred V. Aho, Monica S. Lam, Ravi Sethi, and Jeffrey D. Ullman that covers compiler design.
- [Crafting Interpreters](https://craftinginterpreters.com/): A book by Bob Nystrom that covers building interpreters and compilers.

## Contributing

Contributions to BoLang are welcome. Please feel free to submit a pull request or open an issue.
