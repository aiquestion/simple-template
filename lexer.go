package simple_template

import (
	"fmt"
	"strings"
	"text/scanner"
)

var LexPrintToken = false

var TOKEN_MAP = map[string]rune{
	"if":    TIf,
	"else":  TElse,
	"==":    TEqeq,
	"!=":    TNeq,
	">=":    TGteq,
	"<=":    TLteq,
	"&&":    TAnd,
	"||":    TOr,
	"true":  TTrue,
	"false": TFalse,
}

func Parse(code string) []Expr {
	s := &scanner.Scanner{}
	s.Init(strings.NewReader(code))
	// TODO need to skip comment
	lex := &yyLexerImpl{
		scanner: s,
	}
	parser := yyNewParser()
	parser.Parse(lex)
	return lex.Stmt
}

type yyLexerImpl struct {
	scanner *scanner.Scanner
	Stmt    []Expr
}

func (y *yyLexerImpl) Lex(lval *yySymType) int {
	r := y.scanner.Scan()
	var token Token
	switch r {
	case scanner.EOF:
		return 0
	case scanner.Char:
		fallthrough
	case scanner.String:
		fallthrough
	case scanner.RawString:
		token.Type = TString
	case scanner.Ident:
		text := y.scanner.TokenText()
		token.Str = text
		if t, ok := TOKEN_MAP[text]; ok {
			token.Type = t
		} else {
			token.Type = TIdent
		}
	case scanner.Float:
		fallthrough
	case scanner.Int:
		token.Type = TNumber
	case '=':
		if y.scanner.Peek() == '=' {
			token.Str = "=="
			if t, ok := TOKEN_MAP["=="]; ok {
				token.Type = t
			}
			y.scanner.Scan()
		} else {
			token.Type = r
		}
	case '!':
		if y.scanner.Peek() == '=' {
			token.Str = "!="
			if t, ok := TOKEN_MAP["!="]; ok {
				token.Type = t
			}
			y.scanner.Scan()
		} else {
			token.Type = r
		}
	case '>':
		if y.scanner.Peek() == '=' {
			token.Str = ">="
			if t, ok := TOKEN_MAP[">="]; ok {
				token.Type = t
			}
			y.scanner.Scan()
		} else {
			token.Type = r
		}
	case '<':
		if y.scanner.Peek() == '=' {
			token.Str = "<="
			if t, ok := TOKEN_MAP["<="]; ok {
				token.Type = t
			}
			y.scanner.Scan()
		} else {
			token.Type = r
		}
	default:
		token.Type = r
	}
	if token.Str == "" {
		token.Str = y.scanner.TokenText()
	}
	tpos := y.scanner.Pos()
	token.Pos = Position{
		Source: tpos.Filename,
		Line:   tpos.Line,
		Column: tpos.Column,
	}
	lval.token = token
	if LexPrintToken {
		fmt.Printf("<<token: %s, %s>>\r\n", TokenName(int(token.Type)), token.Str)
	}
	return int(token.Type)
}
func (y *yyLexerImpl) Error(s string) {
	// TODO
	fmt.Println("#######" + s)
}
