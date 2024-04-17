package ast

import (
	"bolang/token"
	"bytes"
	"strings"
)

// Node is the interface that all nodes in the AST implement.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is the interface that all statement nodes in the AST implement.
type Statement interface {
	Node
	statementNode()
}

// Expression is the interface that all expression nodes in the AST implement.
type Expression interface {
	Node
	expressionNode()
}

// Program represents the root node of the AST.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal value of the token associated with the node.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns a string representation of the program.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement represents a let statement in the AST.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral returns the literal value of the token associated with the node.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String returns a string representation of the let statement.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) statementNode() {}

// Identifier represents an identifier in the AST.
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral returns the literal value of the token associated with the node.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String returns a string representation of the identifier.
func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

// ReturnStatement represents a return statement in the AST.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral returns the literal value of the token associated with the node.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns a string representation of the return statement.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement represents an expression statement in the AST.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// TokenLiteral returns the literal value of the token associated with the node.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns a string representation of the expression statement.
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) statementNode() {}

// IntegerLiteral represents an integer literal in the AST.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral returns the literal value of the token associated with the node.
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String returns a string representation of the integer literal.
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) expressionNode() {}

// StringLiteral represents a string literal in the AST.
type StringLiteral struct {
	Token token.Token
	Value string
}

// TokenLiteral returns the literal value of the token associated with the node.
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

// String returns a string representation of the string literal.
func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) expressionNode() {}

// FloatLiteral represents a float literal in the AST.
type FloatLiteral struct {
	Token token.Token
	Value float64
}

// TokenLiteral returns the literal value of the token associated with the node.
func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String returns a string representation of the float literal.
func (fl *FloatLiteral) String() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) expressionNode() {}

// Boolean represents a boolean in the AST.
type Boolean struct {
	Token token.Token
	Value bool
}

// TokenLiteral returns the literal value of the token associated with the node.
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// String returns a string representation of the boolean.
func (b *Boolean) String() string {
	return b.Token.Literal
}

func (b *Boolean) expressionNode() {}

// PrefixExpression represents a prefix expression in the AST.
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// TokenLiteral returns the literal value of the token associated with the node.
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns a string representation of the prefix expression.
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {}

// InfixExpression represents an infix expression in the AST.
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// TokenLiteral returns the literal value of the token associated with the node.
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns a string representation of the infix expression.
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *InfixExpression) expressionNode() {}

// IfExpression represents an if expression in the AST.
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// TokenLiteral returns the literal value of the token associated with the node.
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns a string representation of the if expression.
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

func (ie *IfExpression) expressionNode() {}

// BlockStatement represents a block statement in the AST.
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

// TokenLiteral returns the literal value of the token associated with the node.
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String returns a string representation of the block statement.
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (bs *BlockStatement) statementNode() {}

// FunctionLiteral represents a function literal in the AST.
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

// TokenLiteral returns the literal value of the token associated with the node.
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String returns a string representation of the function literal.
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", ")) // Join the elements of params with a comma separator
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (fl *FunctionLiteral) expressionNode() {}

// CallExpression represents a call expression in the AST.
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

// TokenLiteral returns the literal value of the token associated with the node.
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String returns a string representation of the call expression.
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (ce *CallExpression) expressionNode() {}

// ArrayLiteral represents an array literal in the AST.
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

// TokenLiteral returns the literal value of the token associated with the node.
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

// String returns a string representation of the array literal.
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range al.Elements {
		elements = append(elements, e.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (al *ArrayLiteral) expressionNode() {}

// IndexExpression represents an index expression in the AST.
type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

// TokenLiteral returns the literal value of the token associated with the node.
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns a string representation of the index expression.
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

func (ie *IndexExpression) expressionNode() {}
