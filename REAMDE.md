# BO-LANG (Bo Programming Language)

BoLang is a programming language built in Go. It's a simple language designed for learning purposes.

## Project Structure

The project is organized into several packages:

- [`ast`](command:_github.copilot.openSymbolInFile?%5B%22ast%2Fast.go%22%2C%22ast%22%5D "ast/ast.go"): Contains the abstract syntax tree definitions for the language.
- [`lexer`](command:_github.copilot.openSymbolInFile?%5B%22lexer%2Flexer.go%22%2C%22lexer%22%5D "lexer/lexer.go"): Contains the lexer which converts source code into tokens.
- [`parser`](command:_github.copilot.openSymbolInFile?%5B%22parser%2Fparser.go%22%2C%22parser%22%5D "parser/parser.go"): Contains the parser which converts tokens into an abstract syntax tree.
- [`repl`](command:_github.copilot.openSymbolInFile?%5B%22repl%2Frepl.go%22%2C%22repl%22%5D "repl/repl.go"): Contains the Read-Eval-Print Loop (REPL) for interactive use of the language.
- [`token`](command:_github.copilot.openSymbolInFile?%5B%22token%2Ftoken.go%22%2C%22token%22%5D "token/token.go"): Contains the definitions of the language's tokens.

## Running the REPL

You can start the REPL by running the [`main.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Ffathurrohman%2FCode%2Fbolang%2Fmain.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/fathurrohman/Code/bolang/main.go") file. Once the REPL is running, you can type in BoLang code and see the resulting abstract syntax tree.

The REPL supports a few commands:

- `!help`: Prints a help message.
- `!exit`: Exits the REPL.

## Error Handling

If there are any errors in parsing the code, the REPL will print an error message.

## Contributing

Contributions to BoLang are welcome. Please feel free to submit a pull request or open an issue.
