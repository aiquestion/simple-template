package simple_template

import "fmt"

type ExecuteContext struct {

}

func EvaluateExpression(expr Expr)(interface{}, error){
	return nil, nil
}

func evalIdentityExpr(e *IdentExpr) (interface{},error) {
	// need
	return nil, nil
}

func evalLogicalExpr(e *LogicalExpr) (bool, error) {
	left, err := EvaluateExpression(e.Lhs)
	if err != nil { return false, err }

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
		right, err := EvaluateExpression(e.Rhs)
		if err != nil { return false, err }
		rightBool, ok := TryBool(right)
		if !ok {
			return false, fmt.Errorf("right expr value unknown, %v", left)
		}
		if e.Operator == "||" {
			return leftBool || rightBool, nil
		}else{
			return leftBool && rightBool, nil
		}
	}

	right, err := EvaluateExpression(e.Rhs)
	if err != nil { return false, err }
	switch e.Operator {
	case "==":
		fallthrough
	case "!=":
		leftStr, ok := TryString(left)
		if !ok {
			return false, fmt.Errorf("left expr value unknown, %v", left)
		}
		rightStr, ok := TryString(right)
		if !ok {
			return false, fmt.Errorf("right expr value unknown, %v", left)
		}
		eq := leftStr == rightStr
		if e.Operator == "==" {
			return eq, nil
		}else{
			return !eq, nil
		}
	case ">":
		fallthrough
	case "<":
		fallthrough
	case ">=":
		fallthrough
	case "<=":
		// TODO

	}

	// TODO
	return false, nil
}



