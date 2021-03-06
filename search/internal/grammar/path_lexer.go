// Code generated from search/internal/grammar/Path.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 43, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 25, 10,
	3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 6, 9, 40, 10, 9, 13, 9, 14, 9, 41, 2, 2, 10, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 3, 2, 3, 6, 2, 35, 35, 46, 46, 49, 49,
	63, 63, 2, 44, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 24, 3, 2, 2, 2, 7, 26, 3, 2, 2,
	2, 9, 28, 3, 2, 2, 2, 11, 31, 3, 2, 2, 2, 13, 34, 3, 2, 2, 2, 15, 36, 3,
	2, 2, 2, 17, 39, 3, 2, 2, 2, 19, 20, 7, 49, 2, 2, 20, 4, 3, 2, 2, 2, 21,
	25, 7, 44, 2, 2, 22, 23, 7, 44, 2, 2, 23, 25, 7, 44, 2, 2, 24, 21, 3, 2,
	2, 2, 24, 22, 3, 2, 2, 2, 25, 6, 3, 2, 2, 2, 26, 27, 7, 63, 2, 2, 27, 8,
	3, 2, 2, 2, 28, 29, 7, 35, 2, 2, 29, 30, 7, 63, 2, 2, 30, 10, 3, 2, 2,
	2, 31, 32, 7, 35, 2, 2, 32, 33, 7, 44, 2, 2, 33, 12, 3, 2, 2, 2, 34, 35,
	7, 35, 2, 2, 35, 14, 3, 2, 2, 2, 36, 37, 7, 46, 2, 2, 37, 16, 3, 2, 2,
	2, 38, 40, 10, 2, 2, 2, 39, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 39,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 18, 3, 2, 2, 2, 5, 2, 24, 41, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'/'", "", "'='", "'!='", "'!*'", "'!'", "','",
}

var lexerSymbolicNames = []string{
	"", "AND", "ANY", "EQ", "NEQ", "NONE", "NOT", "OR", "STRING",
}

var lexerRuleNames = []string{
	"AND", "ANY", "EQ", "NEQ", "NONE", "NOT", "OR", "STRING",
}

type PathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPathLexer(input antlr.CharStream) *PathLexer {

	l := new(PathLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Path.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PathLexer tokens.
const (
	PathLexerAND    = 1
	PathLexerANY    = 2
	PathLexerEQ     = 3
	PathLexerNEQ    = 4
	PathLexerNONE   = 5
	PathLexerNOT    = 6
	PathLexerOR     = 7
	PathLexerSTRING = 8
)
