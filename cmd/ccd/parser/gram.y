%{
package parser

import (
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

%}

%union{
	node ast.Node
	str string
	nodeList []ast.Node
}

%type <node> top_level_stmt
%type <node> stmt
%type <node> create_set_stmt
%type <node> info_stmt
%type <node> ping_stmt
%type <node> insert_stmt
%type <node> select_stmt
%type <node> show_stmt

%type <node> select_attr_list
%type <node> query_clause
%type <node> where_clause
%type <node> order_clause
%type <node> limit_clause

%type <node> expression
%type <node> logical_or_expr
%type <node> logical_and_expr
%type <node> equality_expr
%type <node> relational_expr
%type <node> unary_expr
%type <node> postfix_expr
%type <node> primary_expr
%type <nodeList> arg_expr_list
%type <nodeList> arg_expr

%token GT_OR_EQUAL
%token LT_OR_EQUAL
%token NOT_EQUAL

%token AND
%token ASC
%token BY
%token CREATE
%token DESC
%token FILTER
%token FROM
%token INFO
%token INSERT
%token INTO
%token LIMIT
%token NOT
%token OR
%token ORDER
%token PING
%token RETRIEVE
%token SELECT
%token SET
%token SHOW
%token TAG
%token WHERE

%type <str> name

%token <str> IDENT
%token <str> NUMBER
%token <str> SLITERAL

%start main

%%

main:
	top_level_stmt
		{
			yylex.(*lexer).node = $1
		}

top_level_stmt:
	stmt
		{
			$$ = $1
		}

stmt:
	create_set_stmt
		{
			$$ = $1
		}
	| info_stmt
		{
			$$ = $1
		}
	| insert_stmt
		{
			$$ = $1
		}
	| select_stmt
		{
			$$ = $1
		}
	| show_stmt
		{
			$$ = $1
		}
	| ping_stmt
		{
			$$ = $1
		}

create_set_stmt:
	CREATE SET name ';'
		{
			$$ = &ast.CreateSetStmt{SetName: $3}
		}

info_stmt:
	INFO ';'
		{
			$$ = &ast.InfoStmt{Topic: ""}
		}
	| INFO SLITERAL ';'
		{
			$$ = &ast.InfoStmt{Topic: $2}
		}

insert_stmt:
	INSERT INTO name query_clause ';'
		{
			$$ = &ast.InsertStmt{Into: $3, Query: $4}
		}

ping_stmt:
	PING ';'
		{
			$$ = &ast.PingStmt{}
		}

select_stmt:
	SELECT select_attr_list query_clause ';'
		{
			$$ = &ast.SelectStmt{AttrList: $2, Query: $3}
		}

show_stmt:
	SHOW name ';'
		{
			$$ = &ast.ShowStmt{Name: $2}
		}

select_attr_list:
	name
		{
			$$ = &ast.SelectAttrList{Attr: $1}
		}
	| '*'
		{
			$$ = &ast.SelectAttrList{Attr: "*"}
		}

query_clause:
	FROM name where_clause order_clause limit_clause
		{
			$$ = &ast.QueryClause{From: $2, Where: $3, Order: $4, Limit: $5}
		}

where_clause:
	WHERE expression
		{
			$$ = &ast.WhereClause{Valid: true, Condition: $2}
		}
	|
		{
			$$ = &ast.WhereClause{}
		}

order_clause:
	ORDER BY name
		{
			$$ = &ast.OrderClause{Valid: true, Attr: $3, Desc: false}
		}
	| ORDER BY name ASC
		{
			$$ = &ast.OrderClause{Valid: true, Attr: $3, Desc: false}
		}
	| ORDER BY name DESC
		{
			$$ = &ast.OrderClause{Valid: true, Attr: $3, Desc: true}
		}
	|
		{
			$$ = &ast.OrderClause{}
		}

limit_clause:
	LIMIT NUMBER
		{
			$$ = &ast.LimitClause{Valid: true, Value: $2}
		}
	|
		{
			$$ = &ast.LimitClause{}
		}

expression:
	logical_or_expr
		{
			$$ = $1
		}

logical_or_expr:
	logical_and_expr
		{
			$$ = $1
		}
	| logical_or_expr OR logical_and_expr
		{
			$$ = &ast.OrExpr{Expr1: $1, Expr2: $3}
		}

logical_and_expr:
	unary_expr
		{
			$$ = $1
		}
	| logical_and_expr AND unary_expr
		{
			$$ = &ast.AndExpr{Expr1: $1, Expr2: $3}
		}

unary_expr:
	equality_expr
		{
			$$ = $1
		}
	| NOT unary_expr
		{
			$$ = &ast.NotExpr{Expr: $2}
		}

equality_expr:
	relational_expr
		{
			$$ = $1
		}
	| equality_expr '=' relational_expr
		{
			$$ = &ast.EqualExpr{Expr1: $1, Expr2: $3}
		}
	| equality_expr NOT_EQUAL relational_expr
		{
			$$ = &ast.NotEqualExpr{Expr1: $1, Expr2: $3}
		}

relational_expr:
	postfix_expr
		{
			$$ = $1
		}
	| relational_expr '<' postfix_expr
		{
			$$ = &ast.LessThanExpr{Expr1: $1, Expr2: $3}
		}
	| relational_expr '>' postfix_expr
		{
			$$ = &ast.GreaterThanExpr{Expr1: $1, Expr2: $3}
		}
	| relational_expr LT_OR_EQUAL postfix_expr
		{
			$$ = &ast.LessThanOrEqualExpr{Expr1: $1, Expr2: $3}
		}
	| relational_expr GT_OR_EQUAL postfix_expr
		{
			$$ = &ast.GreaterThanOrEqualExpr{Expr1: $1, Expr2: $3}
		}

postfix_expr:
	primary_expr
		{
			$$ = $1
		}
	| FILTER '(' arg_expr_list ')'
		{
			$$ = &ast.FilterExpr{ExprList: $3}
		}
	| TAG '(' arg_expr_list ')'
		{
			$$ = &ast.TagExpr{ExprList: $3}
		}

primary_expr:
	name
		{
			$$ = &ast.Name{Value: $1}
		}
	| SLITERAL
		{
			$$ = &ast.SLiteral{Value: $1}
		}
	| NUMBER
		{
			$$ = &ast.Number{Value: $1}
		}
	| '(' expression ')'
		{
			$$ = &ast.ParenExpr{Expr: $2}
		}

arg_expr_list:
	arg_expr
		{
			$$ = $1
		}
	| arg_expr_list ',' arg_expr
		{
			$$ = append($1, $3...)
		}

arg_expr:
	name
		{
			$$ = []ast.Node{&ast.Name{Value: $1}}
		}

name:
	IDENT
		{
			$$ = $1
		}

/*
boolean:
	TRUE
		{
			$$ = "true"
		}
	| FALSE
		{
			$$ = "false"
		}
*/
