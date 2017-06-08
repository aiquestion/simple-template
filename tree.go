package simple_template

import (
	"fmt"
	"math"
	"reflect"
)

var buildIns = map[string]Func{
	"toNumber": toNumber,
	"toBool":   toBool,
	"isNaN":    isNaN,
}

func NewExecuteContenxt(idents map[string]interface{}, funcs map[string]Func) *ExecuteContext {
	finalFunc := make(map[string]Func)
	for k, v := range buildIns {
		finalFunc[k] = v
	}
	for k, v := range funcs {
		finalFunc[k] = v
	}
	return &ExecuteContext{
		definedIdent: idents,
		definedFunc:  finalFunc,
	}
}

type Func func(inputs []interface{}) (interface{}, error)

type ExecuteContext struct {
	definedIdent map[string]interface{}
	definedFunc  map[string]Func
}

func (ec *ExecuteContext) EvaluateExpression(expr Expr) (interface{}, error) {
	switch vt := expr.(type) {
	case *StringExpr:
		return vt.Value, nil
	case *NumberExpr:
		return vt.Value, nil
	case *TrueExpr:
		return true, nil
	case *FalseExpr:
		return false, nil
	case *IdentExpr:
		return ec.evalIdentityExpr(vt)
	case *LogicalExpr:
		return ec.evalLogicalExpr(vt)
	case *ArithmeticOpExpr:
		return ec.evalArithmethicExpr(vt)
	case *IfExpr:
		return ec.evalIfExpr(vt)
	case *AttrGetExpr:
		return ec.evalAttrGetExpr(vt)
	case *FuncCallExpr:
		return ec.evalFuncCallExpr(vt)
	}
	return nil, fmt.Errorf("unknow expression %T", expr)
}

func (ec *ExecuteContext) evalFuncCallExpr(e *FuncCallExpr) (interface{}, error) {
	if ec.definedFunc == nil {
		return nil, fmt.Errorf("function %v cannot be find", e.Func.Value)
	}

	f := ec.definedFunc[e.Func.Value]
	if f != nil {
		inputs := make([]interface{}, 0, len(e.Args))
		for _, a := range e.Args {
			va, err := ec.EvaluateExpression(a)
			if err != nil {
				return nil, err
			}
			inputs = append(inputs, va)
		}
		return f(inputs)
	}

	return nil, fmt.Errorf("function %v cannot be find", e.Func.Value)
}

func (ec *ExecuteContext) evalIdentityExpr(e *IdentExpr) (interface{}, error) {
	if ec.definedIdent == nil {
		return nil, fmt.Errorf("%v cannot be find", e.Value)
	}
	ival := ec.definedIdent[e.Value]
	if ival != nil {
		return ival, nil
	}
	return nil, fmt.Errorf("%v cannot be find", e.Value)
}

func (ec *ExecuteContext) evalLogicalExpr(e *LogicalExpr) (bool, error) {
	left, err := ec.EvaluateExpression(e.Lhs)
	if err != nil {
		return false, err
	}

	// need special process for these 2 operator
	if e.Operator == "&&" || e.Operator == "||" {
		leftBool, ok := TryBool(left)
		if !ok {
			return false, fmt.Errorf("left expr value unknown, %v", left)
		}
		if e.Operator == "&&" && !leftBool {
			return false, nil
		}
		if e.Operator == "||" && leftBool {
			return true, nil
		}
		right, err := ec.EvaluateExpression(e.Rhs)
		if err != nil {
			return false, err
		}
		rightBool, ok := TryBool(right)
		if !ok {
			return false, fmt.Errorf("right expr value unknown, %v", left)
		}
		if e.Operator == "||" {
			return leftBool || rightBool, nil
		} else {
			return leftBool && rightBool, nil
		}
	}

	right, err := ec.EvaluateExpression(e.Rhs)
	if err != nil {
		return false, err
	}
	switch e.Operator {
	case "==", "!=":
		eq := true
		if e.Operator == "!=" {
			eq = !eq
		}

		// one is number, then compare the other to number
		_, lok := left.(float64)
		_, rok := right.(float64)
		if lok || rok {
			lnum, ok := TryNumber(left)
			if !ok {
				return false, nil
			}
			rnum, ok := TryNumber(right)
			if !ok {
				return false, nil
			}
			if eq {
				return lnum == rnum, nil
			} else {
				return lnum != rnum, nil
			}
		}

		// one is bool
		_, lok = left.(bool)
		_, rok = right.(bool)
		if lok || rok {
			lb, ok := TryBool(left)
			if !ok {
				return false, nil
			}
			rb, ok := TryBool(right)
			if !ok {
				return false, nil
			}
			if eq {
				return lb == rb, nil
			} else {
				return lb != rb, nil
			}
		}

		// then all should be string, or we get an err
		_, lok = left.(string)
		_, rok = right.(string)
		if lok && rok {
			lb, ok := TryBool(left)
			if !ok {
				return false, nil
			}
			rb, ok := TryBool(right)
			if !ok {
				return false, nil
			}
			if eq {
				return lb == rb, nil
			} else {
				return lb != rb, nil
			}
		} else {
			return false, fmt.Errorf("logical expression with left or right side type unknow %v => %v", left, right)
		}
	case ">", "<", ">=", "<=":
		// will convert both side to number
		lnum, ok := TryNumber(left)
		if !ok {
			return false, fmt.Errorf("%v cannot conver to number when compare", left)
		}
		rnum, ok := TryNumber(right)
		if !ok {
			return false, fmt.Errorf("%v cannot conver to number when compare", right)
		}
		switch e.Operator {
		case ">":
			return lnum > rnum, nil
		case "<":
			return lnum < rnum, nil
		case ">=":
			return lnum >= rnum, nil
		case "<=":
			return lnum <= rnum, nil
		default:
			panic("you should not reach here")
		}
	}
	return false, fmt.Errorf("invalid operator %v", e.Operator)
}

func (ec *ExecuteContext) evalArithmethicExpr(e *ArithmeticOpExpr) (interface{}, error) {
	left, err := ec.EvaluateExpression(e.Lhs)
	if err != nil {
		return false, err
	}
	right, err := ec.EvaluateExpression(e.Rhs)
	if err != nil {
		return false, err
	}
	switch e.Operator {
	case "+":
		_, lok := left.(string)
		_, rok := right.(string)
		if lok || rok {
			lstr, ok := TryString(left)
			if !ok {
				return false, fmt.Errorf("%v cannot conver to string when +", left)
			}
			rstr, ok := TryString(right)
			if !ok {
				return false, fmt.Errorf("%v cannot conver to string when +", right)
			}
			return lstr + rstr, nil
		}
		fallthrough
	case "-", "*", "/", "%":
		lnum, ok := TryNumber(left)
		if !ok {
			return false, fmt.Errorf("%v cannot conver to number when +", left)
		}
		rnum, ok := TryNumber(right)
		if !ok {
			return false, fmt.Errorf("%v cannot conver to number when +", right)
		}
		switch e.Operator {
		case "+":
			return lnum + rnum, nil
		case "-":
			return lnum - rnum, nil
		case "*":
			return lnum * rnum, nil
		case "/":
			if rnum == 0 {
				return nil, fmt.Errorf("%v / %v : divided by zero", left, right)
			}
			return lnum / rnum, nil
		case "%":
			lint := math.Trunc(lnum)
			rint := math.Trunc(rnum)
			if lint != lnum || rint != rnum {
				return false, fmt.Errorf("%v or %v cannot be conver to int in % operation", lnum, rnum)
			}
			return int64(lint) % int64(rint), nil
		default:
			panic("you should never reach here")
		}
		return lnum + rnum, nil
	}
	return nil, fmt.Errorf("invalid operator %v", e.Operator)
}

func (ec *ExecuteContext) evalIfExpr(e *IfExpr) (interface{}, error) {
	cond, err := ec.EvaluateExpression(e.Cond)
	if err != nil {
		return nil, err
	}
	if condB, ok := TryBool(cond); ok {
		if condB {
			return ec.EvaluateExpression(e.Then)
		} else {
			if e.Else != nil {
				return ec.EvaluateExpression(e.Else)
			}
			return nil, nil
		}
	} else {
		return nil, fmt.Errorf("condition cannot eval to bool %v", cond)
	}
}

func (ec *ExecuteContext) evalAttrGetExpr(e *AttrGetExpr) (interface{}, error) {
	obj, err := ec.EvaluateExpression(e.Object)
	if err != nil {
		return nil, err
	}
	key, err := ec.EvaluateExpression(e.Key)
	if err != nil {
		return nil, err
	}
	// array
	objr := reflect.ValueOf(obj)
	if objr.Type().Kind() == reflect.Slice || objr.Type().Kind() == reflect.Array {
		index, ok := TryNumber(key)
		if !ok {
			return nil, nil
		}

		if math.Trunc(index) != index {
			return nil, fmt.Errorf("index %v invald in %v", index, obj)
		}

		intIndex := int(index)
		if 0 <= intIndex && intIndex < objr.Len() {
			rs := objr.Index(intIndex)
			return rs.Interface(), nil
		} else {
			return nil, fmt.Errorf("index %v out of range in %v", index, obj)
		}
	}

	// map
	if objr.Type().Kind() == reflect.Map {
		res := objr.MapIndex(reflect.ValueOf(key))
		return res.Interface(), nil
	}

	return nil, fmt.Errorf("object %v is not array or map", obj)
}
