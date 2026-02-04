%{
package parser

import (
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

%}

%union{
	str string
	optlist []ast.Option
	node ast.Node
	selectExpr ast.SelectExpr
	queryExpr *ast.QueryExpr
	whereExpr ast.WhereExpr
	orderExpr ast.OrderExpr
	limitExpr ast.LimitExpr
	pass bool
}

%type <node> top_level_stmt
%type <node> stmt
%type <node> create_set_stmt
%type <node> info_stmt
%type <node> ping_stmt
%type <node> insert_stmt
%type <node> select_stmt
%type <node> show_stmt

%type <selectExpr> select_expression
%type <queryExpr> query_expression
%type <whereExpr> where_expression
%type <orderExpr> order_expression
%type <limitExpr> limit_expression

%token ASC
%token BY
%token CREATE
%token DESC
%token FROM
%token INFO
%token INSERT
%token INTO
%token LIMIT
%token ORDER
%token PING
%token RETRIEVE
%token SELECT
%token SET
%token SHOW
%token WHERE

%type <str> name
%type <str> unreserved_keyword

%token <str> VERSION
%token <str> IDENT NUMBER
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

info_stmt:
	INFO ';'
		{
			$$ = &ast.InfoStmt{Topic: ""}
		}
	| INFO SLITERAL ';'
		{
			$$ = &ast.InfoStmt{Topic: $2}
		}

create_set_stmt:
	CREATE SET name ';'
		{
			$$ = &ast.CreateSetStmt{SetName: $3}
		}

insert_stmt:
	INSERT INTO name query_expression ';'
		{
			$$ = &ast.InsertStmt{Into: $3, Query: $4}
		}

select_stmt:
	SELECT select_expression query_expression ';'
		{
			$$ = &ast.SelectStmt{Select: $2, Query: $3}
		}

query_expression:
	FROM name where_expression order_expression limit_expression
		{
			$$ = &ast.QueryExpr{From: $2, Where: $3, Order: $4, Limit: $5}
		}

where_expression:
	WHERE name '=' SLITERAL
		{
			$$ = &ast.WhereConditionExpr{WhereAttr: $2, WhereValue: $4}
		}
	|
		{
			$$ = &ast.NoWhereExpr{}
		}

order_expression:
	ORDER BY name
		{
			$$ = &ast.OrderValueExpr{Attribute: $3, Desc: false}
		}
	| ORDER BY name ASC
		{
			$$ = &ast.OrderValueExpr{Attribute: $3, Desc: false}
		}
	| ORDER BY name DESC
		{
			$$ = &ast.OrderValueExpr{Attribute: $3, Desc: true}
		}
	|
		{
			$$ = &ast.NoOrderExpr{}
		}

limit_expression:
	LIMIT NUMBER
		{
			$$ = &ast.LimitValueExpr{Value: $2}
		}
	|
		{
			$$ = &ast.NoLimitExpr{}
		}

select_expression:
	name
		{
			$$ = &ast.AttrSelectExpr{Attribute: $1}
		}
	| '*'
		{
			$$ = &ast.StarSelectExpr{}
		}

show_stmt:
	SHOW name ';'
		{
			$$ = &ast.ShowStmt{Name: $2}
		}

ping_stmt:
	PING ';'
		{
			$$ = &ast.PingStmt{}
		}

name:
	IDENT
		{
			$$ = $1
		}
	| unreserved_keyword
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

unreserved_keyword:
	VERSION
