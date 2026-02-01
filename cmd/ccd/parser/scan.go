
//line scan.rl:1
package parser

import (
	"github.com/indexdata/ccms/cmd/ccd/ast"
)


//line scan.go:9
const sql_start int = 2
const sql_first_final int = 2
const sql_error int = 0

const sql_en_main int = 2


//line scan.rl:16


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
	
//line scan.go:36
	{
	 lex.cs = sql_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line scan.rl:36
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0
	
//line scan.go:50
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 2:
		goto st_case_2
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	}
	goto st_out
tr1:
//line scan.rl:62
 lex.te = ( lex.p)+1
{ out.str = string(lex.data[lex.ts+1:lex.te-1]); tok = SLITERAL; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr2:
//line scan.rl:64
 lex.te = ( lex.p)+1

	goto st2
tr4:
//line scan.rl:46
 lex.te = ( lex.p)+1
{ tok = '('; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr5:
//line scan.rl:47
 lex.te = ( lex.p)+1
{ tok = ')'; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr6:
//line scan.rl:48
 lex.te = ( lex.p)+1
{ tok = '*'; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr7:
//line scan.rl:45
 lex.te = ( lex.p)+1
{ tok = ','; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr9:
//line scan.rl:44
 lex.te = ( lex.p)+1
{ tok = ';'; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr10:
//line scan.rl:49
 lex.te = ( lex.p)+1
{ tok = '='; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr21:
//line scan.rl:63
 lex.te = ( lex.p)
( lex.p)--
{ out.str = string(lex.data[lex.ts:lex.te]); tok = NUMBER; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr22:
//line NONE:1
	switch  lex.act {
	case 7:
	{( lex.p) = ( lex.te) - 1
 tok = CREATE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 8:
	{( lex.p) = ( lex.te) - 1
 tok = FROM; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 9:
	{( lex.p) = ( lex.te) - 1
 tok = INFO; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 10:
	{( lex.p) = ( lex.te) - 1
 tok = LIMIT; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 11:
	{( lex.p) = ( lex.te) - 1
 tok = RETRIEVE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 12:
	{( lex.p) = ( lex.te) - 1
 tok = SET; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 13:
	{( lex.p) = ( lex.te) - 1
 tok = SHOW; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 14:
	{( lex.p) = ( lex.te) - 1
 tok = PING; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 15:
	{( lex.p) = ( lex.te) - 1
 out.str = "version"; tok = VERSION; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 16:
	{( lex.p) = ( lex.te) - 1
 tok = SELECT; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 17:
	{( lex.p) = ( lex.te) - 1
 tok = WHERE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 18:
	{( lex.p) = ( lex.te) - 1
 out.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; {( lex.p)++;  lex.cs = 2; goto _out } }
	}
	
	goto st2
tr23:
//line scan.rl:61
 lex.te = ( lex.p)
( lex.p)--
{ out.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
	st2:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
//line NONE:1
 lex.ts = ( lex.p)

//line scan.go:257
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr2
		case 39:
			goto st1
		case 40:
			goto tr4
		case 41:
			goto tr5
		case 42:
			goto tr6
		case 44:
			goto tr7
		case 59:
			goto tr9
		case 61:
			goto tr10
		case 67:
			goto st5
		case 70:
			goto st10
		case 73:
			goto st13
		case 76:
			goto st16
		case 80:
			goto st20
		case 82:
			goto st23
		case 83:
			goto st30
		case 86:
			goto st37
		case 87:
			goto st43
		case 95:
			goto tr11
		case 99:
			goto st5
		case 102:
			goto st10
		case 105:
			goto st13
		case 108:
			goto st16
		case 112:
			goto st20
		case 114:
			goto st23
		case 115:
			goto st30
		case 118:
			goto st37
		case 119:
			goto st43
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr11
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr11
			}
		default:
			goto st3
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if  lex.data[( lex.p)] == 39 {
			goto tr1
		}
		goto st1
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st3
		}
		goto tr21
tr11:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:61
 lex.act = 18;
	goto st4
tr28:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:50
 lex.act = 7;
	goto st4
tr31:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:51
 lex.act = 8;
	goto st4
tr34:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:52
 lex.act = 9;
	goto st4
tr38:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:53
 lex.act = 10;
	goto st4
tr41:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:57
 lex.act = 14;
	goto st4
tr48:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:54
 lex.act = 11;
	goto st4
tr52:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:55
 lex.act = 12;
	goto st4
tr55:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:59
 lex.act = 16;
	goto st4
tr57:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:56
 lex.act = 13;
	goto st4
tr63:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:58
 lex.act = 15;
	goto st4
tr67:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:60
 lex.act = 17;
	goto st4
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
//line scan.go:443
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 95:
			goto tr11
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr22
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 82:
			goto st6
		case 95:
			goto tr11
		case 114:
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto st7
		case 95:
			goto tr11
		case 101:
			goto st7
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 65:
			goto st8
		case 95:
			goto tr11
		case 97:
			goto st8
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 84:
			goto st9
		case 95:
			goto tr11
		case 116:
			goto st9
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto tr28
		case 95:
			goto tr11
		case 101:
			goto tr28
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 82:
			goto st11
		case 95:
			goto tr11
		case 114:
			goto st11
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 79:
			goto st12
		case 95:
			goto tr11
		case 111:
			goto st12
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 77:
			goto tr31
		case 95:
			goto tr11
		case 109:
			goto tr31
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 78:
			goto st14
		case 95:
			goto tr11
		case 110:
			goto st14
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 70:
			goto st15
		case 95:
			goto tr11
		case 102:
			goto st15
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 79:
			goto tr34
		case 95:
			goto tr11
		case 111:
			goto tr34
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 73:
			goto st17
		case 95:
			goto tr11
		case 105:
			goto st17
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 77:
			goto st18
		case 95:
			goto tr11
		case 109:
			goto st18
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 73:
			goto st19
		case 95:
			goto tr11
		case 105:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 84:
			goto tr38
		case 95:
			goto tr11
		case 116:
			goto tr38
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 73:
			goto st21
		case 95:
			goto tr11
		case 105:
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 78:
			goto st22
		case 95:
			goto tr11
		case 110:
			goto st22
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 71:
			goto tr41
		case 95:
			goto tr11
		case 103:
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto st24
		case 95:
			goto tr11
		case 101:
			goto st24
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 84:
			goto st25
		case 95:
			goto tr11
		case 116:
			goto st25
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 82:
			goto st26
		case 95:
			goto tr11
		case 114:
			goto st26
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 73:
			goto st27
		case 95:
			goto tr11
		case 105:
			goto st27
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto st28
		case 95:
			goto tr11
		case 101:
			goto st28
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 86:
			goto st29
		case 95:
			goto tr11
		case 118:
			goto st29
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto tr48
		case 95:
			goto tr11
		case 101:
			goto tr48
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto st31
		case 72:
			goto st35
		case 95:
			goto tr11
		case 101:
			goto st31
		case 104:
			goto st35
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 76:
			goto st32
		case 84:
			goto tr52
		case 95:
			goto tr11
		case 108:
			goto st32
		case 116:
			goto tr52
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto st33
		case 95:
			goto tr11
		case 101:
			goto st33
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 67:
			goto st34
		case 95:
			goto tr11
		case 99:
			goto st34
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 84:
			goto tr55
		case 95:
			goto tr11
		case 116:
			goto tr55
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 79:
			goto st36
		case 95:
			goto tr11
		case 111:
			goto st36
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 87:
			goto tr57
		case 95:
			goto tr11
		case 119:
			goto tr57
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto st38
		case 95:
			goto tr11
		case 101:
			goto st38
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 82:
			goto st39
		case 95:
			goto tr11
		case 114:
			goto st39
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 83:
			goto st40
		case 95:
			goto tr11
		case 115:
			goto st40
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 73:
			goto st41
		case 95:
			goto tr11
		case 105:
			goto st41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 79:
			goto st42
		case 95:
			goto tr11
		case 111:
			goto st42
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 78:
			goto tr63
		case 95:
			goto tr11
		case 110:
			goto tr63
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 72:
			goto st44
		case 95:
			goto tr11
		case 104:
			goto st44
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto st45
		case 95:
			goto tr11
		case 101:
			goto st45
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 82:
			goto st46
		case 95:
			goto tr11
		case 114:
			goto st46
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr11
		case 69:
			goto tr67
		case 95:
			goto tr11
		case 101:
			goto tr67
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr11
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr23
	st_out:
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof
	_test_eof20:  lex.cs = 20; goto _test_eof
	_test_eof21:  lex.cs = 21; goto _test_eof
	_test_eof22:  lex.cs = 22; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof25:  lex.cs = 25; goto _test_eof
	_test_eof26:  lex.cs = 26; goto _test_eof
	_test_eof27:  lex.cs = 27; goto _test_eof
	_test_eof28:  lex.cs = 28; goto _test_eof
	_test_eof29:  lex.cs = 29; goto _test_eof
	_test_eof30:  lex.cs = 30; goto _test_eof
	_test_eof31:  lex.cs = 31; goto _test_eof
	_test_eof32:  lex.cs = 32; goto _test_eof
	_test_eof33:  lex.cs = 33; goto _test_eof
	_test_eof34:  lex.cs = 34; goto _test_eof
	_test_eof35:  lex.cs = 35; goto _test_eof
	_test_eof36:  lex.cs = 36; goto _test_eof
	_test_eof37:  lex.cs = 37; goto _test_eof
	_test_eof38:  lex.cs = 38; goto _test_eof
	_test_eof39:  lex.cs = 39; goto _test_eof
	_test_eof40:  lex.cs = 40; goto _test_eof
	_test_eof41:  lex.cs = 41; goto _test_eof
	_test_eof42:  lex.cs = 42; goto _test_eof
	_test_eof43:  lex.cs = 43; goto _test_eof
	_test_eof44:  lex.cs = 44; goto _test_eof
	_test_eof45:  lex.cs = 45; goto _test_eof
	_test_eof46:  lex.cs = 46; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 3:
			goto tr21
		case 4:
			goto tr22
		case 5:
			goto tr23
		case 6:
			goto tr23
		case 7:
			goto tr23
		case 8:
			goto tr23
		case 9:
			goto tr23
		case 10:
			goto tr23
		case 11:
			goto tr23
		case 12:
			goto tr23
		case 13:
			goto tr23
		case 14:
			goto tr23
		case 15:
			goto tr23
		case 16:
			goto tr23
		case 17:
			goto tr23
		case 18:
			goto tr23
		case 19:
			goto tr23
		case 20:
			goto tr23
		case 21:
			goto tr23
		case 22:
			goto tr23
		case 23:
			goto tr23
		case 24:
			goto tr23
		case 25:
			goto tr23
		case 26:
			goto tr23
		case 27:
			goto tr23
		case 28:
			goto tr23
		case 29:
			goto tr23
		case 30:
			goto tr23
		case 31:
			goto tr23
		case 32:
			goto tr23
		case 33:
			goto tr23
		case 34:
			goto tr23
		case 35:
			goto tr23
		case 36:
			goto tr23
		case 37:
			goto tr23
		case 38:
			goto tr23
		case 39:
			goto tr23
		case 40:
			goto tr23
		case 41:
			goto tr23
		case 42:
			goto tr23
		case 43:
			goto tr23
		case 44:
			goto tr23
		case 45:
			goto tr23
		case 46:
			goto tr23
		}
	}

	_out: {}
	}

//line scan.rl:68


	return tok;
}

func (lex *lexer) Error(e string) {
	lex.err = e
}
