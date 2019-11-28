// Copyright 2019 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

%{
    package promql
%}

%union {
    Node Node
    item item
}

%token	ERROR 
%token	EOF
%token	COMMENT
%token	IDENTIFIER
%token	METRIC_IDENTIFIER
%token	LEFT_PAREN
%token	RIGHT_PAREN
%token	LEFT_BRACE
%token	RIGHT_BRACE
%token	LEFT_BRACKET
%token	RIGHT_BRACKET
%token	COMMA
%token	ASSIGN
%token	COLON
%token	SEMICOLON
%token	STRING
%token	NUMBER
%token	DURATION
%token	BLANK
%token	TIMES
%token	SPACE

%token	operatorsStart
// Operators.
%token	SUB
%token	ADD
%token	MUL
%token	MOD
%token	DIV
%token	LAND
%token	LOR
%token	LUNLESS
%token	EQL
%token	NEQ
%token	LTE
%token	LSS
%token	GTE
%token	GTR
%token	EQL_REGEX
%token	NEQ_REGEX
%token	POW
%token	operatorsEnd

%token	aggregatorsStart
// Aggregators.
%token	AVG
%token	COUNT
%token	SUM
%token	MIN
%token	MAX
%token	STDDEV
%token	STDVAR
%token	TOPK
%token	BOTTOMK
%token	COUNT_VALUES
%token	QUANTILE
%token	aggregatorsEnd
%token
%token	keywordsStart
// Keywords.
%token	OFFSET
%token	BY
%token	WITHOUT
%token	ON
%token	IGNORING
%token	GROUP_LEFT
%token	GROUP_RIGHT
%token	BOOL
%token	keywordsEnd


%token	startSymbolsStart
	// Start symbols for the generated parser.
%token	startSymbolsEnd

%start start;

%%

start:  SPACE

%%