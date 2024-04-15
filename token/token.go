package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	T_ILLEGAL = "ILLEGAL"
	T_EOF     = "EOF"
	T_COMMENT = "COMMENT"

	T_IDENT  = "IDENT"
	T_INT    = "INT"
	T_FLOAT  = "FLOAT"
	T_BOOL   = "BOOL"
	T_STRING = "STRING"

	T_ASSIGN   = "="
	T_PLUS     = "+"
	T_MINUS    = "-"
	T_BANG     = "!"
	T_ASTERISK = "*"
	T_SLASH    = "/"
	T_MOD      = "%"
	T_LT       = "<"
	T_GT       = ">"

	T_LTE    = "<="
	T_GTE    = ">="
	T_EQ     = "=="
	T_NOT_EQ = "!="
	T_ADD_EQ = "+="
	T_SUB_EQ = "-="
	T_MUL_EQ = "*="
	T_DIV_EQ = "/="
	T_MOD_EQ = "%="

	T_INC = "++"
	T_DEC = "--"
	T_AND = "&&"
	T_OR  = "||"

	T_COMMA     = ","
	T_COLON     = ":"
	T_SEMICOLON = ";"
	T_DOT       = "."

	T_LPAREN   = "("
	T_RPAREN   = ")"
	T_LBRACE   = "{"
	T_RBRACE   = "}"
	T_LBRACKET = "["
	T_RBRACKET = "]"

	T_FUNCTION    = "FUNCTION"
	T_LET         = "LET"
	T_NIL         = "NIL"
	T_TRUE        = "TRUE"
	T_FALSE       = "FALSE"
	T_IF          = "IF"
	T_ELSE        = "ELSE"
	T_RETURN      = "RETURN"
	T_PRINT       = "PRINT"
	T_FOR         = "FOR"
	T_FOREVER     = "FOREVER"
	T_BREAK       = "BREAK"
	T_CONTINUE    = "CONTINUE"
	T_INT_TYPE    = "INT_TYPE"
	T_FLOAT_TYPE  = "FLOAT_TYPE"
	T_BOOL_TYPE   = "BOOL_TYPE"
	T_STRING_TYPE = "STRING_TYPE"
	T_VOID_TYPE   = "VOID_TYPE"
	T_ARRAY_TYPE  = "ARRAY_TYPE"
)

var keywords = map[string]TokenType{
	"fn":       T_FUNCTION,
	"let":      T_LET,
	"nil":      T_NIL,
	"true":     T_TRUE,
	"false":    T_FALSE,
	"if":       T_IF,
	"else":     T_ELSE,
	"return":   T_RETURN,
	"print":    T_PRINT,
	"for":      T_FOR,
	"forever":  T_FOREVER,
	"break":    T_BREAK,
	"continue": T_CONTINUE,
	"int":      T_INT_TYPE,
	"float":    T_FLOAT_TYPE,
	"bool":     T_BOOL_TYPE,
	"string":   T_STRING_TYPE,
	"void":     T_VOID_TYPE,
	"array":    T_ARRAY_TYPE,
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Literal: %s}", t.Type, t.Literal)
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return T_IDENT
}

func LookupType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return T_IDENT
}
