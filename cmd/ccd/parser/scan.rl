package parser

import (
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

%%{ 
	machine sql;
	write data;
	access lex.;
	variable p lex.p;
	variable pe lex.pe;

	identifier = [A-Za-z_][0-9A-Za-z_.]*;
	sliteral = ['][^']*['];
}%%

type lexer struct {
	data []byte
	p, pe, cs int
	ts, te, act int

	err string
	str string
	optlist []ast.Option
	node ast.Node
	pass bool
}

func newLexer(data []byte) *lexer {
	lex := &lexer{ 
		data: data,
		pe: len(data),
	}
	%% write init;
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0
	%%{ 
		main := |*
			';' => { tok = ';'; fbreak; };
			',' => { tok = ','; fbreak; };
			'(' => { tok = '('; fbreak; };
			')' => { tok = ')'; fbreak; };
			'*' => { tok = '*'; fbreak; };
			'=' => { tok = '='; fbreak; };
			'asc'i => { tok = ASC; fbreak; };
			'by'i => { tok = BY; fbreak; };
			'create'i => { tok = CREATE; fbreak; };
			'desc'i => { tok = DESC; fbreak; };
			'from'i => { tok = FROM; fbreak; };
			'info'i => { tok = INFO; fbreak; };
			'insert'i => { tok = INSERT; fbreak; };
			'into'i => { tok = INTO; fbreak; };
			'limit'i => { tok = LIMIT; fbreak; };
			'order'i => { tok = ORDER; fbreak; };
			'retrieve'i => { tok = RETRIEVE; fbreak; };
			'set'i => { tok = SET; fbreak; };
			'show'i => { tok = SHOW; fbreak; };
			'ping'i => { tok = PING; fbreak; };
			'version'i => { out.str = "version"; tok = VERSION; fbreak; };
			'select'i => { tok = SELECT; fbreak; };
			'where'i => { tok = WHERE; fbreak; };
			identifier => { out.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak; };
			sliteral => { out.str = string(lex.data[lex.ts+1:lex.te-1]); tok = SLITERAL; fbreak; };
			digit+ => { out.str = string(lex.data[lex.ts:lex.te]); tok = NUMBER; fbreak; };
			space;
		*|;

		write exec;
	}%%

	return tok;
}

func (lex *lexer) Error(e string) {
	lex.err = e
}
