
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
}

func newLexer(data []byte) *lexer {
	lex := &lexer{ 
		data: data,
		pe: len(data),
	}
	
//line scan.go:35
	{
	 lex.cs = sql_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line scan.rl:35
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0
	
//line scan.go:49
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
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	}
	goto st_out
tr1:
//line scan.rl:93
 lex.te = ( lex.p)+1
{ out.str = string(lex.data[lex.ts+1:lex.te-1]); tok = SLITERAL; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr2:
//line scan.rl:95
 lex.te = ( lex.p)+1

	goto st2
tr4:
//line scan.rl:45
 lex.te = ( lex.p)+1
{ tok = '('; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr5:
//line scan.rl:46
 lex.te = ( lex.p)+1
{ tok = ')'; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr6:
//line scan.rl:47
 lex.te = ( lex.p)+1
{ tok = '*'; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr7:
//line scan.rl:44
 lex.te = ( lex.p)+1
{ tok = ','; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr9:
//line scan.rl:43
 lex.te = ( lex.p)+1
{ tok = ';'; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr11:
//line scan.rl:48
 lex.te = ( lex.p)+1
{ tok = '='; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr30:
//line scan.rl:94
 lex.te = ( lex.p)
( lex.p)--
{ out.str = string(lex.data[lex.ts:lex.te]); tok = NUMBER; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr31:
//line scan.rl:49
 lex.te = ( lex.p)
( lex.p)--
{ tok = '<'; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr32:
//line scan.rl:51
 lex.te = ( lex.p)+1
{ tok = LT_OR_EQUAL; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr33:
//line scan.rl:53
 lex.te = ( lex.p)+1
{ tok = NOT_EQUAL; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr34:
//line scan.rl:50
 lex.te = ( lex.p)
( lex.p)--
{ tok = '>'; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr35:
//line scan.rl:52
 lex.te = ( lex.p)+1
{ tok = GT_OR_EQUAL; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr36:
//line scan.rl:92
 lex.te = ( lex.p)
( lex.p)--
{ out.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr41:
//line NONE:1
	switch  lex.act {
	case 12:
	{( lex.p) = ( lex.te) - 1
 tok = ADD; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 13:
	{( lex.p) = ( lex.te) - 1
 tok = ALL; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 14:
	{( lex.p) = ( lex.te) - 1
 tok = ALTER; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 15:
	{( lex.p) = ( lex.te) - 1
 tok = AND; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 16:
	{( lex.p) = ( lex.te) - 1
 tok = ASC; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 17:
	{( lex.p) = ( lex.te) - 1
 tok = BY; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 18:
	{( lex.p) = ( lex.te) - 1
 tok = CREATE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 19:
	{( lex.p) = ( lex.te) - 1
 tok = COUNT; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 20:
	{( lex.p) = ( lex.te) - 1
 tok = DELETE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 21:
	{( lex.p) = ( lex.te) - 1
 tok = DESC; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 22:
	{( lex.p) = ( lex.te) - 1
 tok = DROP; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 23:
	{( lex.p) = ( lex.te) - 1
 tok = ENCRYPTED; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 24:
	{( lex.p) = ( lex.te) - 1
 tok = FILTER; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 25:
	{( lex.p) = ( lex.te) - 1
 tok = FROM; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 26:
	{( lex.p) = ( lex.te) - 1
 tok = FUND; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 27:
	{( lex.p) = ( lex.te) - 1
 tok = ILIKE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 28:
	{( lex.p) = ( lex.te) - 1
 tok = INFO; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 29:
	{( lex.p) = ( lex.te) - 1
 tok = INSERT; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 30:
	{( lex.p) = ( lex.te) - 1
 tok = INTO; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 31:
	{( lex.p) = ( lex.te) - 1
 tok = LIKE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 32:
	{( lex.p) = ( lex.te) - 1
 tok = LIMIT; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 33:
	{( lex.p) = ( lex.te) - 1
 tok = NOT; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 34:
	{( lex.p) = ( lex.te) - 1
 tok = OFFSET; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 36:
	{( lex.p) = ( lex.te) - 1
 tok = ORDER; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 37:
	{( lex.p) = ( lex.te) - 1
 tok = PASSWORD; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 38:
	{( lex.p) = ( lex.te) - 1
 tok = PROJECT; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 39:
	{( lex.p) = ( lex.te) - 1
 tok = PROPERTY; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 40:
	{( lex.p) = ( lex.te) - 1
 tok = RETRIEVE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 41:
	{( lex.p) = ( lex.te) - 1
 tok = SET; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 42:
	{( lex.p) = ( lex.te) - 1
 tok = SHOW; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 43:
	{( lex.p) = ( lex.te) - 1
 tok = TAG; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 44:
	{( lex.p) = ( lex.te) - 1
 tok = TO; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 45:
	{( lex.p) = ( lex.te) - 1
 tok = PING; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 46:
	{( lex.p) = ( lex.te) - 1
 tok = SELECT; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 47:
	{( lex.p) = ( lex.te) - 1
 tok = USER; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 48:
	{( lex.p) = ( lex.te) - 1
 tok = WHERE; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 49:
	{( lex.p) = ( lex.te) - 1
 tok = WITH; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 50:
	{( lex.p) = ( lex.te) - 1
 out.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; {( lex.p)++;  lex.cs = 2; goto _out } }
	}
	
	goto st2
tr115:
//line scan.rl:77
 lex.te = ( lex.p)
( lex.p)--
{ tok = OR; {( lex.p)++;  lex.cs = 2; goto _out } }
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

//line scan.go:505
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
		case 60:
			goto st4
		case 61:
			goto tr11
		case 62:
			goto st5
		case 65:
			goto st6
		case 66:
			goto st14
		case 67:
			goto st15
		case 68:
			goto st23
		case 69:
			goto st31
		case 70:
			goto st39
		case 73:
			goto st48
		case 76:
			goto st58
		case 78:
			goto st63
		case 79:
			goto st65
		case 80:
			goto st73
		case 82:
			goto st91
		case 83:
			goto st98
		case 84:
			goto st105
		case 85:
			goto st107
		case 87:
			goto st110
		case 95:
			goto tr19
		case 97:
			goto st6
		case 98:
			goto st14
		case 99:
			goto st15
		case 100:
			goto st23
		case 101:
			goto st31
		case 102:
			goto st39
		case 105:
			goto st48
		case 108:
			goto st58
		case 110:
			goto st63
		case 111:
			goto st65
		case 112:
			goto st73
		case 114:
			goto st91
		case 115:
			goto st98
		case 116:
			goto st105
		case 117:
			goto st107
		case 119:
			goto st110
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 103 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr19
				}
			case  lex.data[( lex.p)] >= 71:
				goto tr19
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
		goto tr30
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 61:
			goto tr32
		case 62:
			goto tr33
		}
		goto tr31
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if  lex.data[( lex.p)] == 61 {
			goto tr35
		}
		goto tr34
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 68:
			goto st8
		case 76:
			goto st9
		case 78:
			goto st12
		case 83:
			goto st13
		case 95:
			goto tr19
		case 100:
			goto st8
		case 108:
			goto st9
		case 110:
			goto st12
		case 115:
			goto st13
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
tr19:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:92
 lex.act = 50;
	goto st7
tr42:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:54
 lex.act = 12;
	goto st7
tr43:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:55
 lex.act = 13;
	goto st7
tr46:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:56
 lex.act = 14;
	goto st7
tr47:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:57
 lex.act = 15;
	goto st7
tr48:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:58
 lex.act = 16;
	goto st7
tr49:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:59
 lex.act = 17;
	goto st7
tr54:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:61
 lex.act = 19;
	goto st7
tr58:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:60
 lex.act = 18;
	goto st7
tr65:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:62
 lex.act = 20;
	goto st7
tr66:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:63
 lex.act = 21;
	goto st7
tr68:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:64
 lex.act = 22;
	goto st7
tr76:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:65
 lex.act = 23;
	goto st7
tr83:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:66
 lex.act = 24;
	goto st7
tr85:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:67
 lex.act = 25;
	goto st7
tr87:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:68
 lex.act = 26;
	goto st7
tr92:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:69
 lex.act = 27;
	goto st7
tr96:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:70
 lex.act = 28;
	goto st7
tr99:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:71
 lex.act = 29;
	goto st7
tr100:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:72
 lex.act = 30;
	goto st7
tr104:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:73
 lex.act = 31;
	goto st7
tr106:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:74
 lex.act = 32;
	goto st7
tr108:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:75
 lex.act = 33;
	goto st7
tr114:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:76
 lex.act = 34;
	goto st7
tr118:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:78
 lex.act = 36;
	goto st7
tr127:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:79
 lex.act = 37;
	goto st7
tr129:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:87
 lex.act = 45;
	goto st7
tr135:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:80
 lex.act = 38;
	goto st7
tr139:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:81
 lex.act = 39;
	goto st7
tr146:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:82
 lex.act = 40;
	goto st7
tr150:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:83
 lex.act = 41;
	goto st7
tr153:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:88
 lex.act = 46;
	goto st7
tr155:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:84
 lex.act = 42;
	goto st7
tr157:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:86
 lex.act = 44;
	goto st7
tr158:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:85
 lex.act = 43;
	goto st7
tr161:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:89
 lex.act = 47;
	goto st7
tr166:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:90
 lex.act = 48;
	goto st7
tr168:
//line NONE:1
 lex.te = ( lex.p)+1

//line scan.rl:91
 lex.act = 49;
	goto st7
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
//line scan.go:966
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 95:
			goto tr19
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr41
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 68:
			goto tr42
		case 95:
			goto tr19
		case 100:
			goto tr42
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 76:
			goto tr43
		case 84:
			goto st10
		case 95:
			goto tr19
		case 108:
			goto tr43
		case 116:
			goto st10
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st11
		case 95:
			goto tr19
		case 101:
			goto st11
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto tr46
		case 95:
			goto tr19
		case 114:
			goto tr46
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 68:
			goto tr47
		case 95:
			goto tr19
		case 100:
			goto tr47
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 67:
			goto tr48
		case 95:
			goto tr19
		case 99:
			goto tr48
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 89:
			goto tr49
		case 95:
			goto tr19
		case 121:
			goto tr49
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto st16
		case 82:
			goto st19
		case 95:
			goto tr19
		case 111:
			goto st16
		case 114:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 85:
			goto st17
		case 95:
			goto tr19
		case 117:
			goto st17
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 78:
			goto st18
		case 95:
			goto tr19
		case 110:
			goto st18
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto tr54
		case 95:
			goto tr19
		case 116:
			goto tr54
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st20
		case 95:
			goto tr19
		case 101:
			goto st20
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 65:
			goto st21
		case 95:
			goto tr19
		case 97:
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto st22
		case 95:
			goto tr19
		case 116:
			goto st22
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto tr58
		case 95:
			goto tr19
		case 101:
			goto tr58
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st24
		case 82:
			goto st29
		case 95:
			goto tr19
		case 101:
			goto st24
		case 114:
			goto st29
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 76:
			goto st25
		case 83:
			goto st28
		case 95:
			goto tr19
		case 108:
			goto st25
		case 115:
			goto st28
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st26
		case 95:
			goto tr19
		case 101:
			goto st26
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto st27
		case 95:
			goto tr19
		case 116:
			goto st27
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto tr65
		case 95:
			goto tr19
		case 101:
			goto tr65
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 67:
			goto tr66
		case 95:
			goto tr19
		case 99:
			goto tr66
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto st30
		case 95:
			goto tr19
		case 111:
			goto st30
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 80:
			goto tr68
		case 95:
			goto tr19
		case 112:
			goto tr68
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 78:
			goto st32
		case 95:
			goto tr19
		case 110:
			goto st32
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 67:
			goto st33
		case 95:
			goto tr19
		case 99:
			goto st33
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto st34
		case 95:
			goto tr19
		case 114:
			goto st34
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 89:
			goto st35
		case 95:
			goto tr19
		case 121:
			goto st35
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 80:
			goto st36
		case 95:
			goto tr19
		case 112:
			goto st36
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto st37
		case 95:
			goto tr19
		case 116:
			goto st37
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st38
		case 95:
			goto tr19
		case 101:
			goto st38
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 68:
			goto tr76
		case 95:
			goto tr19
		case 100:
			goto tr76
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 73:
			goto st40
		case 82:
			goto st44
		case 85:
			goto st46
		case 95:
			goto tr19
		case 105:
			goto st40
		case 114:
			goto st44
		case 117:
			goto st46
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 76:
			goto st41
		case 95:
			goto tr19
		case 108:
			goto st41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto st42
		case 95:
			goto tr19
		case 116:
			goto st42
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st43
		case 95:
			goto tr19
		case 101:
			goto st43
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto tr83
		case 95:
			goto tr19
		case 114:
			goto tr83
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto st45
		case 95:
			goto tr19
		case 111:
			goto st45
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 77:
			goto tr85
		case 95:
			goto tr19
		case 109:
			goto tr85
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 78:
			goto st47
		case 95:
			goto tr19
		case 110:
			goto st47
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st47:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 68:
			goto tr87
		case 95:
			goto tr19
		case 100:
			goto tr87
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st48:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 76:
			goto st49
		case 78:
			goto st52
		case 95:
			goto tr19
		case 108:
			goto st49
		case 110:
			goto st52
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st49:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 73:
			goto st50
		case 95:
			goto tr19
		case 105:
			goto st50
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st50:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 75:
			goto st51
		case 95:
			goto tr19
		case 107:
			goto st51
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st51:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto tr92
		case 95:
			goto tr19
		case 101:
			goto tr92
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st52:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 70:
			goto st53
		case 83:
			goto st54
		case 84:
			goto st57
		case 95:
			goto tr19
		case 102:
			goto st53
		case 115:
			goto st54
		case 116:
			goto st57
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st53:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto tr96
		case 95:
			goto tr19
		case 111:
			goto tr96
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st54:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st55
		case 95:
			goto tr19
		case 101:
			goto st55
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st55:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto st56
		case 95:
			goto tr19
		case 114:
			goto st56
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st56:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto tr99
		case 95:
			goto tr19
		case 116:
			goto tr99
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st57:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto tr100
		case 95:
			goto tr19
		case 111:
			goto tr100
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st58:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 73:
			goto st59
		case 95:
			goto tr19
		case 105:
			goto st59
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st59:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 75:
			goto st60
		case 77:
			goto st61
		case 95:
			goto tr19
		case 107:
			goto st60
		case 109:
			goto st61
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st60:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto tr104
		case 95:
			goto tr19
		case 101:
			goto tr104
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st61:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 73:
			goto st62
		case 95:
			goto tr19
		case 105:
			goto st62
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st62:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto tr106
		case 95:
			goto tr19
		case 116:
			goto tr106
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st63:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto st64
		case 95:
			goto tr19
		case 111:
			goto st64
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st64:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto tr108
		case 95:
			goto tr19
		case 116:
			goto tr108
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st65:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 70:
			goto st66
		case 82:
			goto st70
		case 95:
			goto tr19
		case 102:
			goto st66
		case 114:
			goto st70
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st66:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 70:
			goto st67
		case 95:
			goto tr19
		case 102:
			goto st67
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st67:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 83:
			goto st68
		case 95:
			goto tr19
		case 115:
			goto st68
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st68:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st69
		case 95:
			goto tr19
		case 101:
			goto st69
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st69:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto tr114
		case 95:
			goto tr19
		case 116:
			goto tr114
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st70:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 68:
			goto st71
		case 95:
			goto tr19
		case 100:
			goto st71
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr115
	st71:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st72
		case 95:
			goto tr19
		case 101:
			goto st72
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st72:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto tr118
		case 95:
			goto tr19
		case 114:
			goto tr118
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st73:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 65:
			goto st74
		case 73:
			goto st80
		case 82:
			goto st82
		case 95:
			goto tr19
		case 97:
			goto st74
		case 105:
			goto st80
		case 114:
			goto st82
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st74:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 83:
			goto st75
		case 95:
			goto tr19
		case 115:
			goto st75
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st75:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 83:
			goto st76
		case 95:
			goto tr19
		case 115:
			goto st76
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st76:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 87:
			goto st77
		case 95:
			goto tr19
		case 119:
			goto st77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st77:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto st78
		case 95:
			goto tr19
		case 111:
			goto st78
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st78:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto st79
		case 95:
			goto tr19
		case 114:
			goto st79
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st79:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 68:
			goto tr127
		case 95:
			goto tr19
		case 100:
			goto tr127
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st80:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 78:
			goto st81
		case 95:
			goto tr19
		case 110:
			goto st81
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st81:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 71:
			goto tr129
		case 95:
			goto tr19
		case 103:
			goto tr129
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st82:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto st83
		case 95:
			goto tr19
		case 111:
			goto st83
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st83:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 74:
			goto st84
		case 80:
			goto st87
		case 95:
			goto tr19
		case 106:
			goto st84
		case 112:
			goto st87
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st84:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st85
		case 95:
			goto tr19
		case 101:
			goto st85
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st85:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 67:
			goto st86
		case 95:
			goto tr19
		case 99:
			goto st86
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st86:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof86
		}
	st_case_86:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto tr135
		case 95:
			goto tr19
		case 116:
			goto tr135
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st87:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st88
		case 95:
			goto tr19
		case 101:
			goto st88
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st88:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto st89
		case 95:
			goto tr19
		case 114:
			goto st89
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st89:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto st90
		case 95:
			goto tr19
		case 116:
			goto st90
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st90:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 89:
			goto tr139
		case 95:
			goto tr19
		case 121:
			goto tr139
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st91:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st92
		case 95:
			goto tr19
		case 101:
			goto st92
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st92:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto st93
		case 95:
			goto tr19
		case 116:
			goto st93
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st93:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto st94
		case 95:
			goto tr19
		case 114:
			goto st94
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st94:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 73:
			goto st95
		case 95:
			goto tr19
		case 105:
			goto st95
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st95:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st96
		case 95:
			goto tr19
		case 101:
			goto st96
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st96:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof96
		}
	st_case_96:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 86:
			goto st97
		case 95:
			goto tr19
		case 118:
			goto st97
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st97:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto tr146
		case 95:
			goto tr19
		case 101:
			goto tr146
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st98:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st99
		case 72:
			goto st103
		case 95:
			goto tr19
		case 101:
			goto st99
		case 104:
			goto st103
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st99:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 76:
			goto st100
		case 84:
			goto tr150
		case 95:
			goto tr19
		case 108:
			goto st100
		case 116:
			goto tr150
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st100:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st101
		case 95:
			goto tr19
		case 101:
			goto st101
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st101:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 67:
			goto st102
		case 95:
			goto tr19
		case 99:
			goto st102
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st102:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto tr153
		case 95:
			goto tr19
		case 116:
			goto tr153
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st103:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 79:
			goto st104
		case 95:
			goto tr19
		case 111:
			goto st104
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st104:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 87:
			goto tr155
		case 95:
			goto tr19
		case 119:
			goto tr155
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st105:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 65:
			goto st106
		case 79:
			goto tr157
		case 95:
			goto tr19
		case 97:
			goto st106
		case 111:
			goto tr157
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st106:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 71:
			goto tr158
		case 95:
			goto tr19
		case 103:
			goto tr158
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st107:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 83:
			goto st108
		case 95:
			goto tr19
		case 115:
			goto st108
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st108:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof108
		}
	st_case_108:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st109
		case 95:
			goto tr19
		case 101:
			goto st109
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st109:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof109
		}
	st_case_109:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto tr161
		case 95:
			goto tr19
		case 114:
			goto tr161
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st110:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof110
		}
	st_case_110:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 72:
			goto st111
		case 73:
			goto st114
		case 95:
			goto tr19
		case 104:
			goto st111
		case 105:
			goto st114
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st111:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof111
		}
	st_case_111:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto st112
		case 95:
			goto tr19
		case 101:
			goto st112
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st112:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof112
		}
	st_case_112:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 82:
			goto st113
		case 95:
			goto tr19
		case 114:
			goto st113
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st113:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 69:
			goto tr166
		case 95:
			goto tr19
		case 101:
			goto tr166
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st114:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof114
		}
	st_case_114:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 84:
			goto st115
		case 95:
			goto tr19
		case 116:
			goto st115
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
	st115:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof115
		}
	st_case_115:
		switch  lex.data[( lex.p)] {
		case 46:
			goto tr19
		case 72:
			goto tr168
		case 95:
			goto tr19
		case 104:
			goto tr168
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto tr36
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
	_test_eof47:  lex.cs = 47; goto _test_eof
	_test_eof48:  lex.cs = 48; goto _test_eof
	_test_eof49:  lex.cs = 49; goto _test_eof
	_test_eof50:  lex.cs = 50; goto _test_eof
	_test_eof51:  lex.cs = 51; goto _test_eof
	_test_eof52:  lex.cs = 52; goto _test_eof
	_test_eof53:  lex.cs = 53; goto _test_eof
	_test_eof54:  lex.cs = 54; goto _test_eof
	_test_eof55:  lex.cs = 55; goto _test_eof
	_test_eof56:  lex.cs = 56; goto _test_eof
	_test_eof57:  lex.cs = 57; goto _test_eof
	_test_eof58:  lex.cs = 58; goto _test_eof
	_test_eof59:  lex.cs = 59; goto _test_eof
	_test_eof60:  lex.cs = 60; goto _test_eof
	_test_eof61:  lex.cs = 61; goto _test_eof
	_test_eof62:  lex.cs = 62; goto _test_eof
	_test_eof63:  lex.cs = 63; goto _test_eof
	_test_eof64:  lex.cs = 64; goto _test_eof
	_test_eof65:  lex.cs = 65; goto _test_eof
	_test_eof66:  lex.cs = 66; goto _test_eof
	_test_eof67:  lex.cs = 67; goto _test_eof
	_test_eof68:  lex.cs = 68; goto _test_eof
	_test_eof69:  lex.cs = 69; goto _test_eof
	_test_eof70:  lex.cs = 70; goto _test_eof
	_test_eof71:  lex.cs = 71; goto _test_eof
	_test_eof72:  lex.cs = 72; goto _test_eof
	_test_eof73:  lex.cs = 73; goto _test_eof
	_test_eof74:  lex.cs = 74; goto _test_eof
	_test_eof75:  lex.cs = 75; goto _test_eof
	_test_eof76:  lex.cs = 76; goto _test_eof
	_test_eof77:  lex.cs = 77; goto _test_eof
	_test_eof78:  lex.cs = 78; goto _test_eof
	_test_eof79:  lex.cs = 79; goto _test_eof
	_test_eof80:  lex.cs = 80; goto _test_eof
	_test_eof81:  lex.cs = 81; goto _test_eof
	_test_eof82:  lex.cs = 82; goto _test_eof
	_test_eof83:  lex.cs = 83; goto _test_eof
	_test_eof84:  lex.cs = 84; goto _test_eof
	_test_eof85:  lex.cs = 85; goto _test_eof
	_test_eof86:  lex.cs = 86; goto _test_eof
	_test_eof87:  lex.cs = 87; goto _test_eof
	_test_eof88:  lex.cs = 88; goto _test_eof
	_test_eof89:  lex.cs = 89; goto _test_eof
	_test_eof90:  lex.cs = 90; goto _test_eof
	_test_eof91:  lex.cs = 91; goto _test_eof
	_test_eof92:  lex.cs = 92; goto _test_eof
	_test_eof93:  lex.cs = 93; goto _test_eof
	_test_eof94:  lex.cs = 94; goto _test_eof
	_test_eof95:  lex.cs = 95; goto _test_eof
	_test_eof96:  lex.cs = 96; goto _test_eof
	_test_eof97:  lex.cs = 97; goto _test_eof
	_test_eof98:  lex.cs = 98; goto _test_eof
	_test_eof99:  lex.cs = 99; goto _test_eof
	_test_eof100:  lex.cs = 100; goto _test_eof
	_test_eof101:  lex.cs = 101; goto _test_eof
	_test_eof102:  lex.cs = 102; goto _test_eof
	_test_eof103:  lex.cs = 103; goto _test_eof
	_test_eof104:  lex.cs = 104; goto _test_eof
	_test_eof105:  lex.cs = 105; goto _test_eof
	_test_eof106:  lex.cs = 106; goto _test_eof
	_test_eof107:  lex.cs = 107; goto _test_eof
	_test_eof108:  lex.cs = 108; goto _test_eof
	_test_eof109:  lex.cs = 109; goto _test_eof
	_test_eof110:  lex.cs = 110; goto _test_eof
	_test_eof111:  lex.cs = 111; goto _test_eof
	_test_eof112:  lex.cs = 112; goto _test_eof
	_test_eof113:  lex.cs = 113; goto _test_eof
	_test_eof114:  lex.cs = 114; goto _test_eof
	_test_eof115:  lex.cs = 115; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 3:
			goto tr30
		case 4:
			goto tr31
		case 5:
			goto tr34
		case 6:
			goto tr36
		case 7:
			goto tr41
		case 8:
			goto tr36
		case 9:
			goto tr36
		case 10:
			goto tr36
		case 11:
			goto tr36
		case 12:
			goto tr36
		case 13:
			goto tr36
		case 14:
			goto tr36
		case 15:
			goto tr36
		case 16:
			goto tr36
		case 17:
			goto tr36
		case 18:
			goto tr36
		case 19:
			goto tr36
		case 20:
			goto tr36
		case 21:
			goto tr36
		case 22:
			goto tr36
		case 23:
			goto tr36
		case 24:
			goto tr36
		case 25:
			goto tr36
		case 26:
			goto tr36
		case 27:
			goto tr36
		case 28:
			goto tr36
		case 29:
			goto tr36
		case 30:
			goto tr36
		case 31:
			goto tr36
		case 32:
			goto tr36
		case 33:
			goto tr36
		case 34:
			goto tr36
		case 35:
			goto tr36
		case 36:
			goto tr36
		case 37:
			goto tr36
		case 38:
			goto tr36
		case 39:
			goto tr36
		case 40:
			goto tr36
		case 41:
			goto tr36
		case 42:
			goto tr36
		case 43:
			goto tr36
		case 44:
			goto tr36
		case 45:
			goto tr36
		case 46:
			goto tr36
		case 47:
			goto tr36
		case 48:
			goto tr36
		case 49:
			goto tr36
		case 50:
			goto tr36
		case 51:
			goto tr36
		case 52:
			goto tr36
		case 53:
			goto tr36
		case 54:
			goto tr36
		case 55:
			goto tr36
		case 56:
			goto tr36
		case 57:
			goto tr36
		case 58:
			goto tr36
		case 59:
			goto tr36
		case 60:
			goto tr36
		case 61:
			goto tr36
		case 62:
			goto tr36
		case 63:
			goto tr36
		case 64:
			goto tr36
		case 65:
			goto tr36
		case 66:
			goto tr36
		case 67:
			goto tr36
		case 68:
			goto tr36
		case 69:
			goto tr36
		case 70:
			goto tr115
		case 71:
			goto tr36
		case 72:
			goto tr36
		case 73:
			goto tr36
		case 74:
			goto tr36
		case 75:
			goto tr36
		case 76:
			goto tr36
		case 77:
			goto tr36
		case 78:
			goto tr36
		case 79:
			goto tr36
		case 80:
			goto tr36
		case 81:
			goto tr36
		case 82:
			goto tr36
		case 83:
			goto tr36
		case 84:
			goto tr36
		case 85:
			goto tr36
		case 86:
			goto tr36
		case 87:
			goto tr36
		case 88:
			goto tr36
		case 89:
			goto tr36
		case 90:
			goto tr36
		case 91:
			goto tr36
		case 92:
			goto tr36
		case 93:
			goto tr36
		case 94:
			goto tr36
		case 95:
			goto tr36
		case 96:
			goto tr36
		case 97:
			goto tr36
		case 98:
			goto tr36
		case 99:
			goto tr36
		case 100:
			goto tr36
		case 101:
			goto tr36
		case 102:
			goto tr36
		case 103:
			goto tr36
		case 104:
			goto tr36
		case 105:
			goto tr36
		case 106:
			goto tr36
		case 107:
			goto tr36
		case 108:
			goto tr36
		case 109:
			goto tr36
		case 110:
			goto tr36
		case 111:
			goto tr36
		case 112:
			goto tr36
		case 113:
			goto tr36
		case 114:
			goto tr36
		case 115:
			goto tr36
		}
	}

	_out: {}
	}

//line scan.rl:99


	return tok;
}

func (lex *lexer) Error(e string) {
	lex.err = e
}
