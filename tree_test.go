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
				` if (age != 29 || language[1] ) {language[0]} else {"world"}`,
				`golang`,
			},
			{
				`(2+4)*5`,
				`30`,
			},
			{
				` age != 29 || marry`,
				`true`,
			},
			{
				`test(29, name)`,
				`29,truman`,
			},
			{
				`if (isNaN(toNumber("29"))){"nan"}else{"num"}`,
				`num`,
			},
			{
				`1+1+"123"`,
				`2123`,
			},
			{
				`numMap[1] + "123"`,
				`val1123`,
			},
			{
				`numMap[3] + "123"`,
				`123`,
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
			"numMap": map[string]string{
				"1": "val1",
				"2": "val2",
			},
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

func BenchmarkExecuteContext_EvaluateExpression(b *testing.B) {

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
	for i := 0; i < b.N; i++ {
		ress, errs := Parse(`if (age != 29 || language[1] ) {language[0]} else {"world"}`)
		if errs != nil {
			b.Fatalf("err: %v", errs)
		}
		res := ress[0]
		val, err := ec.EvaluateExpression(res)
		if err != nil {
			b.Fatalf("err: %v", err)
		}
		if fmt.Sprintf("%v", val) != "golang" {
			b.Fatal(`expected "golang"`)
		}
	}
}
