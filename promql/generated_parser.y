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

import (
        "math"
        "sort"
        "strconv"
        "time"

        "github.com/prometheus/prometheus/pkg/labels"
        "github.com/prometheus/prometheus/pkg/value"
)
%}

%union {
    node      Node
    item      Item
    matchers  []*labels.Matcher
    matcher   *labels.Matcher
    label     labels.Label
    labels    labels.Labels
    strings   []string
    series    []sequenceValue
    uint      uint64
    float     float64
    string    string
    duration  time.Duration
}


%token <item> ASSIGN BLANK COLON COMMA COMMENT DURATION EOF ERROR IDENTIFIER LEFT_BRACE LEFT_BRACKET LEFT_PAREN METRIC_IDENTIFIER NUMBER RIGHT_BRACE RIGHT_BRACKET RIGHT_PAREN SEMICOLON SPACE STRING TIMES

// Operators.
%token	operatorsStart
%token <item> ADD DIV EQL EQL_REGEX GTE GTR LAND LOR LSS LTE LUNLESS MOD MUL NEQ NEQ_REGEX POW SUB
%token	operatorsEnd

// Aggregators.
%token	aggregatorsStart
%token <item> AVG BOTTOMK COUNT COUNT_VALUES MAX MIN QUANTILE STDDEV STDVAR SUM TOPK
%token	aggregatorsEnd

// Keywords.
%token	keywordsStart
%token <item> BOOL BY GROUP_LEFT GROUP_RIGHT IGNORING OFFSET ON WITHOUT
%token keywordsEnd


// Start symbols for the generated parser.
%token	startSymbolsStart
%token START_METRIC START_SERIES_DESCRIPTION START_EXPRESSION START_METRIC_SELECTOR
%token	startSymbolsEnd


// Type definitions for grammar rules.
%type <matchers> label_match_list label_matchers
%type <matcher> label_matcher

%type <item> aggregate_op grouping_label match_op maybe_label metric_identifier unary_op

%type <labels> label_set label_set_list metric
%type <label> label_set_item
%type <strings> grouping_label_list grouping_labels maybe_grouping_labels
%type <series> series_item series_values
%type <uint> uint
%type <float> number series_value signed_number
%type <node> aggregate_expr aggregate_modifier bin_modifier binary_expr bool_modifier expr function_call function_call_args function_call_body group_modifiers matrix_selector number_literal offset_expr on_or_ignoring paren_expr string_literal subquery_expr unary_expr vector_selector
%type <string> string
%type <duration> duration maybe_duration

%start start

// Operators are listed with increasing precedence.
%left LOR
%left LAND LUNLESS
%left EQL GTE GTR LSS LTE NEQ
%left ADD SUB
%left MUL DIV MOD
%right POW

// Offset modifiers do not have associativity.
%nonassoc OFFSET

// This ensures that it is always attempted to parse range or subquery selectors when a left
// bracket is encountered.
%right LEFT_BRACKET

%%

start           : 
                START_METRIC metric
                     { yylex.(*parser).generatedParserResult = $2 }
                | START_SERIES_DESCRIPTION series_description
                | START_EXPRESSION /* empty */ EOF
                        { yylex.(*parser).errorf("no expression found in input")}
                | START_EXPRESSION expr
                     { yylex.(*parser).generatedParserResult = $2 }
                | START_METRIC_SELECTOR vector_selector 
                     { yylex.(*parser).generatedParserResult = $2 }
                | start EOF
                | error /* If none of the more detailed error messages are triggered, we fall back to this. */
                        { yylex.(*parser).unexpected("","") }
                ;

expr            :
                paren_expr
                | binary_expr
                | unary_expr
                | offset_expr
                | number_literal
                | string_literal
                | vector_selector
                | matrix_selector
                | subquery_expr
                | function_call
                | aggregate_expr
                ;

unary_expr      :
                /* gives the rule the same prec as POW, not a perfect solution */
                unary_op expr %prec POW
                        {
                        if nl, ok := $2.(*NumberLiteral); ok {
                                if $1.Typ == SUB {
                                        nl.Val *= -1
                                }
                                $$ = nl
                        } else {
                                $$ = &UnaryExpr{Op: $1.Typ, Expr: $2.(Expr)}
                        }
                        }
                ;

unary_op        :
                ADD
                | SUB
                ;

binary_expr     :
                expr POW bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr MUL bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr DIV bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr MOD bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr ADD bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr SUB bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr EQL bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr NEQ bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr LTE bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr LSS bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr GTE bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr GTR bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr LAND bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr LUNLESS bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                | expr LOR bin_modifier expr
                        { $$ = yylex.(*parser).newBinaryExpression($1, $2, $3, $4)} 
                ;

// Using left recursion for the modifier rules, helps to keep the parser stack small and 
// reduces allocations

bin_modifier    :
                group_modifiers
                ;

bool_modifier   :
                /* empty */
                        { $$ = &BinaryExpr{
                        VectorMatching: &VectorMatching{Card: CardOneToOne},
                        }
                        }
                | BOOL
                        { $$ = &BinaryExpr{
                        VectorMatching: &VectorMatching{Card: CardOneToOne},
                        ReturnBool:     true,
                        }
                        }
                ;

on_or_ignoring  :
                bool_modifier IGNORING grouping_labels
                        {
                        $$ = $1
                        $$.(*BinaryExpr).VectorMatching.MatchingLabels = $3
                        } 
                | bool_modifier ON grouping_labels
                        {
                        $$ = $1
                        $$.(*BinaryExpr).VectorMatching.MatchingLabels = $3
                        $$.(*BinaryExpr).VectorMatching.On = true
                        } 
                ;

group_modifiers:
                bool_modifier /* empty */
                | on_or_ignoring /* empty */
                | on_or_ignoring GROUP_LEFT maybe_grouping_labels
                        {
                        $$ = $1
                        $$.(*BinaryExpr).VectorMatching.Card = CardManyToOne
                        $$.(*BinaryExpr).VectorMatching.Include = $3
                        }
                | on_or_ignoring GROUP_RIGHT maybe_grouping_labels
                        {
                        $$ = $1
                        $$.(*BinaryExpr).VectorMatching.Card = CardOneToMany
                        $$.(*BinaryExpr).VectorMatching.Include = $3
                        }
                ;

maybe_grouping_labels:
                /* empty */
                        { $$ = nil }
                | grouping_labels
                ;

paren_expr      : LEFT_PAREN expr RIGHT_PAREN
                        { $$ = &ParenExpr{Expr: $2.(Expr)} }
                ;


number_literal  :
                number
                        { $$ = &NumberLiteral{$1}}
                ;

string_literal  :
                string
                        { $$ = &StringLiteral{$1}}
                ;

string          :
                STRING
                        { $$ = yylex.(*parser).unquoteString($1.Val) }
                ;

label_matchers  : 
                LEFT_BRACE label_match_list RIGHT_BRACE
                        { $$ = $2 }
                | LEFT_BRACE label_match_list COMMA RIGHT_BRACE
                        { $$ = $2 }
                | LEFT_BRACE RIGHT_BRACE
                        { $$ = []*labels.Matcher{} }

                ;

label_match_list:
                label_match_list COMMA label_matcher
                        { $$ = append($1, $3)}
                | label_matcher
                        { $$ = []*labels.Matcher{$1}}
                | label_match_list error
                        { yylex.(*parser).unexpected("label matching", "\",\" or \"}\"") }
                ;

label_matcher   :
                IDENTIFIER match_op STRING
                        { $$ = yylex.(*parser).newLabelMatcher($1, $2, $3) }
                | IDENTIFIER match_op error
                        { yylex.(*parser).unexpected("label matching", "string")}
                | IDENTIFIER error 
                        { yylex.(*parser).unexpected("label matching", "label matching operator") } 
                | error
                        { yylex.(*parser).unexpected("label matching", "identifier or \"}\"")}
                ;

match_op        :
                EQL {$$ =$1}
                | NEQ {$$=$1}
                | EQL_REGEX {$$=$1}
                | NEQ_REGEX {$$=$1}
                ;


metric          :
                metric_identifier label_set
                        { $$ = append($2, labels.Label{Name: labels.MetricName, Value: $1.Val}); sort.Sort($$) }
                | label_set 
                        {$$ = $1}
                ;


metric_identifier
                :
                METRIC_IDENTIFIER {$$=$1}
                | IDENTIFIER      {$$=$1}

label_set       :
                LEFT_BRACE label_set_list RIGHT_BRACE
                        { $$ = labels.New($2...) }
                | LEFT_BRACE label_set_list COMMA RIGHT_BRACE
                        { $$ = labels.New($2...) }
                | LEFT_BRACE RIGHT_BRACE
                        { $$ = labels.New() }
                | /* empty */
                        { $$ = labels.New() }
                ;

label_set_list  :
                label_set_list COMMA label_set_item
                        { $$ = append($1, $3) }
                | label_set_item
                        { $$ = []labels.Label{$1} }
                | label_set_list error
                        { yylex.(*parser).unexpected("label set", "\",\" or \"}\"", ) }
                
                ;

label_set_item  :
                IDENTIFIER EQL STRING
                        { $$ = labels.Label{Name: $1.Val, Value: yylex.(*parser).unquoteString($3.Val) } } 
                | IDENTIFIER EQL error
                        { yylex.(*parser).unexpected("label set", "string")}
                | IDENTIFIER error
                        { yylex.(*parser).unexpected("label set", "\"=\"")}
                | error
                        { yylex.(*parser).unexpected("label set", "identifier or \"}\"") }
                ;

grouping_labels :
                LEFT_PAREN grouping_label_list RIGHT_PAREN
                        { $$ = $2 }
                | LEFT_PAREN grouping_label_list COMMA RIGHT_PAREN
                        { $$ = $2 }
                | LEFT_PAREN RIGHT_PAREN
                        { $$ = []string{} }
                | error
                        { yylex.(*parser).unexpected("grouping opts", "\"(\"") }
                ;


grouping_label_list:
                grouping_label_list COMMA grouping_label
                        { $$ = append($1, $3.Val) }
                | grouping_label
                        { $$ = []string{$1.Val} }
                | grouping_label_list error
                        { yylex.(*parser).unexpected("grouping opts", "\",\" or \")\"") }
                ;

grouping_label  :
                maybe_label
                        {
                        if !isLabel($1.Val) {
                                yylex.(*parser).unexpected("grouping opts", "label")
                        }
                        $$ = $1
                        }
                | error
                        { yylex.(*parser).unexpected("grouping opts", "label") }
                ;


/* inside of grouping options label names can be recognized as keywords by the lexer */
maybe_label     :
                IDENTIFIER
                | METRIC_IDENTIFIER
                | LAND
                | LOR
                | LUNLESS
                | AVG
                | COUNT
                | SUM
                | MIN
                | MAX
                | STDDEV
                | STDVAR
                | TOPK
                | BOTTOMK
                | COUNT_VALUES
                | QUANTILE
                | OFFSET
                | BY
                | ON
                | IGNORING
                | GROUP_LEFT
                | GROUP_RIGHT
                | BOOL
                ;

// The series description grammar is only used inside unit tests.
offset_expr:
                expr OFFSET DURATION
                        {
                        offset, err := parseDuration($3.Val)
                        if err != nil {
                                yylex.(*parser).error(err)
                        }
                        yylex.(*parser).addOffset($1, offset)
                        $$ = $1
                        }
                | expr OFFSET error
                        { yylex.(*parser).unexpected("offset", "duration") }
                ;

vector_selector:
                metric_identifier label_matchers
                        { $$ = yylex.(*parser).newVectorSelector($1.Val, $2) }
                | metric_identifier 
                        { $$ = yylex.(*parser).newVectorSelector($1.Val, nil) }
                | label_matchers
                        { $$ = yylex.(*parser).newVectorSelector("", $1) }
                ;

matrix_selector :
                expr LEFT_BRACKET duration RIGHT_BRACKET
                        {
                        vs, ok := $1.(*VectorSelector)
                        if !ok{
                                yylex.(*parser).errorf("matrix selectors only allowed for vector selectors")
                        }
                        if vs.Offset != 0{
                                yylex.(*parser).errorf("no offset modifiers allowed before range selector")
                        }
                        $$ = &MatrixSelector{
                                Name: vs.Name,
                                Offset: vs.Offset,
                                LabelMatchers: vs.LabelMatchers,
                                Range: $3,
                        }
                        }
                ;

subquery_expr   :
                expr LEFT_BRACKET duration COLON maybe_duration RIGHT_BRACKET
                        {
                        $$ = &SubqueryExpr{
                                Expr:  $1.(Expr),
                                Range: $3,
                                Step:  $5,
                        }
                        }
                | expr LEFT_BRACKET duration COLON duration error
                        { yylex.(*parser).unexpected("subquery selector", "\"]\"") }
                | expr LEFT_BRACKET duration COLON error
                        { yylex.(*parser).unexpected("subquery selector", "duration or \"]\"") }
                | expr LEFT_BRACKET duration error
                        { yylex.(*parser).unexpected("subquery or matrix selector", "\":\" or \"]\"") }
                | expr LEFT_BRACKET error
                        { yylex.(*parser).unexpected("subquery selector", "duration") }
                ;
                        

maybe_duration  :
                /* empty */
                        {$$ = 0}
                | duration
                ;

duration        :
                DURATION
                        {
                        var err error
                        $$, err = parseDuration($1.Val)
                        if err != nil {
                                yylex.(*parser).error(err)
                        }
                        }
                ;
                
function_call   :
                IDENTIFIER function_call_body
                        {
                        fn, exist := getFunction($1.Val)
                        if !exist{
                                yylex.(*parser).errorf("unknown function with name %q", $1.Val)
                        } 
                        $$ = &Call{
                                Func: fn,
                                Args: $2.(Expressions),
                        }
                        }
                ;

function_call_body:
                LEFT_PAREN function_call_args RIGHT_PAREN
                        { $$ = $2 }
                | LEFT_PAREN RIGHT_PAREN
                        {$$ = Expressions{}}
                ;

function_call_args:
                function_call_args COMMA expr
                        { $$ = append($1.(Expressions), $3.(Expr)) }
                | expr 
                        { $$ = Expressions{$1.(Expr)} }
                ;

// TODO param support
// Consider unifying calls and aggregate expressions
aggregate_expr  :
                aggregate_op aggregate_modifier function_call_body
                        { $$ = yylex.(*parser).newAggregateExpr($1, $2, $3) }
                | aggregate_op function_call_body aggregate_modifier
                        { $$ = yylex.(*parser).newAggregateExpr($1, $3, $2) }
                | aggregate_op function_call_body
                        { $$ = yylex.(*parser).newAggregateExpr($1, &AggregateExpr{}, $2) }
                | aggregate_op error
                        { yylex.(*parser).unexpected("aggregation","") }
                ;

aggregate_op    :
                AVG
                | COUNT
                | SUM
                | MIN
                | MAX
                | STDDEV
                | STDVAR
                | TOPK
                | BOTTOMK
                | COUNT_VALUES
                | QUANTILE
                ;

aggregate_modifier:
                BY grouping_labels
                        {
                        $$ = &AggregateExpr{
                                Grouping: $2,
                        }
                        }
                | WITHOUT grouping_labels
                        {
                        $$ = &AggregateExpr{
                                Grouping: $2,
                                Without:  true,
                        }
                        }
                ;

series_description:
                metric series_values
                        {
                        yylex.(*parser).generatedParserResult = &seriesDescription{
                                labels: $1,
                                values: $2,
                        }
                        }
                ;

series_values   :
                /*empty*/
                        { $$ = []sequenceValue{} }
                | series_values SPACE series_item
                        { $$ = append($1, $3...) }
                | series_values SPACE
                        { $$ = $1 }
                | error
                        { yylex.(*parser).unexpected("series values", "") }
                ;

series_item     :
                BLANK
                        { $$ = []sequenceValue{{omitted: true}}}
                | BLANK TIMES uint
                        {
                        $$ = []sequenceValue{}
                        for i:=uint64(0); i < $3; i++{
                                $$ = append($$, sequenceValue{omitted: true})
                        }
                        }
                | series_value
                        { $$ = []sequenceValue{{value: $1}}}
                | series_value TIMES uint
                        {
                        $$ = []sequenceValue{}
                        for i:=uint64(0); i <= $3; i++{
                                $$ = append($$, sequenceValue{value: $1})
                        }
                        }
                | series_value signed_number TIMES uint
                        {
                        $$ = []sequenceValue{}
                        for i:=uint64(0); i <= $4; i++{
                                $$ = append($$, sequenceValue{value: $1})
                                $1 += $2
                        }
                        }
uint            :
                NUMBER
                        {
                        var err error
                        $$, err = strconv.ParseUint($1.Val, 10, 64)
                        if err != nil {
                                yylex.(*parser).errorf("invalid repetition in series values: %s", err)
                        }
                        }
                ;

signed_number   :
                ADD number
                        { $$ = $2 }
                | SUB number
                        { $$ = -$2 }
                ;

series_value    :
                IDENTIFIER
                        {
                        if $1.Val != "stale" {
                                yylex.(*parser).unexpected("series values", "number or \"stale\"")
                        }
                        $$ = math.Float64frombits(value.StaleNaN)
                        }
                | number
                        { $$ = $1 }
                | signed_number
                        { $$ = $1 }
                ;



number          :
                NUMBER
                        {$$ = yylex.(*parser).number($1.Val) }
                ;

%%
