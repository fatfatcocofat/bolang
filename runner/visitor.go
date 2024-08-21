package runner

import (
	"bo/parser"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type BoVisitor struct {
	*parser.BaseBoVisitor
	symbolTable map[string]interface{}
}

func NewBoVisitor() *BoVisitor {
	return &BoVisitor{
		symbolTable: make(map[string]interface{}),
	}
}

func (v *BoVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProgram(ctx)
	case *parser.StatementContext:
		return v.VisitStatement(ctx)
	case *parser.ExpressionContext:
		return v.VisitExpression(ctx)
	case *parser.FunctionCallContext:
		return v.VisitFunctionCall(ctx)
	default:
		panic(fmt.Sprintf("Visit -> unhandled type: %T", ctx))
	}
}

func (v *BoVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, statement := range ctx.AllStatement() {
		v.Visit(statement)
	}

	return nil
}

func (v *BoVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	switch ctx := ctx.GetChild(0).(type) {
	case *parser.RequireStatementContext:
		return v.VisitRequireStatement(ctx)
	case *parser.VariableDeclarationContext:
		varType := ctx.TypeSpec().GetText()
		varName := ctx.ID().GetText()
		varValue := v.Visit(ctx.Expression()) // Evaluate the expression

		switch varType {
		case "int":
			v.symbolTable[varName] = varValue.(int)
		case "float":
			v.symbolTable[varName] = varValue.(float64)
		case "bool":
			v.symbolTable[varName] = varValue.(bool)
		case "string":
			v.symbolTable[varName] = varValue.(string)
		default:
			panic(fmt.Sprintf("VisitStatement -> unhandled variable type: %s", varType))
		}

		return nil
	case *parser.FunctionCallContext:
		return v.VisitFunctionCall(ctx)
	default:
		panic(fmt.Sprintf("VisitStatement -> unhandled statement type: %T", ctx))
	}
}

func (v *BoVisitor) VisitRequireStatement(ctx *parser.RequireStatementContext) interface{} {
	importPath := ctx.ImportPath().GetText()

	fmt.Printf("Importing module: %s\n", importPath)

	return nil
}

func (v *BoVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.INT() != nil {
		val, _ := strconv.Atoi(ctx.INT().GetText())
		return val
	} else if ctx.FLOAT() != nil {
		val, _ := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
		return val
	} else if ctx.STRING() != nil {
		// Remove the quotes from the string
		return ctx.STRING().GetText()[1 : len(ctx.STRING().GetText())-1]
	} else if ctx.BOOL() != nil {
		// Convert the string to a boolean
		return ctx.BOOL().GetText() == "true"
	} else if ctx.ID() != nil {
		// Look up the variable in the symbol table and return its value (if it exists)
		return v.symbolTable[ctx.ID().GetText()]
	} else {
		panic(fmt.Sprintf("VisitExpression -> unhandled expression type: %T", ctx))
	}
}

func (v *BoVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	funcName := ctx.ID().GetText()

	switch funcName {
	case "println":
		for _, arg := range ctx.FunctionParameters().AllExpression() {
			fmt.Println(v.Visit(arg))
		}
	default:
		panic(fmt.Sprintf("VisitFunctionCall -> unhandled function call: %s", funcName))
	}

	return nil
}
