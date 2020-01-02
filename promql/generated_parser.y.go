// Code generated by goyacc -o promql/generated_parser.y.go promql/generated_parser.y. DO NOT EDIT.

//line promql/generated_parser.y:15
package promql

import __yyfmt__ "fmt"

//line promql/generated_parser.y:15

import (
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/value"
)

//line promql/generated_parser.y:28
type yySymType struct {
	yys      int
	node     Node
	item     Item
	matchers []*labels.Matcher
	matcher  *labels.Matcher
	label    labels.Label
	labels   labels.Labels
	strings  []string
	series   []sequenceValue
	uint     uint64
	float    float64
	string   string
	duration time.Duration
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
const START_METRIC = 57411
const START_GROUPING_LABELS = 57412
const START_SERIES_DESCRIPTION = 57413
const START_EXPRESSION = 57414
const START_METRIC_SELECTOR = 57415
const startSymbolsEnd = 57416

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
	"START_METRIC",
	"START_GROUPING_LABELS",
	"START_SERIES_DESCRIPTION",
	"START_EXPRESSION",
	"START_METRIC_SELECTOR",
	"startSymbolsEnd",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line promql/generated_parser.y:676

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 22,
	1, 145,
	5, 145,
	24, 145,
	-2, 0,
	-1, 193,
	1, 48,
	5, 48,
	7, 48,
	8, 48,
	10, 48,
	11, 48,
	13, 48,
	15, 48,
	19, 48,
	20, 48,
	26, 48,
	27, 48,
	28, 48,
	29, 48,
	30, 48,
	31, 48,
	32, 48,
	33, 48,
	34, 48,
	35, 48,
	36, 48,
	37, 48,
	38, 48,
	39, 48,
	42, 48,
	45, 48,
	46, 48,
	47, 48,
	48, 48,
	49, 48,
	50, 48,
	51, 48,
	52, 48,
	53, 48,
	54, 48,
	55, 48,
	58, 48,
	59, 48,
	60, 48,
	-2, 0,
	-1, 194,
	1, 48,
	5, 48,
	7, 48,
	8, 48,
	10, 48,
	11, 48,
	13, 48,
	15, 48,
	19, 48,
	20, 48,
	26, 48,
	27, 48,
	28, 48,
	29, 48,
	30, 48,
	31, 48,
	32, 48,
	33, 48,
	34, 48,
	35, 48,
	36, 48,
	37, 48,
	38, 48,
	39, 48,
	42, 48,
	45, 48,
	46, 48,
	47, 48,
	48, 48,
	49, 48,
	50, 48,
	51, 48,
	52, 48,
	53, 48,
	54, 48,
	55, 48,
	58, 48,
	59, 48,
	60, 48,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 438

var yyAct = [...]int{

	23, 228, 164, 221, 37, 189, 220, 72, 18, 67,
	193, 194, 139, 148, 196, 195, 61, 8, 216, 74,
	130, 191, 190, 122, 75, 76, 187, 182, 144, 143,
	231, 215, 165, 163, 229, 232, 118, 119, 181, 45,
	45, 185, 211, 123, 138, 191, 190, 137, 77, 78,
	79, 40, 131, 132, 10, 180, 177, 135, 133, 134,
	136, 212, 80, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 176, 117, 91, 92, 11, 93, 94,
	95, 96, 97, 2, 3, 4, 5, 6, 7, 142,
	128, 120, 17, 213, 69, 14, 63, 140, 214, 68,
	126, 62, 141, 127, 178, 69, 174, 63, 121, 64,
	68, 20, 62, 30, 99, 66, 9, 60, 19, 12,
	166, 58, 169, 170, 21, 22, 1, 225, 172, 173,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 162, 175, 192, 179, 171, 188, 183,
	197, 198, 199, 200, 201, 202, 203, 204, 205, 206,
	207, 208, 209, 210, 16, 15, 16, 15, 11, 39,
	17, 38, 168, 13, 146, 13, 147, 145, 74, 34,
	33, 32, 31, 75, 76, 29, 71, 28, 27, 26,
	25, 24, 217, 186, 184, 218, 219, 98, 70, 223,
	224, 222, 65, 42, 36, 73, 129, 77, 78, 79,
	59, 0, 0, 0, 226, 227, 0, 0, 230, 0,
	0, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 0, 233, 91, 92, 0, 93, 94, 95,
	96, 97, 41, 15, 35, 0, 11, 0, 0, 0,
	0, 0, 0, 0, 46, 45, 0, 0, 0, 0,
	0, 44, 43, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 56,
	57, 116, 0, 0, 124, 125, 0, 0, 0, 0,
	0, 0, 0, 0, 105, 104, 101, 103, 102, 112,
	114, 113, 106, 107, 108, 109, 110, 111, 0, 0,
	100, 0, 0, 167, 0, 0, 116, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 115, 124, 125, 105,
	104, 101, 103, 102, 112, 114, 113, 106, 107, 108,
	109, 110, 111, 0, 0, 100, 0, 0, 0, 116,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 115, 105, 104, 101, 103, 102, 112, 114, 113,
	106, 107, 108, 109, 110, 111, 0, 0, 100, 41,
	15, 35, 0, 11, 0, 0, 0, 0, 0, 0,
	0, 46, 45, 0, 115, 0, 0, 0, 44, 43,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57,
}
var yyPact = [...]int{

	15, 111, 66, 159, 109, 159, 382, 157, -1000, -1000,
	-1000, 105, -1000, 81, -1000, -1000, -1000, 103, -1000, 176,
	-1000, -1000, 112, 346, -1000, -1000, -1000, -1000, -1000, -1000,
	61, -1000, -1000, -1000, -1000, 382, 382, -1000, -1000, 66,
	-1000, 99, 235, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 88,
	-1000, -1000, 18, -1000, -1000, 45, -1000, -1000, 10, -1000,
	87, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 5, -1000,
	-52, -52, -52, -52, -52, -52, -52, -52, -52, -52,
	-52, -52, -52, -52, -52, 12, 11, 11, 313, 346,
	-1000, 382, 382, 278, 109, 109, -1000, 94, -1000, 54,
	-1000, -1000, -1000, -1000, -1000, -1000, 92, -1000, 36, -1000,
	-1000, 17, -1000, 19, 382, -1000, -53, -47, -1000, 382,
	382, 382, 382, 382, 382, 382, 382, 382, 382, 382,
	382, 382, 382, -1000, 25, -1000, 47, -1000, 83, 346,
	346, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 8, -5, -1000, -1000, -1000,
	20, 20, 346, 109, 109, 109, 109, 346, 346, 346,
	346, 346, 346, 346, 346, 346, 346, 346, 346, 346,
	346, 11, -1000, -1000, 382, 14, 14, 7, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 21, -1000, 346, -1000, -1000,
	-1000, 14, -1000, -1000,
}
var yyPgo = [...]int{

	0, 51, 210, 16, 206, 169, 7, 205, 204, 203,
	202, 95, 119, 9, 3, 198, 6, 197, 194, 1,
	193, 5, 4, 191, 190, 189, 188, 187, 185, 113,
	182, 181, 180, 179, 0, 28, 177, 176, 174, 172,
	23, 171, 2, 127, 126, 124, 124,
}
var yyR1 = [...]int{

	0, 44, 44, 44, 44, 44, 44, 44, 44, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 24, 8, 8, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 35,
	37, 37, 38, 38, 38, 36, 36, 36, 16, 16,
	23, 46, 27, 28, 41, 1, 1, 1, 2, 2,
	2, 3, 3, 3, 3, 4, 4, 4, 4, 12,
	12, 5, 5, 11, 11, 11, 11, 10, 10, 10,
	13, 13, 13, 13, 14, 14, 14, 14, 15, 15,
	15, 6, 6, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 26, 29, 29, 29,
	30, 31, 43, 43, 42, 32, 39, 39, 33, 33,
	33, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 40, 40, 45, 17, 17, 17, 17, 18,
	18, 18, 18, 18, 19, 21, 21, 20, 20, 20,
	22,
}
var yyR2 = [...]int{

	0, 2, 2, 2, 2, 2, 2, 2, 1, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 1, 1, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 1,
	0, 1, 1, 3, 3, 1, 3, 3, 0, 1,
	3, 1, 1, 1, 1, 3, 4, 2, 3, 1,
	2, 3, 3, 2, 1, 1, 1, 1, 1, 2,
	1, 1, 1, 3, 4, 2, 0, 3, 1, 2,
	3, 3, 2, 1, 3, 4, 2, 1, 3, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 2, 1, 1,
	4, 6, 0, 1, 1, 4, 3, 1, 3, 3,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 0, 3, 2, 1, 1,
	3, 1, 3, 4, 1, 2, 2, 1, 1, 1,
	1,
}
var yyChk = [...]int{

	-1000, -44, 68, 69, 70, 71, 72, 73, 2, 5,
	-1, 11, -12, -5, -11, 8, 7, 11, -14, 9,
	2, -45, -12, -34, -23, -24, -25, -26, -27, -28,
	-29, -30, -31, -32, -33, 9, -8, -22, -41, -5,
	-1, 7, -9, 27, 26, 20, 19, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55, -29, -2,
	12, -3, 7, 2, -11, -10, 12, -13, 7, 2,
	-15, 10, -6, -7, 2, 7, 8, 31, 32, 33,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 58, 59, 61, 62, 63, 64, 65, -17, 2,
	42, 28, 30, 29, 27, 26, 34, 35, 36, 37,
	38, 39, 31, 33, 32, 58, 13, 13, -34, -34,
	-1, 9, -40, -34, 59, 60, 12, 15, 2, -4,
	2, 34, 35, 40, 41, 12, 15, 2, 34, 2,
	10, 15, 2, 24, -35, -36, -38, -37, 65, -35,
	-35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, -35, 21, -42, 21, -42, 10, -39, -34,
	-34, -40, -14, -14, 12, -3, 19, 2, 12, -13,
	19, 2, 10, -6, -18, 22, -20, 7, -22, -21,
	27, 26, -34, 63, 64, 62, 61, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, 17, 14, 10, 15, 23, 23, -21, -22, -22,
	-16, -14, -16, -14, -14, -43, -42, -34, -19, 20,
	-19, 23, 14, -19,
}
var yyDef = [...]int{

	0, -2, 0, 76, 0, 76, 9, 0, 8, 7,
	1, 0, 2, 76, 70, 71, 72, 0, 3, 0,
	87, 4, -2, 5, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 9, 9, 52, 53, 118,
	119, 72, 9, 22, 23, 160, 54, 131, 132, 133,
	134, 135, 136, 137, 138, 139, 140, 141, 6, 0,
	57, 59, 0, 64, 69, 0, 75, 78, 0, 83,
	0, 86, 89, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 144, 148,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 0, 0, 0, 0, 21,
	117, 9, 9, 130, 0, 0, 55, 0, 60, 0,
	63, 65, 66, 67, 68, 73, 0, 79, 0, 82,
	84, 0, 90, 147, 9, 39, 45, 42, 41, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 116, 0, 124, 0, 50, 0, 127,
	128, 129, 142, 143, 56, 58, 61, 62, 74, 77,
	80, 81, 85, 88, 146, 149, 151, 157, 158, 159,
	0, 0, 24, -2, -2, 0, 0, 25, 26, 27,
	28, 29, 30, 31, 32, 33, 34, 35, 36, 37,
	38, 122, 120, 125, 9, 0, 0, 0, 155, 156,
	46, 49, 47, 43, 44, 0, 123, 126, 150, 154,
	152, 0, 121, 153,
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
	72, 73, 74,
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
//line promql/generated_parser.y:146
		{
			yylex.(*parser).generatedParserResult.(*VectorSelector).LabelMatchers = yyDollar[2].matchers
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:148
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].labels
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:150
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].strings
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:153
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].node
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:155
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].node
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:158
		{
			yylex.(*parser).unexpected("", "")
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:163
		{
			yylex.(*parser).errorf("No expression found in input")
			yyVAL.node = nil
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:179
		{
			if nl, ok := yyDollar[2].node.(*NumberLiteral); ok {
				if yyDollar[1].item.Typ == SUB {
					nl.Val *= -1
				}
				yyVAL.node = nl
			} else {
				yyVAL.node = &UnaryExpr{Op: yyDollar[1].item.Typ, Expr: yyDollar[2].node.(Expr)}
			}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:198
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:200
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:202
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:204
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:206
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:208
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:210
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:212
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:214
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:216
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:218
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:220
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:222
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:224
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:226
		{
			yyVAL.node = yylex.(*parser).newBinaryExpression(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].node)
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:238
		{
			yyVAL.node = &BinaryExpr{
				VectorMatching: &VectorMatching{Card: CardOneToOne},
			}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:243
		{
			yyVAL.node = &BinaryExpr{
				VectorMatching: &VectorMatching{Card: CardOneToOne},
				ReturnBool:     true,
			}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:253
		{
			yyVAL.node = yyDollar[1].node
			yyVAL.node.(*BinaryExpr).VectorMatching.MatchingLabels = yyDollar[3].strings
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:258
		{
			yyVAL.node = yyDollar[1].node
			yyVAL.node.(*BinaryExpr).VectorMatching.MatchingLabels = yyDollar[3].strings
			yyVAL.node.(*BinaryExpr).VectorMatching.On = true
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:268
		{
			yyVAL.node = yyDollar[1].node
			yyVAL.node.(*BinaryExpr).VectorMatching.Card = CardManyToOne
			yyVAL.node.(*BinaryExpr).VectorMatching.Include = yyDollar[3].strings
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:274
		{
			yyVAL.node = yyDollar[1].node
			yyVAL.node.(*BinaryExpr).VectorMatching.Card = CardOneToMany
			yyVAL.node.(*BinaryExpr).VectorMatching.Include = yyDollar[3].strings
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:283
		{
			yyVAL.strings = nil
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:288
		{
			yyVAL.node = &ParenExpr{Expr: yyDollar[2].node.(Expr)}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:298
		{
			yyVAL.node = &NumberLiteral{yyDollar[1].float}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:303
		{
			yyVAL.node = &StringLiteral{yyDollar[1].string}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:308
		{
			yyVAL.string = yylex.(*parser).unquoteString(yyDollar[1].item.Val)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:313
		{
			yyVAL.matchers = yyDollar[2].matchers
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:315
		{
			yyVAL.matchers = yyDollar[2].matchers
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:317
		{
			yyVAL.matchers = []*labels.Matcher{}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:323
		{
			yyVAL.matchers = append(yyDollar[1].matchers, yyDollar[3].matcher)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:325
		{
			yyVAL.matchers = []*labels.Matcher{yyDollar[1].matcher}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:327
		{
			yylex.(*parser).unexpected("label matching", "\",\" or \"}\"")
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:332
		{
			yyVAL.matcher = yylex.(*parser).newLabelMatcher(yyDollar[1].item, yyDollar[2].item, yyDollar[3].item)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:334
		{
			yylex.(*parser).unexpected("label matching", "string")
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:336
		{
			yylex.(*parser).unexpected("label matching", "label matching operator")
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:338
		{
			yylex.(*parser).unexpected("label matching", "identifier or \"}\"")
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:342
		{
			yyVAL.item = yyDollar[1].item
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:343
		{
			yyVAL.item = yyDollar[1].item
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:344
		{
			yyVAL.item = yyDollar[1].item
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:345
		{
			yyVAL.item = yyDollar[1].item
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:351
		{
			yyVAL.labels = append(yyDollar[2].labels, labels.Label{Name: labels.MetricName, Value: yyDollar[1].item.Val})
			sort.Sort(yyVAL.labels)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:353
		{
			yyVAL.labels = yyDollar[1].labels
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:360
		{
			yyVAL.item = yyDollar[1].item
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:361
		{
			yyVAL.item = yyDollar[1].item
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:365
		{
			yyVAL.labels = labels.New(yyDollar[2].labels...)
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:367
		{
			yyVAL.labels = labels.New(yyDollar[2].labels...)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:369
		{
			yyVAL.labels = labels.New()
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:371
		{
			yyVAL.labels = labels.New()
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:376
		{
			yyVAL.labels = append(yyDollar[1].labels, yyDollar[3].label)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:378
		{
			yyVAL.labels = []labels.Label{yyDollar[1].label}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:380
		{
			yylex.(*parser).unexpected("label set", "\",\" or \"}\"")
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:386
		{
			yyVAL.label = labels.Label{Name: yyDollar[1].item.Val, Value: yylex.(*parser).unquoteString(yyDollar[3].item.Val)}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:388
		{
			yylex.(*parser).unexpected("label set", "string")
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:390
		{
			yylex.(*parser).unexpected("label set", "\"=\"")
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:392
		{
			yylex.(*parser).unexpected("label set", "identifier or \"}\"")
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:397
		{
			yyVAL.strings = yyDollar[2].strings
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:399
		{
			yyVAL.strings = yyDollar[2].strings
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:401
		{
			yyVAL.strings = []string{}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:403
		{
			yylex.(*parser).unexpected("grouping opts", "\"(\"")
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:409
		{
			yyVAL.strings = append(yyDollar[1].strings, yyDollar[3].item.Val)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:411
		{
			yyVAL.strings = []string{yyDollar[1].item.Val}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:413
		{
			yylex.(*parser).unexpected("grouping opts", "\",\" or \"}\"")
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:418
		{
			if !isLabel(yyDollar[1].item.Val) {
				yylex.(*parser).unexpected("grouping opts", "label")
			}
			yyVAL.item = yyDollar[1].item
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:425
		{
			yylex.(*parser).unexpected("grouping opts", "label")
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:458
		{
			offset, err := parseDuration(yyDollar[3].item.Val)
			if err != nil {
				yylex.(*parser).error(err)
			}
			yylex.(*parser).addOffset(yyDollar[1].node, offset)
			yyVAL.node = yyDollar[1].node
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:470
		{
			yyVAL.node = &VectorSelector{Name: yyDollar[1].item.Val, LabelMatchers: yyDollar[2].matchers}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:472
		{
			yyVAL.node = &VectorSelector{Name: yyDollar[1].item.Val}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:474
		{
			yyVAL.node = &VectorSelector{LabelMatchers: yyDollar[1].matchers}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:479
		{
			yyVAL.node = yyDollar[1].node
			yyVAL.node.(*MatrixSelector).Range = yyDollar[3].duration
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
//line promql/generated_parser.y:487
		{
			yyVAL.node = &SubqueryExpr{
				Expr:  yyDollar[1].node.(Expr),
				Range: yyDollar[3].duration,
				Step:  yyDollar[5].duration,
			}
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:499
		{
			yyVAL.duration = 0
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:505
		{
			var err error
			yyVAL.duration, err = parseDuration(yyDollar[1].item.Val)
			if err != nil {
				yylex.(*parser).error(err)
			}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:516
		{
			fn, exist := getFunction(yyDollar[1].item.Val)
			if !exist {
				yylex.(*parser).errorf("unknown function with name %q", yyDollar[1].item.Val)
			}
			yyVAL.node = &Call{
				Func: fn,
				Args: yyDollar[3].node.(Expressions),
			}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:530
		{
			yyVAL.node = append(yyDollar[1].node.(Expressions), yyDollar[3].node.(Expr))
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:532
		{
			yyVAL.node = Expressions{yyDollar[1].node.(Expr)}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:539
		{
			yyVAL.node = yyDollar[2].node
			yyVAL.node.(*AggregateExpr).Op = yyDollar[1].item.Typ
			yyVAL.node.(*AggregateExpr).Expr = yyDollar[3].node.(Expr)
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:545
		{
			yyVAL.node = yyDollar[3].node
			yyVAL.node.(*AggregateExpr).Op = yyDollar[1].item.Typ
			yyVAL.node.(*AggregateExpr).Expr = yyDollar[2].node.(Expr)
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:551
		{
			yyVAL.node = &AggregateExpr{
				Op:   yyDollar[1].item.Typ,
				Expr: yyDollar[2].node.(Expr),
			}
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:575
		{
			yyVAL.node = &AggregateExpr{
				Grouping: yyDollar[2].strings,
			}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:581
		{
			yyVAL.node = &AggregateExpr{
				Grouping: yyDollar[2].strings,
				Without:  true,
			}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:591
		{
			yylex.(*parser).generatedParserResult = &seriesDescription{
				labels: yyDollar[1].labels,
				values: yyDollar[2].series,
			}
		}
	case 145:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:601
		{
			yyVAL.series = []sequenceValue{}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:603
		{
			yyVAL.series = append(yyDollar[1].series, yyDollar[3].series...)
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:605
		{
			yyVAL.series = yyDollar[1].series
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:607
		{
			yylex.(*parser).unexpected("series values", "")
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:612
		{
			yyVAL.series = []sequenceValue{{omitted: true}}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:614
		{
			yyVAL.series = []sequenceValue{}
			for i := uint64(0); i < yyDollar[3].uint; i++ {
				yyVAL.series = append(yyVAL.series, sequenceValue{omitted: true})
			}
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:621
		{
			yyVAL.series = []sequenceValue{{value: yyDollar[1].float}}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:623
		{
			yyVAL.series = []sequenceValue{}
			for i := uint64(0); i <= yyDollar[3].uint; i++ {
				yyVAL.series = append(yyVAL.series, sequenceValue{value: yyDollar[1].float})
			}
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:630
		{
			yyVAL.series = []sequenceValue{}
			for i := uint64(0); i <= yyDollar[4].uint; i++ {
				yyVAL.series = append(yyVAL.series, sequenceValue{value: yyDollar[1].float})
				yyDollar[1].float += yyDollar[2].float
			}
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:639
		{
			var err error
			yyVAL.uint, err = strconv.ParseUint(yyDollar[1].item.Val, 10, 64)
			if err != nil {
				yylex.(*parser).errorf("invalid repitition in series values: %s", err)
			}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:650
		{
			yyVAL.float = yyDollar[2].float
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:652
		{
			yyVAL.float = -yyDollar[2].float
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:657
		{
			if yyDollar[1].item.Val != "stale" {
				yylex.(*parser).unexpected("series values", "number or \"stale\"")
			}
			yyVAL.float = math.Float64frombits(value.StaleNaN)
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:664
		{
			yyVAL.float = yyDollar[1].float
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:666
		{
			yyVAL.float = yyDollar[1].float
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:673
		{
			yyVAL.float = yylex.(*parser).number(yyDollar[1].item.Val)
		}
	}
	goto yystack /* stack new state and value */
}
