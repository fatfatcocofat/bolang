// Code generated from Bo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Bo
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BoParser.
type BoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BoParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by BoParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by BoParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by BoParser#functionParameters.
	VisitFunctionParameters(ctx *FunctionParametersContext) interface{}

	// Visit a parse tree produced by BoParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by BoParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by BoParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by BoParser#requireStatement.
	VisitRequireStatement(ctx *RequireStatementContext) interface{}

	// Visit a parse tree produced by BoParser#importPath.
	VisitImportPath(ctx *ImportPathContext) interface{}
}
