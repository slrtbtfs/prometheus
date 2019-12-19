// Code generated by goyacc -o promql/generated_parser.y.go promql/generated_parser.y. DO NOT EDIT.

//line promql/generated_parser.y:15
package promql

import __yyfmt__ "fmt"

//line promql/generated_parser.y:15

import (
	"math"
	"sort"
	"strconv"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/value"
)

//line promql/generated_parser.y:27
type yySymType struct {
	yys      int
	node     Node
	item     Item
	matchers interface{}
	matcher  interface{}
	label    labels.Label
	labels   interface{}
	strings  interface{}
	series   interface{}
	uint     uint64
	float    float64
}

const ERROR = 57346
const EOF = 57347
const COMMENT = 57348
const IDENTIFIER = 57349
const METRIC_IDENTIFIER = 57350
const LEFT_PAREN = 57351
const RIGHT_PAREN = 57352
const LEFT_BRACE = 57353
const RIGHT_BRACE = 57354
const LEFT_BRACKET = 57355
const RIGHT_BRACKET = 57356
const COMMA = 57357
const ASSIGN = 57358
const COLON = 57359
const SEMICOLON = 57360
const STRING = 57361
const NUMBER = 57362
const DURATION = 57363
const BLANK = 57364
const TIMES = 57365
const SPACE = 57366
const operatorsStart = 57367
const SUB = 57368
const ADD = 57369
const MUL = 57370
const MOD = 57371
const DIV = 57372
const LAND = 57373
const LOR = 57374
const LUNLESS = 57375
const EQL = 57376
const NEQ = 57377
const LTE = 57378
const LSS = 57379
const GTE = 57380
const GTR = 57381
const EQL_REGEX = 57382
const NEQ_REGEX = 57383
const POW = 57384
const operatorsEnd = 57385
const aggregatorsStart = 57386
const AVG = 57387
const COUNT = 57388
const SUM = 57389
const MIN = 57390
const MAX = 57391
const STDDEV = 57392
const STDVAR = 57393
const TOPK = 57394
const BOTTOMK = 57395
const COUNT_VALUES = 57396
const QUANTILE = 57397
const aggregatorsEnd = 57398
const keywordsStart = 57399
const OFFSET = 57400
const BY = 57401
const WITHOUT = 57402
const ON = 57403
const IGNORING = 57404
const GROUP_LEFT = 57405
const GROUP_RIGHT = 57406
const BOOL = 57407
const keywordsEnd = 57408
const startSymbolsStart = 57409
const START_LABELS = 57410
const START_LABEL_SET = 57411
const START_METRIC = 57412
const START_GROUPING_LABELS = 57413
const START_SERIES_DESCRIPTION = 57414
const startSymbolsEnd = 57415

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ERROR",
	"EOF",
	"COMMENT",
	"IDENTIFIER",
	"METRIC_IDENTIFIER",
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"LEFT_BRACKET",
	"RIGHT_BRACKET",
	"COMMA",
	"ASSIGN",
	"COLON",
	"SEMICOLON",
	"STRING",
	"NUMBER",
	"DURATION",
	"BLANK",
	"TIMES",
	"SPACE",
	"operatorsStart",
	"SUB",
	"ADD",
	"MUL",
	"MOD",
	"DIV",
	"LAND",
	"LOR",
	"LUNLESS",
	"EQL",
	"NEQ",
	"LTE",
	"LSS",
	"GTE",
	"GTR",
	"EQL_REGEX",
	"NEQ_REGEX",
	"POW",
	"operatorsEnd",
	"aggregatorsStart",
	"AVG",
	"COUNT",
	"SUM",
	"MIN",
	"MAX",
	"STDDEV",
	"STDVAR",
	"TOPK",
	"BOTTOMK",
	"COUNT_VALUES",
	"QUANTILE",
	"aggregatorsEnd",
	"keywordsStart",
	"OFFSET",
	"BY",
	"WITHOUT",
	"ON",
	"IGNORING",
	"GROUP_LEFT",
	"GROUP_RIGHT",
	"BOOL",
	"keywordsEnd",
	"startSymbolsStart",
	"START_LABELS",
	"START_LABEL_SET",
	"START_METRIC",
	"START_GROUPING_LABELS",
	"START_SERIES_DESCRIPTION",
	"startSymbolsEnd",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line promql/generated_parser.y:383

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 4,
	1, 31,
	5, 31,
	-2, 0,
	-1, 22,
	1, 31,
	5, 31,
	24, 31,
	-2, 0,
	-1, 63,
	1, 71,
	5, 71,
	24, 71,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 159

var yyAct = [...]int{

	107, 97, 98, 38, 36, 25, 68, 96, 39, 40,
	30, 35, 7, 103, 13, 92, 101, 100, 110, 102,
	99, 8, 94, 108, 90, 77, 101, 100, 99, 12,
	75, 10, 41, 42, 43, 21, 82, 63, 69, 70,
	73, 89, 20, 74, 71, 72, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 76, 62, 55,
	56, 22, 57, 58, 59, 60, 61, 1, 95, 86,
	38, 84, 93, 81, 34, 39, 40, 66, 2, 3,
	4, 5, 6, 80, 91, 88, 85, 64, 19, 28,
	65, 78, 37, 15, 14, 67, 79, 11, 104, 41,
	42, 43, 105, 106, 109, 23, 9, 0, 33, 0,
	0, 111, 0, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 0, 0, 55, 56, 0, 57,
	58, 59, 60, 61, 32, 27, 0, 16, 0, 31,
	26, 0, 18, 17, 87, 83, 12, 32, 27, 0,
	0, 0, 31, 26, 0, 0, 0, 29, 24,
}
var yyPact = [...]int{

	10, 16, 20, 18, 135, 33, -1000, -1000, -1000, -1000,
	146, -1000, 145, -1000, 18, -1000, -1000, -1000, -1000, -1000,
	1, -1000, 135, 75, -1000, -1000, 4, -1000, 28, -1000,
	-1000, 23, -1000, -1000, 81, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 34, -1000, 133, -1000, 67, -1000, -1000,
	-1000, -1000, -1000, -1000, 132, -1000, 22, -1000, -1000, 68,
	-1000, -9, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 0, -1000, -4, -10, -1000, -1000, -1000, -1000,
	8, 8, 3, 3, -5, -1000, -1000, -1000, -1000, -1000,
	3, -1000,
}
var yyPgo = [...]int{

	0, 106, 105, 5, 95, 94, 4, 92, 89, 93,
	14, 10, 88, 74, 73, 72, 0, 68, 2, 1,
	67, 61, 58,
}
var yyR1 = [...]int{

	0, 20, 20, 20, 20, 21, 20, 20, 20, 1,
	1, 1, 2, 2, 2, 3, 3, 3, 3, 4,
	4, 4, 4, 10, 10, 10, 5, 5, 9, 9,
	9, 9, 8, 8, 8, 11, 11, 11, 11, 12,
	12, 12, 13, 13, 13, 6, 6, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	22, 14, 14, 14, 14, 15, 15, 15, 15, 15,
	16, 18, 18, 17, 17, 17, 19,
}
var yyR2 = [...]int{

	0, 2, 2, 2, 2, 0, 3, 2, 1, 3,
	4, 2, 3, 1, 2, 3, 3, 2, 1, 1,
	1, 1, 1, 2, 1, 1, 1, 1, 3, 4,
	2, 0, 3, 1, 2, 3, 3, 2, 1, 3,
	2, 1, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 0, 3, 2, 1, 1, 3, 1, 3, 4,
	1, 2, 2, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -20, 68, 69, 70, 71, 72, 2, 5, -1,
	11, -9, 11, -10, -5, -9, 2, 8, 7, -12,
	9, 2, -21, -2, 12, -3, 7, 2, -8, 12,
	-11, 7, 2, -9, -13, 10, -6, -7, 2, 7,
	8, 31, 32, 33, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 58, 59, 61, 62, 63,
	64, 65, -22, -10, 12, 15, 2, -4, 2, 34,
	35, 40, 41, 12, 15, 2, 34, 2, 10, 15,
	2, -14, 2, 12, -3, 19, 2, 12, -11, 19,
	2, -6, 24, -15, 22, -17, 7, -19, -18, 20,
	27, 26, 23, 23, -18, -19, -19, -16, 20, -16,
	23, -16,
}
var yyDef = [...]int{

	0, -2, 0, 31, -2, 0, 5, 8, 7, 1,
	0, 2, 0, 3, 31, 24, 25, 26, 27, 4,
	0, 41, -2, 0, 11, 13, 0, 18, 0, 30,
	33, 0, 38, 23, 0, 40, 43, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 6, -2, 9, 0, 14, 0, 17, 19,
	20, 21, 22, 28, 0, 34, 0, 37, 39, 0,
	44, 70, 74, 10, 12, 15, 16, 29, 32, 35,
	36, 42, 73, 72, 75, 77, 83, 84, 85, 86,
	0, 0, 0, 0, 0, 81, 82, 76, 80, 78,
	0, 79,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73,
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
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:139
		{
			yylex.(*parser).generatedParserResult.(*VectorSelector).LabelMatchers = yyDollar[2].matchers.([]*labels.Matcher)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:141
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].labels
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:143
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].labels
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:145
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].strings
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:146
		{
			yylex.(*parser).generatedParserResult = &seriesDescription{}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:149
		{
			yylex.(*parser).unexpected("", "")
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:155
		{
			yyVAL.matchers = yyDollar[2].matchers
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:157
		{
			yyVAL.matchers = yyDollar[2].matchers
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:159
		{
			yyVAL.matchers = []*labels.Matcher{}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:165
		{
			yyVAL.matchers = append(yyDollar[1].matchers.([]*labels.Matcher), yyDollar[3].matcher.(*labels.Matcher))
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:167
		{
			yyVAL.matchers = []*labels.Matcher{yyDollar[1].matcher.(*labels.Matcher)}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:169
		{
			yylex.(*parser).unexpected("label matching", "\",\" or \"}\"")
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:174
		{
			yyVAL.matcher = yylex.(*parser).newLabelMatcher(yyDollar[1].item, yyDollar[2].item, yyDollar[3].item)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:176
		{
			yylex.(*parser).unexpected("label matching", "string")
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:178
		{
			yylex.(*parser).unexpected("label matching", "label matching operator")
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:180
		{
			yylex.(*parser).unexpected("label matching", "identifier or \"}\"")
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:184
		{
			yyVAL.item = yyDollar[1].item
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:185
		{
			yyVAL.item = yyDollar[1].item
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:186
		{
			yyVAL.item = yyDollar[1].item
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:187
		{
			yyVAL.item = yyDollar[1].item
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:193
		{
			yyVAL.labels = append(yyDollar[2].labels.(labels.Labels), labels.Label{Name: labels.MetricName, Value: yyDollar[1].item.Val})
			sort.Sort(yyVAL.labels.(labels.Labels))
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:195
		{
			yyVAL.labels = yyDollar[1].labels
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:197
		{
			yylex.(*parser).errorf("missing metric name or metric selector")
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:202
		{
			yyVAL.item = yyDollar[1].item
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:203
		{
			yyVAL.item = yyDollar[1].item
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:207
		{
			yyVAL.labels = labels.New(yyDollar[2].labels.([]labels.Label)...)
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:209
		{
			yyVAL.labels = labels.New(yyDollar[2].labels.([]labels.Label)...)
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:211
		{
			yyVAL.labels = labels.New()
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:213
		{
			yyVAL.labels = labels.New()
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:218
		{
			yyVAL.labels = append(yyDollar[1].labels.([]labels.Label), yyDollar[3].label)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:220
		{
			yyVAL.labels = []labels.Label{yyDollar[1].label}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:222
		{
			yylex.(*parser).unexpected("label set", "\",\" or \"}\"")
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:228
		{
			yyVAL.label = labels.Label{Name: yyDollar[1].item.Val, Value: yylex.(*parser).unquoteString(yyDollar[3].item.Val)}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:230
		{
			yylex.(*parser).unexpected("label set", "string")
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:232
		{
			yylex.(*parser).unexpected("label set", "\"=\"")
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:234
		{
			yylex.(*parser).unexpected("label set", "identifier or \"}\"")
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:239
		{
			yyVAL.strings = yyDollar[2].strings
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:241
		{
			yyVAL.strings = []string{}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:243
		{
			yylex.(*parser).unexpected("grouping opts", "\"(\"")
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:249
		{
			yyVAL.strings = append(yyDollar[1].strings.([]string), yyDollar[3].item.Val)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:251
		{
			yyVAL.strings = []string{yyDollar[1].item.Val}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:253
		{
			yylex.(*parser).unexpected("grouping opts", "\",\" or \"}\"")
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:258
		{
			if !isLabel(yyDollar[1].item.Val) {
				yylex.(*parser).unexpected("grouping opts", "label")
			}
			yyVAL.item = yyDollar[1].item
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:265
		{
			yylex.(*parser).unexpected("grouping opts", "label")
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:298
		{
			yylex.(*parser).generatedParserResult = &seriesDescription{
				labels: yyDollar[1].labels.(labels.Labels),
				values: yyDollar[2].series.([]sequenceValue),
			}
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:308
		{
			yyVAL.series = []sequenceValue{}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:310
		{
			yyVAL.series = append(yyDollar[1].series.([]sequenceValue), yyDollar[3].series.([]sequenceValue)...)
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:312
		{
			yyVAL.series = yyDollar[1].series.([]sequenceValue)
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:314
		{
			yylex.(*parser).unexpected("series values", "")
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:319
		{
			yyVAL.series = []sequenceValue{{omitted: true}}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:321
		{
			yyVAL.series = []sequenceValue{}
			for i := uint64(0); i < yyDollar[3].uint; i++ {
				yyVAL.series = append(yyVAL.series.([]sequenceValue), sequenceValue{omitted: true})
			}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:328
		{
			yyVAL.series = []sequenceValue{{value: yyDollar[1].float}}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:330
		{
			yyVAL.series = []sequenceValue{}
			for i := uint64(0); i <= yyDollar[3].uint; i++ {
				yyVAL.series = append(yyVAL.series.([]sequenceValue), sequenceValue{value: yyDollar[1].float})
			}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:337
		{
			yyVAL.series = []sequenceValue{}
			for i := uint64(0); i <= yyDollar[4].uint; i++ {
				yyVAL.series = append(yyVAL.series.([]sequenceValue), sequenceValue{value: yyDollar[1].float})
				yyDollar[1].float += yyDollar[2].float
			}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:346
		{
			var err error
			yyVAL.uint, err = strconv.ParseUint(yyDollar[1].item.Val, 10, 64)
			if err != nil {
				yylex.(*parser).errorf("invalid repitition in series values: %s", err)
			}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:357
		{
			yyVAL.float = yyDollar[2].float
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:359
		{
			yyVAL.float = -yyDollar[2].float
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:364
		{
			if yyDollar[1].item.Val != "stale" {
				yylex.(*parser).unexpected("series values", "number or \"stale\"")
			}
			yyVAL.float = math.Float64frombits(value.StaleNaN)
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:371
		{
			yyVAL.float = yyDollar[1].float
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:373
		{
			yyVAL.float = yyDollar[1].float
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:380
		{
			yyVAL.float = yylex.(*parser).number(yyDollar[1].item.Val)
		}
	}
	goto yystack /* stack new state and value */
}
