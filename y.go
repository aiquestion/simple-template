//line parser.go.y:2
package simple_template

import __yyfmt__ "fmt"

//line parser.go.y:2
/**
THIS FILE IS AUTO-GENERATE BY GOYACC
DO NOT EDIT
**/

import "strconv"

//line parser.go.y:20
type yySymType struct {
	yys   int
	token Token
	expr  Expr
	exprs []Expr
}

const TIf = 57346
const TElse = 57347
const TEqeq = 57348
const TNeq = 57349
const TGteq = 57350
const TLteq = 57351
const TAnd = 57352
const TOr = 57353
const TTrue = 57354
const TFalse = 57355
const TIdent = 57356
const TString = 57357
const TNumber = 57358

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TIf",
	"TElse",
	"TEqeq",
	"TNeq",
	"TGteq",
	"TLteq",
	"TAnd",
	"TOr",
	"TTrue",
	"TFalse",
	"TIdent",
	"TString",
	"TNumber",
	"'('",
	"'{'",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"','",
	"')'",
	"'}'",
	"'['",
	"']'",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:195

func TokenName(c int) string {
	if c >= TAnd && c-TAnd < len(yyToknames) {
		if yyToknames[c-TAnd] != "" {
			return yyToknames[c-TAnd]
		}
	}
	return string([]byte{byte(c)})
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 36
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 245

var yyAct = [...]int{

	2, 55, 33, 54, 1, 23, 24, 25, 26, 27,
	32, 28, 14, 34, 30, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 14, 70, 53, 21, 22, 19, 20, 16, 15,
	66, 59, 14, 65, 25, 26, 27, 17, 18, 23,
	24, 25, 26, 27, 36, 29, 63, 64, 62, 56,
	21, 22, 19, 20, 16, 15, 69, 67, 35, 31,
	10, 71, 11, 17, 18, 23, 24, 25, 26, 27,
	7, 0, 0, 0, 60, 21, 22, 19, 20, 16,
	15, 0, 0, 0, 0, 0, 0, 0, 17, 18,
	23, 24, 25, 26, 27, 0, 0, 72, 21, 22,
	19, 20, 16, 15, 0, 0, 0, 0, 0, 0,
	0, 17, 18, 23, 24, 25, 26, 27, 0, 0,
	68, 21, 22, 19, 20, 16, 15, 0, 0, 0,
	0, 0, 0, 0, 17, 18, 23, 24, 25, 26,
	27, 0, 61, 21, 22, 19, 20, 16, 15, 0,
	0, 0, 0, 0, 0, 0, 17, 18, 23, 24,
	25, 26, 27, 0, 57, 21, 22, 19, 20, 16,
	15, 0, 21, 22, 19, 20, 16, 0, 17, 18,
	23, 24, 25, 26, 27, 17, 18, 23, 24, 25,
	26, 27, 21, 22, 19, 20, 0, 0, 0, 0,
	0, 0, 0, 0, 8, 17, 18, 23, 24, 25,
	26, 27, 5, 6, 13, 3, 4, 12, 9, 0,
	8, 0, 0, 0, 0, 0, 0, 58, 5, 6,
	13, 3, 4, 12, 9,
}
var yyPact = [...]int{

	226, -14, 169, -1000, -1000, -1000, -1000, -18, 38, 226,
	-1000, -1000, 226, 37, 226, 226, 226, 226, 226, 226,
	226, 226, 226, 226, 226, 226, 226, 226, 226, 226,
	5, -25, 28, -1000, 147, -1000, 210, 169, 176, 196,
	-16, -16, -16, -16, -16, -16, 21, 21, -1000, -1000,
	-1000, 54, 125, -1000, -1000, 226, 226, -1000, -1000, 16,
	-1000, 22, -1000, 28, 169, -1000, 226, 102, 61, 14,
	226, 79, -1000,
}
var yyPgo = [...]int{

	0, 0, 80, 72, 70, 2, 69, 68, 4,
}
var yyR1 = [...]int{

	0, 8, 8, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 3, 7,
	7, 4, 4, 6, 6, 5,
}
var yyR2 = [...]int{

	0, 1, 3, 1, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 7, 11, 3, 3, 1, 1, 3, 2, 2,
	3, 1, 4, 1, 3, 3,
}
var yyChk = [...]int{

	-1000, -8, -1, 15, 16, 12, 13, -2, 4, 18,
	-4, -3, 17, 14, 26, 11, 10, 19, 20, 8,
	9, 6, 7, 21, 22, 23, 24, 25, 29, 17,
	-8, -6, -1, -5, -1, -7, 17, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, 28, 28, 26, 31, 27, 27, -8,
	30, 27, -5, -1, -1, 27, 18, -1, 28, 5,
	18, -1, 28,
}
var yyDef = [...]int{

	0, -2, 1, 3, 4, 5, 6, 20, 0, 0,
	25, 26, 0, 31, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 1, 33, 0, 28, 0, 2, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 0, 0, 23, 24, 0, 0, 27, 29, 0,
	32, 0, 34, 0, 35, 30, 0, 0, 21, 0,
	0, 0, 22,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 25, 3, 3,
	17, 27, 23, 21, 26, 22, 3, 24, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 31, 3,
	20, 3, 19, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 29, 3, 30, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 18, 3, 28,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:39
		{
			yyVAL.exprs = []Expr{yyDollar[1].expr}
			if l, ok := yylex.(*yyLexerImpl); ok {
				l.Stmt = yyVAL.exprs
			}
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:45
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
			if l, ok := yylex.(*yyLexerImpl); ok {
				l.Stmt = yyVAL.exprs
			}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:53
		{
			yyVAL.expr = &StringExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:57
		{
			v, _ := strconv.ParseFloat(yyDollar[1].token.Str, 0)
			yyVAL.expr = &NumberExpr{Value: v}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:62
		{
			yyVAL.expr = &TrueExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:66
		{
			yyVAL.expr = &FalseExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:70
		{
			yyVAL.expr = &LogicalExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:74
		{
			yyVAL.expr = &LogicalExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:78
		{
			yyVAL.expr = &LogicalExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.expr = &LogicalExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:86
		{
			yyVAL.expr = &LogicalExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:90
		{
			yyVAL.expr = &LogicalExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:94
		{
			yyVAL.expr = &LogicalExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:98
		{
			yyVAL.expr = &LogicalExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.expr = &ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.expr = &ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:110
		{
			yyVAL.expr = &ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.expr = &ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:118
		{
			yyVAL.expr = &ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:122
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:125
		{
			yyVAL.expr = &IfExpr{Cond: yyDollar[3].expr, Then: yyDollar[6].expr}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 22:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line parser.go.y:129
		{
			yyVAL.expr = &IfExpr{Cond: yyDollar[3].expr, Then: yyDollar[6].expr, Else: yyDollar[10].expr}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.expr = &ArrayConstructExpr{Members: yyDollar[2].exprs}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:137
		{
			yyVAL.expr = &MapConstructExpr{Members: yyDollar[2].exprs}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:156
		{
			yyVAL.expr = &FuncCallExpr{Func: &IdentExpr{Value: yyDollar[1].token.Str}, Args: yyDollar[2].exprs}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.exprs = []Expr{}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:165
		{
			yyVAL.exprs = yyDollar[2].exprs
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.expr = &IdentExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Line())
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.expr = &AttrGetExpr{Object: yyDollar[1].expr, Key: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:182
		{
			yyVAL.exprs = []Expr{yyDollar[1].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:185
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.expr = &MapItemExpr{Key: yyDollar[1].expr, Value: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	}
	goto yystack /* stack new state and value */
}
