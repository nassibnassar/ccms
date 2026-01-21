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
%type <node> ping_stmt
%type <node> retrieve_stmt
%type <node> show_filters_stmt
%type <node> show_sets_stmt

%token CREATE
%token FILTERS
%token FROM
%token HELP
%token LIMIT
%token PING
%token RETRIEVE
%token SET
%token SETS
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
	| help_stmt
		{
			$$ = $1
		}
	| retrieve_stmt
		{
			$$ = $1
		}
	| show_filters_stmt
		{
			$$ = $1
		}
	| show_sets_stmt
		{
			$$ = $1
		}
	| ping_stmt
		{
			$$ = $1
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

retrieve_stmt:
	RETRIEVE name FROM name LIMIT NUMBER
		{
			$$ = &ast.RetrieveStmt{Attribute: $2, Set: $4, Limit: $6}
		}
	| RETRIEVE name FROM name
		{
			$$ = &ast.RetrieveStmt{Attribute: $2, Set: $4, Limit: "20"}
		}

show_filters_stmt:
	SHOW FILTERS
		{
			$$ = &ast.ShowFiltersStmt{}
		}

show_sets_stmt:
	SHOW SETS
		{
			$$ = &ast.ShowSetsStmt{}
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
