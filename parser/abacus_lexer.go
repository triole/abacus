// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 90, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 5, 3, 39, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 6, 12,
	58, 10, 12, 13, 12, 14, 12, 59, 3, 12, 3, 12, 6, 12, 64, 10, 12, 13, 12,
	14, 12, 65, 5, 12, 68, 10, 12, 3, 13, 3, 13, 7, 13, 72, 10, 13, 12, 13,
	14, 13, 75, 11, 13, 3, 14, 5, 14, 78, 10, 14, 3, 15, 3, 15, 5, 15, 82,
	10, 15, 3, 16, 6, 16, 85, 10, 16, 13, 16, 14, 16, 86, 3, 16, 3, 16, 2,
	2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 2, 25, 13, 27, 2, 29, 2, 31, 14, 3, 2, 4, 5, 2, 67, 92, 97, 97,
	99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 93, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2,
	5, 38, 3, 2, 2, 2, 7, 40, 3, 2, 2, 2, 9, 42, 3, 2, 2, 2, 11, 44, 3, 2,
	2, 2, 13, 46, 3, 2, 2, 2, 15, 48, 3, 2, 2, 2, 17, 50, 3, 2, 2, 2, 19, 52,
	3, 2, 2, 2, 21, 54, 3, 2, 2, 2, 23, 57, 3, 2, 2, 2, 25, 69, 3, 2, 2, 2,
	27, 77, 3, 2, 2, 2, 29, 81, 3, 2, 2, 2, 31, 84, 3, 2, 2, 2, 33, 34, 7,
	63, 2, 2, 34, 4, 3, 2, 2, 2, 35, 39, 7, 96, 2, 2, 36, 37, 7, 44, 2, 2,
	37, 39, 7, 44, 2, 2, 38, 35, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39, 6, 3,
	2, 2, 2, 40, 41, 7, 44, 2, 2, 41, 8, 3, 2, 2, 2, 42, 43, 7, 49, 2, 2, 43,
	10, 3, 2, 2, 2, 44, 45, 7, 45, 2, 2, 45, 12, 3, 2, 2, 2, 46, 47, 7, 47,
	2, 2, 47, 14, 3, 2, 2, 2, 48, 49, 7, 48, 2, 2, 49, 16, 3, 2, 2, 2, 50,
	51, 7, 42, 2, 2, 51, 18, 3, 2, 2, 2, 52, 53, 7, 43, 2, 2, 53, 20, 3, 2,
	2, 2, 54, 55, 5, 23, 12, 2, 55, 22, 3, 2, 2, 2, 56, 58, 4, 50, 59, 2, 57,
	56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2,
	2, 60, 67, 3, 2, 2, 2, 61, 63, 5, 15, 8, 2, 62, 64, 4, 50, 59, 2, 63, 62,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2,
	66, 68, 3, 2, 2, 2, 67, 61, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 24, 3,
	2, 2, 2, 69, 73, 5, 27, 14, 2, 70, 72, 5, 29, 15, 2, 71, 70, 3, 2, 2, 2,
	72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 26, 3,
	2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 78, 9, 2, 2, 2, 77, 76, 3, 2, 2, 2, 78,
	28, 3, 2, 2, 2, 79, 82, 5, 27, 14, 2, 80, 82, 4, 50, 59, 2, 81, 79, 3,
	2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 30, 3, 2, 2, 2, 83, 85, 9, 3, 2, 2, 84,
	83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2,
	2, 87, 88, 3, 2, 2, 2, 88, 89, 8, 16, 2, 2, 89, 32, 3, 2, 2, 2, 11, 2,
	38, 59, 65, 67, 73, 77, 81, 86, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "", "'*'", "'/'", "'+'", "'-'", "'.'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "EQ", "POW", "MUL", "DIV", "ADD", "SUB", "POINT", "LPAREN", "RPAREN",
	"SCIENTIFIC_NUMBER", "VARIABLE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"EQ", "POW", "MUL", "DIV", "ADD", "SUB", "POINT", "LPAREN", "RPAREN", "SCIENTIFIC_NUMBER",
	"NUMBER", "VARIABLE", "VALID_ID_START", "VALID_ID_CHAR", "WHITESPACE",
}

type AbacusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewAbacusLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *AbacusLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAbacusLexer(input antlr.CharStream) *AbacusLexer {
	l := new(AbacusLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Abacus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AbacusLexer tokens.
const (
	AbacusLexerEQ                = 1
	AbacusLexerPOW               = 2
	AbacusLexerMUL               = 3
	AbacusLexerDIV               = 4
	AbacusLexerADD               = 5
	AbacusLexerSUB               = 6
	AbacusLexerPOINT             = 7
	AbacusLexerLPAREN            = 8
	AbacusLexerRPAREN            = 9
	AbacusLexerSCIENTIFIC_NUMBER = 10
	AbacusLexerVARIABLE          = 11
	AbacusLexerWHITESPACE        = 12
)
