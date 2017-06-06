package simple_template

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExecuteContext_EvaluateExpression(t *testing.T) {
	Convey("test eval expression", t, func() {
		testTable := []*struct {
			Input  string
			Output string
		}{
			{
				`name`,
				`truman`,
			},
			{
				`age`,
				`29`,
			},
			{
				`marry`,
				`true`,
			},
			{
				`123`,
				`123`,
			},
			{
				`"456"`,
				`456`,
			},
			{
				`true`,
				`true`,
			},
			{
				`language[0] + "=>" + language[1]+ "=>" + language[2]`,
				`golang=>erlang=>lisp`,
			},
			{
				`age*2 + 2 - 10 + 16 / 2 * 2`,
				`66`,
			},
			{
				` if (age == 29) {"hello"} else {"world"}`,
				`hello`,
			},
			{
				` if (age != 29 && marry) {name} else {"world"}`,
				`world`,
			},
			{
				`test(29, name)`,
				`29,truman`,
			},
		}

		test := func(input []interface{}) (interface{}, error) {
			return fmt.Sprintf("%v,%v", input[0], input[1]), nil
		}
		ec := NewExecuteContenxt(map[string]interface{}{
			"name":     "truman",
			"age":      float64(29),
			"marry":    true,
			"language": []interface{}{"golang", "erlang", "lisp"},
		}, map[string]Func{
			"test": test,
		})
		//yyDebug = 99
		yyErrorVerbose = true

		for _, c := range testTable {
			ress, errs := Parse(c.Input)
			So(errs, ShouldBeNil)
			So(len(ress), ShouldBeGreaterThan, 0)
			lastRes := ress[len(ress)-1]
			val, err := ec.EvaluateExpression(lastRes)
			So(err, ShouldBeNil)
			So(fmt.Sprintf("%v", val), ShouldEqual, c.Output)
		}
	})
}
