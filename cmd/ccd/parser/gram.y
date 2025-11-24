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
	pass bool
}

%type <node> top_level_stmt
%type <node> stmt
%type <node> create_set_stmt
%type <node> help_stmt
%type <node> list_stmt
%type <node> ping_stmt

%token CREATE
%token HELP
%token LIST
%token PING
%token SET

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
	| help_stmt
		{
			$$ = $1
		}
	| list_stmt
		{
			$$ = $1
		}
	| ping_stmt
		{
			$$ = $1
		}
	| IDENT
		{
			yylex.(*lexer).pass = true
			// $$ = nil
		}

help_stmt:
	HELP
		{
			$$ = &ast.HelpStmt{}
		}

create_set_stmt:
	CREATE SET name
		{
			$$ = &ast.CreateSetStmt{SetName: $3}
		}

list_stmt:
	LIST name
		{
			$$ = &ast.ListStmt{Name: $2}
		}

ping_stmt:
	PING
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
