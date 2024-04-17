package lexer

import (
	"bolang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	fn add(x: int, y: int) int {
		return x + y;
	}

	let result = add(five, ten);
	print(result);

	fn hello(name: string) string {
		return "Hello, " + name;
	}

	fn imVoid() void {
		// do nothing
	}

	if (5 < 10 && 10 >= 5 || true) {
		!true;
	} else {
		false;
	}

	5 == 5;
	5 != 10;
	a = [1, 2, 3];

	"foobar";
	"foo bar";

	>= <=;

	1 + 2;
	1 - 2;
	1 * 2;
	1 / 2;
	1 % 2;

	int a;
	float b;
	bool c;
	string d;
	array e;

	a++;
	a--;

	a += 1;
	a -= 1;
	a *= 1;
	a /= 1;
	a %= 1;

	.;
	&!;
	|&;

	1.0;
	1.0 + 2.0;

	for (let i = 0; i < 10; i += 1) {
		print(i);
	}

	forever {
		break;
	}

	continue;
	nil;

	let myMap = {1: 2, 3: 4};

	print("Hello, World!");
	print(1);
	print(1.0);
	print(true);
	print(false);
	print(nil);
	print([1, 2, 3]);
	print({1: 2, 3: 4});
	print("Hello, " + "World!");
	print(1 + 2);
	print(myMap[1]);
	print(myVar);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.T_LET, "let"},
		{token.T_IDENT, "five"},
		{token.T_ASSIGN, "="},
		{token.T_INT, "5"},
		{token.T_SEMICOLON, ";"},
		{token.T_LET, "let"},
		{token.T_IDENT, "ten"},
		{token.T_ASSIGN, "="},
		{token.T_INT, "10"},
		{token.T_SEMICOLON, ";"},
		{token.T_FUNCTION, "fn"},
		{token.T_IDENT, "add"},
		{token.T_LPAREN, "("},
		{token.T_IDENT, "x"},
		{token.T_COLON, ":"},
		{token.T_INT_TYPE, "int"},
		{token.T_COMMA, ","},
		{token.T_IDENT, "y"},
		{token.T_COLON, ":"},
		{token.T_INT_TYPE, "int"},
		{token.T_RPAREN, ")"},
		{token.T_INT_TYPE, "int"},
		{token.T_LBRACE, "{"},
		{token.T_RETURN, "return"},
		{token.T_IDENT, "x"},
		{token.T_PLUS, "+"},
		{token.T_IDENT, "y"},
		{token.T_SEMICOLON, ";"},
		{token.T_RBRACE, "}"},
		{token.T_LET, "let"},
		{token.T_IDENT, "result"},
		{token.T_ASSIGN, "="},
		{token.T_IDENT, "add"},
		{token.T_LPAREN, "("},
		{token.T_IDENT, "five"},
		{token.T_COMMA, ","},
		{token.T_IDENT, "ten"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_IDENT, "result"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_FUNCTION, "fn"},
		{token.T_IDENT, "hello"},
		{token.T_LPAREN, "("},
		{token.T_IDENT, "name"},
		{token.T_COLON, ":"},
		{token.T_STRING_TYPE, "string"},
		{token.T_RPAREN, ")"},
		{token.T_STRING_TYPE, "string"},
		{token.T_LBRACE, "{"},
		{token.T_RETURN, "return"},
		{token.T_STRING, "Hello, "},
		{token.T_PLUS, "+"},
		{token.T_IDENT, "name"},
		{token.T_SEMICOLON, ";"},
		{token.T_RBRACE, "}"},
		{token.T_FUNCTION, "fn"},
		{token.T_IDENT, "imVoid"},
		{token.T_LPAREN, "("},
		{token.T_RPAREN, ")"},
		{token.T_VOID_TYPE, "void"},
		{token.T_LBRACE, "{"},
		{token.T_COMMENT, " do nothing"},
		{token.T_RBRACE, "}"},
		{token.T_IF, "if"},
		{token.T_LPAREN, "("},
		{token.T_INT, "5"},
		{token.T_LT, "<"},
		{token.T_INT, "10"},
		{token.T_AND, "&&"},
		{token.T_INT, "10"},
		{token.T_GTE, ">="},
		{token.T_INT, "5"},
		{token.T_OR, "||"},
		{token.T_TRUE, "true"},
		{token.T_RPAREN, ")"},
		{token.T_LBRACE, "{"},
		{token.T_BANG, "!"},
		{token.T_TRUE, "true"},
		{token.T_SEMICOLON, ";"},
		{token.T_RBRACE, "}"},
		{token.T_ELSE, "else"},
		{token.T_LBRACE, "{"},
		{token.T_FALSE, "false"},
		{token.T_SEMICOLON, ";"},
		{token.T_RBRACE, "}"},
		{token.T_INT, "5"},
		{token.T_EQ, "=="},
		{token.T_INT, "5"},
		{token.T_SEMICOLON, ";"},
		{token.T_INT, "5"},
		{token.T_NOT_EQ, "!="},
		{token.T_INT, "10"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "a"},
		{token.T_ASSIGN, "="},
		{token.T_LBRACKET, "["},
		{token.T_INT, "1"},
		{token.T_COMMA, ","},
		{token.T_INT, "2"},
		{token.T_COMMA, ","},
		{token.T_INT, "3"},
		{token.T_RBRACKET, "]"},
		{token.T_SEMICOLON, ";"},
		{token.T_STRING, "foobar"},
		{token.T_SEMICOLON, ";"},
		{token.T_STRING, "foo bar"},
		{token.T_SEMICOLON, ";"},
		{token.T_GTE, ">="},
		{token.T_LTE, "<="},
		{token.T_SEMICOLON, ";"},
		{token.T_INT, "1"},
		{token.T_PLUS, "+"},
		{token.T_INT, "2"},
		{token.T_SEMICOLON, ";"},
		{token.T_INT, "1"},
		{token.T_MINUS, "-"},
		{token.T_INT, "2"},
		{token.T_SEMICOLON, ";"},
		{token.T_INT, "1"},
		{token.T_ASTERISK, "*"},
		{token.T_INT, "2"},
		{token.T_SEMICOLON, ";"},
		{token.T_INT, "1"},
		{token.T_SLASH, "/"},
		{token.T_INT, "2"},
		{token.T_SEMICOLON, ";"},
		{token.T_INT, "1"},
		{token.T_MOD, "%"},
		{token.T_INT, "2"},
		{token.T_SEMICOLON, ";"},
		{token.T_INT_TYPE, "int"},
		{token.T_IDENT, "a"},
		{token.T_SEMICOLON, ";"},
		{token.T_FLOAT_TYPE, "float"},
		{token.T_IDENT, "b"},
		{token.T_SEMICOLON, ";"},
		{token.T_BOOL_TYPE, "bool"},
		{token.T_IDENT, "c"},
		{token.T_SEMICOLON, ";"},
		{token.T_STRING_TYPE, "string"},
		{token.T_IDENT, "d"},
		{token.T_SEMICOLON, ";"},
		{token.T_ARRAY_TYPE, "array"},
		{token.T_IDENT, "e"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "a"},
		{token.T_INC, "++"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "a"},
		{token.T_DEC, "--"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "a"},
		{token.T_ADD_EQ, "+="},
		{token.T_INT, "1"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "a"},
		{token.T_SUB_EQ, "-="},
		{token.T_INT, "1"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "a"},
		{token.T_MUL_EQ, "*="},
		{token.T_INT, "1"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "a"},
		{token.T_DIV_EQ, "/="},
		{token.T_INT, "1"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "a"},
		{token.T_MOD_EQ, "%="},
		{token.T_INT, "1"},
		{token.T_SEMICOLON, ";"},
		{token.T_DOT, "."},
		{token.T_SEMICOLON, ";"},
		{token.T_ILLEGAL, "&!"},
		{token.T_SEMICOLON, ";"},
		{token.T_ILLEGAL, "|&"},
		{token.T_SEMICOLON, ";"},
		{token.T_FLOAT, "1.0"},
		{token.T_SEMICOLON, ";"},
		{token.T_FLOAT, "1.0"},
		{token.T_PLUS, "+"},
		{token.T_FLOAT, "2.0"},
		{token.T_SEMICOLON, ";"},
		{token.T_FOR, "for"},
		{token.T_LPAREN, "("},
		{token.T_LET, "let"},
		{token.T_IDENT, "i"},
		{token.T_ASSIGN, "="},
		{token.T_INT, "0"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "i"},
		{token.T_LT, "<"},
		{token.T_INT, "10"},
		{token.T_SEMICOLON, ";"},
		{token.T_IDENT, "i"},
		{token.T_ADD_EQ, "+="},
		{token.T_INT, "1"},
		{token.T_RPAREN, ")"},
		{token.T_LBRACE, "{"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_IDENT, "i"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_RBRACE, "}"},
		{token.T_FOREVER, "forever"},
		{token.T_LBRACE, "{"},
		{token.T_BREAK, "break"},
		{token.T_SEMICOLON, ";"},
		{token.T_RBRACE, "}"},
		{token.T_CONTINUE, "continue"},
		{token.T_SEMICOLON, ";"},
		{token.T_NIL, "nil"},
		{token.T_SEMICOLON, ";"},
		{token.T_LET, "let"},
		{token.T_IDENT, "myMap"},
		{token.T_ASSIGN, "="},
		{token.T_LBRACE, "{"},
		{token.T_INT, "1"},
		{token.T_COLON, ":"},
		{token.T_INT, "2"},
		{token.T_COMMA, ","},
		{token.T_INT, "3"},
		{token.T_COLON, ":"},
		{token.T_INT, "4"},
		{token.T_RBRACE, "}"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_STRING, "Hello, World!"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_INT, "1"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_FLOAT, "1.0"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_TRUE, "true"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_FALSE, "false"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_NIL, "nil"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_LBRACKET, "["},
		{token.T_INT, "1"},
		{token.T_COMMA, ","},
		{token.T_INT, "2"},
		{token.T_COMMA, ","},
		{token.T_INT, "3"},
		{token.T_RBRACKET, "]"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_LBRACE, "{"},
		{token.T_INT, "1"},
		{token.T_COLON, ":"},
		{token.T_INT, "2"},
		{token.T_COMMA, ","},
		{token.T_INT, "3"},
		{token.T_COLON, ":"},
		{token.T_INT, "4"},
		{token.T_RBRACE, "}"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_STRING, "Hello, "},
		{token.T_PLUS, "+"},
		{token.T_STRING, "World!"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_INT, "1"},
		{token.T_PLUS, "+"},
		{token.T_INT, "2"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_IDENT, "myMap"},
		{token.T_LBRACKET, "["},
		{token.T_INT, "1"},
		{token.T_RBRACKET, "]"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_PRINT, "print"},
		{token.T_LPAREN, "("},
		{token.T_IDENT, "myVar"},
		{token.T_RPAREN, ")"},
		{token.T_SEMICOLON, ";"},
		{token.T_EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
