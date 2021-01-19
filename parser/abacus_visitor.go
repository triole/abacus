// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by AbacusParser.
type AbacusVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AbacusParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by AbacusParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by AbacusParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by AbacusParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by AbacusParser#Pow.
	VisitPow(ctx *PowContext) interface{}

	// Visit a parse tree produced by AbacusParser#AtomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by AbacusParser#Parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by AbacusParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by AbacusParser#Variable.
	VisitVariable(ctx *VariableContext) interface{}
}
