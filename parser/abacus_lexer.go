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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 30, 189,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 5, 18, 130, 10, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 152, 10, 26,
	3, 27, 3, 27, 3, 28, 6, 28, 157, 10, 28, 13, 28, 14, 28, 158, 3, 28, 3,
	28, 6, 28, 163, 10, 28, 13, 28, 14, 28, 164, 5, 28, 167, 10, 28, 3, 29,
	3, 29, 7, 29, 171, 10, 29, 12, 29, 14, 29, 174, 11, 29, 3, 30, 5, 30, 177,
	10, 30, 3, 31, 3, 31, 5, 31, 181, 10, 31, 3, 32, 6, 32, 184, 10, 32, 13,
	32, 14, 32, 185, 3, 32, 3, 32, 2, 2, 33, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49,
	26, 51, 27, 53, 28, 55, 2, 57, 29, 59, 2, 61, 2, 63, 30, 3, 2, 4, 5, 2,
	67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 194, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 63, 3,
	2, 2, 2, 3, 65, 3, 2, 2, 2, 5, 70, 3, 2, 2, 2, 7, 73, 3, 2, 2, 2, 9, 79,
	3, 2, 2, 2, 11, 84, 3, 2, 2, 2, 13, 88, 3, 2, 2, 2, 15, 92, 3, 2, 2, 2,
	17, 96, 3, 2, 2, 2, 19, 100, 3, 2, 2, 2, 21, 106, 3, 2, 2, 2, 23, 108,
	3, 2, 2, 2, 25, 112, 3, 2, 2, 2, 27, 116, 3, 2, 2, 2, 29, 120, 3, 2, 2,
	2, 31, 122, 3, 2, 2, 2, 33, 124, 3, 2, 2, 2, 35, 129, 3, 2, 2, 2, 37, 131,
	3, 2, 2, 2, 39, 133, 3, 2, 2, 2, 41, 135, 3, 2, 2, 2, 43, 137, 3, 2, 2,
	2, 45, 139, 3, 2, 2, 2, 47, 141, 3, 2, 2, 2, 49, 143, 3, 2, 2, 2, 51, 151,
	3, 2, 2, 2, 53, 153, 3, 2, 2, 2, 55, 156, 3, 2, 2, 2, 57, 168, 3, 2, 2,
	2, 59, 176, 3, 2, 2, 2, 61, 180, 3, 2, 2, 2, 63, 183, 3, 2, 2, 2, 65, 66,
	7, 117, 2, 2, 66, 67, 7, 115, 2, 2, 67, 68, 7, 116, 2, 2, 68, 69, 7, 118,
	2, 2, 69, 4, 3, 2, 2, 2, 70, 71, 7, 110, 2, 2, 71, 72, 7, 112, 2, 2, 72,
	6, 3, 2, 2, 2, 73, 74, 7, 104, 2, 2, 74, 75, 7, 110, 2, 2, 75, 76, 7, 113,
	2, 2, 76, 77, 7, 113, 2, 2, 77, 78, 7, 116, 2, 2, 78, 8, 3, 2, 2, 2, 79,
	80, 7, 101, 2, 2, 80, 81, 7, 103, 2, 2, 81, 82, 7, 107, 2, 2, 82, 83, 7,
	110, 2, 2, 83, 10, 3, 2, 2, 2, 84, 85, 7, 103, 2, 2, 85, 86, 7, 122, 2,
	2, 86, 87, 7, 114, 2, 2, 87, 12, 3, 2, 2, 2, 88, 89, 7, 117, 2, 2, 89,
	90, 7, 107, 2, 2, 90, 91, 7, 112, 2, 2, 91, 14, 3, 2, 2, 2, 92, 93, 7,
	101, 2, 2, 93, 94, 7, 113, 2, 2, 94, 95, 7, 117, 2, 2, 95, 16, 3, 2, 2,
	2, 96, 97, 7, 118, 2, 2, 97, 98, 7, 99, 2, 2, 98, 99, 7, 112, 2, 2, 99,
	18, 3, 2, 2, 2, 100, 101, 7, 116, 2, 2, 101, 102, 7, 113, 2, 2, 102, 103,
	7, 119, 2, 2, 103, 104, 7, 112, 2, 2, 104, 105, 7, 102, 2, 2, 105, 20,
	3, 2, 2, 2, 106, 107, 7, 46, 2, 2, 107, 22, 3, 2, 2, 2, 108, 109, 7, 110,
	2, 2, 109, 110, 7, 113, 2, 2, 110, 111, 7, 105, 2, 2, 111, 24, 3, 2, 2,
	2, 112, 113, 7, 111, 2, 2, 113, 114, 7, 107, 2, 2, 114, 115, 7, 112, 2,
	2, 115, 26, 3, 2, 2, 2, 116, 117, 7, 111, 2, 2, 117, 118, 7, 99, 2, 2,
	118, 119, 7, 122, 2, 2, 119, 28, 3, 2, 2, 2, 120, 121, 7, 63, 2, 2, 121,
	30, 3, 2, 2, 2, 122, 123, 7, 62, 2, 2, 123, 32, 3, 2, 2, 2, 124, 125, 7,
	64, 2, 2, 125, 34, 3, 2, 2, 2, 126, 130, 7, 96, 2, 2, 127, 128, 7, 44,
	2, 2, 128, 130, 7, 44, 2, 2, 129, 126, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2,
	130, 36, 3, 2, 2, 2, 131, 132, 7, 44, 2, 2, 132, 38, 3, 2, 2, 2, 133, 134,
	7, 49, 2, 2, 134, 40, 3, 2, 2, 2, 135, 136, 7, 45, 2, 2, 136, 42, 3, 2,
	2, 2, 137, 138, 7, 47, 2, 2, 138, 44, 3, 2, 2, 2, 139, 140, 7, 48, 2, 2,
	140, 46, 3, 2, 2, 2, 141, 142, 7, 42, 2, 2, 142, 48, 3, 2, 2, 2, 143, 144,
	7, 43, 2, 2, 144, 50, 3, 2, 2, 2, 145, 146, 7, 114, 2, 2, 146, 152, 7,
	107, 2, 2, 147, 152, 7, 103, 2, 2, 148, 149, 7, 114, 2, 2, 149, 150, 7,
	106, 2, 2, 150, 152, 7, 107, 2, 2, 151, 145, 3, 2, 2, 2, 151, 147, 3, 2,
	2, 2, 151, 148, 3, 2, 2, 2, 152, 52, 3, 2, 2, 2, 153, 154, 5, 55, 28, 2,
	154, 54, 3, 2, 2, 2, 155, 157, 4, 50, 59, 2, 156, 155, 3, 2, 2, 2, 157,
	158, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 166,
	3, 2, 2, 2, 160, 162, 5, 45, 23, 2, 161, 163, 4, 50, 59, 2, 162, 161, 3,
	2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2,
	2, 165, 167, 3, 2, 2, 2, 166, 160, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167,
	56, 3, 2, 2, 2, 168, 172, 5, 59, 30, 2, 169, 171, 5, 61, 31, 2, 170, 169,
	3, 2, 2, 2, 171, 174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2,
	2, 2, 173, 58, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 175, 177, 9, 2, 2, 2,
	176, 175, 3, 2, 2, 2, 177, 60, 3, 2, 2, 2, 178, 181, 5, 59, 30, 2, 179,
	181, 4, 50, 59, 2, 180, 178, 3, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 62,
	3, 2, 2, 2, 182, 184, 9, 3, 2, 2, 183, 182, 3, 2, 2, 2, 184, 185, 3, 2,
	2, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2,
	187, 188, 8, 32, 2, 2, 188, 64, 3, 2, 2, 2, 12, 2, 129, 151, 158, 164,
	166, 172, 176, 180, 185, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'sqrt'", "'ln'", "'floor'", "'ceil'", "'exp'", "'sin'", "'cos'", "'tan'",
	"'round'", "','", "'log'", "'min'", "'max'", "'='", "'<'", "'>'", "", "'*'",
	"'/'", "'+'", "'-'", "'.'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "EQ", "LS", "GR",
	"POW", "MUL", "DIV", "ADD", "SUB", "POINT", "LPAREN", "RPAREN", "CONSTANT",
	"SCIENTIFIC_NUMBER", "VARIABLE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "EQ", "LS", "GR", "POW", "MUL", "DIV",
	"ADD", "SUB", "POINT", "LPAREN", "RPAREN", "CONSTANT", "SCIENTIFIC_NUMBER",
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
	AbacusLexerT__0              = 1
	AbacusLexerT__1              = 2
	AbacusLexerT__2              = 3
	AbacusLexerT__3              = 4
	AbacusLexerT__4              = 5
	AbacusLexerT__5              = 6
	AbacusLexerT__6              = 7
	AbacusLexerT__7              = 8
	AbacusLexerT__8              = 9
	AbacusLexerT__9              = 10
	AbacusLexerT__10             = 11
	AbacusLexerT__11             = 12
	AbacusLexerT__12             = 13
	AbacusLexerEQ                = 14
	AbacusLexerLS                = 15
	AbacusLexerGR                = 16
	AbacusLexerPOW               = 17
	AbacusLexerMUL               = 18
	AbacusLexerDIV               = 19
	AbacusLexerADD               = 20
	AbacusLexerSUB               = 21
	AbacusLexerPOINT             = 22
	AbacusLexerLPAREN            = 23
	AbacusLexerRPAREN            = 24
	AbacusLexerCONSTANT          = 25
	AbacusLexerSCIENTIFIC_NUMBER = 26
	AbacusLexerVARIABLE          = 27
	AbacusLexerWHITESPACE        = 28
)
