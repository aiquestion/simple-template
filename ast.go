package simple_template

import (
	"fmt"
	"strings"
)

type Position struct {
	Source string
	Line   int
	Column int
}

type Token struct {
	Type rune
	Str  string
	Pos  Position
}

// TODO print better readible type
func (self *Token) String() string {
	return fmt.Sprintf("<%s,%s>", TokenName(int(self.Type)), self.Str)
}
func (self *Token) Line() int {
	return self.Pos.Line
}

type Expr interface {
	SetLine(line int)
	Line() int
	String() string
}

type BaseExpr struct {
	Pos Position
}

func (b *BaseExpr) String() string {
	return "<base, nil>"
}

func (e *BaseExpr) SetLine(line int) {
	e.Pos.Line = line
}
func (e *BaseExpr) Line() int {
	return e.Pos.Line
}

type StringExpr struct {
	BaseExpr

	Value string
}

func (b *StringExpr) String() string {
	return fmt.Sprintf("<strexpr,%s>", b.Value)
}

type NumberExpr struct {
	BaseExpr

	Value float64
}

func (e *NumberExpr) String() string {
	return fmt.Sprintf("<numexpr,%v>", e.Value)
}

type TrueExpr struct {
	BaseExpr
}

func (e *TrueExpr) String() string {
	return "<true expr>"
}

type FalseExpr struct {
	BaseExpr
}

func (e *FalseExpr) String() string {
	return "<false expr>"
}

type IdentExpr struct {
	BaseExpr

	Value string
}

func (e *IdentExpr) String() string {
	return fmt.Sprintf("<identxpr,%s>", e.Value)
}

type LogicalExpr struct {
	BaseExpr

	Lhs      Expr
	Operator string
	Rhs      Expr
}

func (e *LogicalExpr) String() string {
	return fmt.Sprintf("<%s,%s,%s>", e.Lhs.String(), e.Operator, e.Rhs.String())
}

type ArithmeticOpExpr struct {
	BaseExpr

	Lhs      Expr
	Operator string
	Rhs      Expr
}

func (e *ArithmeticOpExpr) String() string {
	return fmt.Sprintf("<%s,%s,%s>", e.Lhs.String(), e.Operator, e.Rhs.String())
}

type FuncCallExpr struct {
	BaseExpr

	Func *IdentExpr
	Args []Expr
}

func (e *FuncCallExpr) String() string {
	args := make([]string, 0, len(e.Args))
	for _, a := range e.Args {
		args = append(args, a.String())
	}
	return fmt.Sprintf("<%s(%s)>", e.Func.String(), strings.Join(args, ","))
}

type AttrGetExpr struct {
	BaseExpr

	Object Expr
	Key    Expr
}

func (e *AttrGetExpr) String() string {
	return fmt.Sprintf("<%s[%s]>", e.Object.String(), e.Key.String())
}

type IfExpr struct {
	BaseExpr

	Cond Expr
	Then Expr
	Else Expr
}

func (e *IfExpr) String() string {
	if e.Else == nil {
		return fmt.Sprintf("<if(%s)then{%s}>", e.Cond.String(), e.Then.String())
	} else {
		return fmt.Sprintf("<if(%s)then{%s}else{%s}>", e.Cond.String(), e.Then.String(), e.Else.String())
	}
}
