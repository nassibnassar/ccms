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
	pass bool
}

%type <node> top_level_stmt
%type <node> stmt
%type <node> create_set_stmt
%type <node> info_stmt
%type <node> ping_stmt
%type <node> retrieve_stmt
%type <node> show_stmt

%type <selectExpr> select_expression

%token CREATE
%token FROM
%token INFO
%token LIMIT
%token PING
%token RETRIEVE
%token SELECT
%token SET
%token SHOW

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
	| retrieve_stmt
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

retrieve_stmt:
	SELECT '*' FROM name LIMIT NUMBER ';'
		{
			$$ = &ast.SelectStmt{Select: &ast.StarSelectExpr{}, Set: $4, Limit: $6, Retrieve: false}
		}
	| SELECT '*' FROM name ';'
		{
			$$ = &ast.SelectStmt{Select: &ast.StarSelectExpr{}, Set: $4, Limit: "1000", Retrieve: false}
		}
	| RETRIEVE select_expression FROM name LIMIT NUMBER ';'
		{
			$$ = &ast.SelectStmt{Select: $2, Set: $4, Limit: $6, Retrieve: true}
		}
	| RETRIEVE select_expression FROM name ';'
		{
			$$ = &ast.SelectStmt{Select: $2, Set: $4, Limit: "20", Retrieve: true}
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
