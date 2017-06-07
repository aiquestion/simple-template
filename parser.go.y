%{
package simple_template

/**
THIS FILE IS AUTO-GENERATE BY GOYACC
DO NOT EDIT
**/

import "strconv"
%}
%type<expr> expr
%type<expr> prefixexp
%type<expr> functioncall
%type<expr> var
%type<exprs> args
%type<exprs> exprlist

%union {
	token Token
	expr Expr
	exprs []Expr
}

%token<token> TIf TElse TEqeq TNeq TGteq TLteq TAnd TOr TTrue TFalse
%token<token> TIdent TString TNumber '('

%left TOr
%left TAnd
%left '>' '<' TGteq TLteq TEqeq TNeq
%left '+' '-'
%left '*' '/' '%'

%start exprlist

%%
exprlist:
    expr {
        $$ = []Expr{$1}
        if l, ok := yylex.(*yyLexerImpl); ok {
            l.Stmt = $$
        }
    }|
    exprlist ',' expr {
        $$ = append($1, $3)
        if l, ok := yylex.(*yyLexerImpl); ok {
            l.Stmt = $$
        }
    }

expr:
	TString {
		$$ = &StringExpr{Value: $1.Str}
		$$.SetLine($1.Line())
	}|
	TNumber {
	    v, _ := strconv.ParseFloat($1.Str, 0)
		$$ = &NumberExpr{Value: v}
		$$.SetLine($1.Line())
	}|
	TTrue {
		$$ = &TrueExpr{}
		$$.SetLine($1.Line())
	}|
	TFalse {
		$$ = &FalseExpr{}
		$$.SetLine($1.Line())
	}|
	expr TOr expr {
		$$ = &LogicalExpr{Lhs: $1, Operator: "||", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr TAnd expr {
		$$ = &LogicalExpr{Lhs: $1, Operator: "&&", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr '>' expr {
		$$ = &LogicalExpr{Lhs: $1, Operator: ">", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr '<' expr {
		$$ = &LogicalExpr{Lhs: $1, Operator: "<", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr TGteq expr {
		$$ = &LogicalExpr{Lhs: $1, Operator: ">=", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr TLteq expr {
		$$ = &LogicalExpr{Lhs: $1, Operator: "<=", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr TEqeq expr {
		$$ = &LogicalExpr{Lhs: $1, Operator: "==", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr TNeq expr {
		$$ = &LogicalExpr{Lhs: $1, Operator: "!=", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr '+' expr {
		$$ = &ArithmeticOpExpr{Lhs: $1, Operator: "+", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr '-' expr {
		$$ = &ArithmeticOpExpr{Lhs: $1, Operator: "-", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr '*' expr {
		$$ = &ArithmeticOpExpr{Lhs: $1, Operator: "*", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr '/' expr {
		$$ = &ArithmeticOpExpr{Lhs: $1, Operator: "/", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	expr '%' expr {
		$$ = &ArithmeticOpExpr{Lhs: $1, Operator: "%", Rhs: $3}
		$$.SetLine($1.Line())
	}|
	prefixexp {
		$$ = $1
	}|
	TIf '(' expr ')' '{' expr '}' {
	    $$ = &IfExpr{Cond: $3, Then: $6}
	}|
	TIf '(' expr ')' '{' expr '}' TElse '{' expr '}' {
	    $$ = &IfExpr{Cond: $3, Then: $6, Else: $10}
	}

prefixexp:
    var {
        $$ = $1
    }|
    functioncall {
        $$ = $1
    }|
    '(' expr ')' {
        $$ = $2
        $$.SetLine($1.Line())
    }

functioncall:
    TIdent args {
        $$ = &FuncCallExpr{Func: &IdentExpr{Value:$1.Str}, Args: $2}
        $$.SetLine($1.Line())
    }

args :
    '(' ')' {
        $$ = []Expr{}
    }|
    '(' exprlist ')' {
        $$ = $2
    }



var:
    TIdent {
        $$ = &IdentExpr{Value:$1.Str}
        $$.SetLine($1.Line())
    }|
    prefixexp '[' expr ']' {
        $$ = &AttrGetExpr{Object: $1, Key:$3}
        $$.SetLine($1.Line())
    }

%%


func TokenName(c int) string {
	if c >= TAnd && c-TAnd < len(yyToknames) {
		if yyToknames[c-TAnd] != "" {
			return yyToknames[c-TAnd]
		}
	}
    return string([]byte{byte(c)})
}

