package simple_template

import (
	"fmt"
	"math"
)

func toNumber(inputs []interface{}) (interface{}, error) {
	if inputs == nil || len(inputs) == 0 {
		return math.NaN(), fmt.Errorf("call toNumber with no argument")
	}
	if res, ok := TryNumber(inputs[0]); ok {
		return res, nil
	} else {
		return math.NaN(), nil
	}
}

func isNaN(inputs []interface{}) (interface{}, error) {
	if inputs == nil || len(inputs) == 0 {
		return false, fmt.Errorf("call isNaN with no argument")
	}
	if res, ok := TryNumber(inputs[0]); ok {
		return math.IsNaN(res), nil
	} else {
		return true, nil
	}
}

func toBool(inputs []interface{}) (interface{}, error) {
	if inputs == nil || len(inputs) == 0 {
		return false, fmt.Errorf("call toBool with no argument")
	}
	if res, ok := TryBool(inputs[0]); ok {
		return res, nil
	} else {
		return false, nil
	}
}
