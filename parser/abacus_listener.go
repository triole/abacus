// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AbacusListener is a complete listener for a parse tree produced by AbacusParser.
type AbacusListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterEqualComparison is called when entering the EqualComparison production.
	EnterEqualComparison(c *EqualComparisonContext)

	// EnterLessComparison is called when entering the LessComparison production.
	EnterLessComparison(c *LessComparisonContext)

	// EnterGreaterComparison is called when entering the GreaterComparison production.
	EnterGreaterComparison(c *GreaterComparisonContext)

	// EnterLessOrEqualComparison is called when entering the LessOrEqualComparison production.
	EnterLessOrEqualComparison(c *LessOrEqualComparisonContext)

	// EnterGreaterOrEqualComparison is called when entering the GreaterOrEqualComparison production.
	EnterGreaterOrEqualComparison(c *GreaterOrEqualComparisonContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterPow is called when entering the Pow production.
	EnterPow(c *PowContext)

	// EnterAtomExpr is called when entering the AtomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterParentheses is called when entering the Parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterFuncExpr is called when entering the FuncExpr production.
	EnterFuncExpr(c *FuncExprContext)

	// EnterConstant is called when entering the Constant production.
	EnterConstant(c *ConstantContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterPlusSign is called when entering the PlusSign production.
	EnterPlusSign(c *PlusSignContext)

	// EnterMinusSign is called when entering the MinusSign production.
	EnterMinusSign(c *MinusSignContext)

	// EnterSqrtFunction is called when entering the SqrtFunction production.
	EnterSqrtFunction(c *SqrtFunctionContext)

	// EnterLnFunction is called when entering the LnFunction production.
	EnterLnFunction(c *LnFunctionContext)

	// EnterLogDefFunction is called when entering the LogDefFunction production.
	EnterLogDefFunction(c *LogDefFunctionContext)

	// EnterLog2Function is called when entering the Log2Function production.
	EnterLog2Function(c *Log2FunctionContext)

	// EnterLog10Function is called when entering the Log10Function production.
	EnterLog10Function(c *Log10FunctionContext)

	// EnterFloorFunction is called when entering the FloorFunction production.
	EnterFloorFunction(c *FloorFunctionContext)

	// EnterCeilFunction is called when entering the CeilFunction production.
	EnterCeilFunction(c *CeilFunctionContext)

	// EnterExpFunction is called when entering the ExpFunction production.
	EnterExpFunction(c *ExpFunctionContext)

	// EnterSinFunction is called when entering the SinFunction production.
	EnterSinFunction(c *SinFunctionContext)

	// EnterCosFunction is called when entering the CosFunction production.
	EnterCosFunction(c *CosFunctionContext)

	// EnterTanFunction is called when entering the TanFunction production.
	EnterTanFunction(c *TanFunctionContext)

	// EnterRoundDefFunction is called when entering the RoundDefFunction production.
	EnterRoundDefFunction(c *RoundDefFunctionContext)

	// EnterRound2Function is called when entering the Round2Function production.
	EnterRound2Function(c *Round2FunctionContext)

	// EnterLogFunction is called when entering the LogFunction production.
	EnterLogFunction(c *LogFunctionContext)

	// EnterMinFunction is called when entering the MinFunction production.
	EnterMinFunction(c *MinFunctionContext)

	// EnterMaxFunction is called when entering the MaxFunction production.
	EnterMaxFunction(c *MaxFunctionContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitEqualComparison is called when exiting the EqualComparison production.
	ExitEqualComparison(c *EqualComparisonContext)

	// ExitLessComparison is called when exiting the LessComparison production.
	ExitLessComparison(c *LessComparisonContext)

	// ExitGreaterComparison is called when exiting the GreaterComparison production.
	ExitGreaterComparison(c *GreaterComparisonContext)

	// ExitLessOrEqualComparison is called when exiting the LessOrEqualComparison production.
	ExitLessOrEqualComparison(c *LessOrEqualComparisonContext)

	// ExitGreaterOrEqualComparison is called when exiting the GreaterOrEqualComparison production.
	ExitGreaterOrEqualComparison(c *GreaterOrEqualComparisonContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitPow is called when exiting the Pow production.
	ExitPow(c *PowContext)

	// ExitAtomExpr is called when exiting the AtomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitParentheses is called when exiting the Parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitFuncExpr is called when exiting the FuncExpr production.
	ExitFuncExpr(c *FuncExprContext)

	// ExitConstant is called when exiting the Constant production.
	ExitConstant(c *ConstantContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitPlusSign is called when exiting the PlusSign production.
	ExitPlusSign(c *PlusSignContext)

	// ExitMinusSign is called when exiting the MinusSign production.
	ExitMinusSign(c *MinusSignContext)

	// ExitSqrtFunction is called when exiting the SqrtFunction production.
	ExitSqrtFunction(c *SqrtFunctionContext)

	// ExitLnFunction is called when exiting the LnFunction production.
	ExitLnFunction(c *LnFunctionContext)

	// ExitLogDefFunction is called when exiting the LogDefFunction production.
	ExitLogDefFunction(c *LogDefFunctionContext)

	// ExitLog2Function is called when exiting the Log2Function production.
	ExitLog2Function(c *Log2FunctionContext)

	// ExitLog10Function is called when exiting the Log10Function production.
	ExitLog10Function(c *Log10FunctionContext)

	// ExitFloorFunction is called when exiting the FloorFunction production.
	ExitFloorFunction(c *FloorFunctionContext)

	// ExitCeilFunction is called when exiting the CeilFunction production.
	ExitCeilFunction(c *CeilFunctionContext)

	// ExitExpFunction is called when exiting the ExpFunction production.
	ExitExpFunction(c *ExpFunctionContext)

	// ExitSinFunction is called when exiting the SinFunction production.
	ExitSinFunction(c *SinFunctionContext)

	// ExitCosFunction is called when exiting the CosFunction production.
	ExitCosFunction(c *CosFunctionContext)

	// ExitTanFunction is called when exiting the TanFunction production.
	ExitTanFunction(c *TanFunctionContext)

	// ExitRoundDefFunction is called when exiting the RoundDefFunction production.
	ExitRoundDefFunction(c *RoundDefFunctionContext)

	// ExitRound2Function is called when exiting the Round2Function production.
	ExitRound2Function(c *Round2FunctionContext)

	// ExitLogFunction is called when exiting the LogFunction production.
	ExitLogFunction(c *LogFunctionContext)

	// ExitMinFunction is called when exiting the MinFunction production.
	ExitMinFunction(c *MinFunctionContext)

	// ExitMaxFunction is called when exiting the MaxFunction production.
	ExitMaxFunction(c *MaxFunctionContext)
}
