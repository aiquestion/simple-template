package simple_template

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Lexer(t *testing.T) {
	Convey("test lexer", t, func() {
		testTable := []*struct {
			Input  string
			Output []string
		}{
			{
				`if(event["did"] == "123456"){ 123 // test comment
				 }else{ 789 }`,
				[]string{`<if(<<<identxpr,event>[<strexpr,did>]>,==,<strexpr,123456>>)then{<numexpr,123>}else{<numexpr,789>}>`},
			},
			{
				`did`,
				[]string{`<identxpr,did>`},
			},
			{
				`did[1][2]["test"]`,
				[]string{`<<<<identxpr,did>[<numexpr,1>]>[<numexpr,2>]>[<strexpr,test>]>`},
			},
			{
				`{1, 3, 4, 5}`,
				[]string{`<array(<numexpr,1>,<numexpr,3>,<numexpr,4>,<numexpr,5>)>`},
			},
			{
				`{
				"name" : 123,
				"age" : 234,
				sleep: true
				}`,
				[]string{`<map(<strexpr,name>=><numexpr,123>,<strexpr,age>=><numexpr,234>,<identxpr,sleep>=><true expr>)>`},
			},
		}
		for _, testCase := range testTable {
			ress, errs := Parse(testCase.Input)
			So(errs, ShouldBeNil)
			So(len(ress), ShouldEqual, len(testCase.Output))

			for index, v := range ress {
				So(v.String(), ShouldEqual, testCase.Output[index])
			}
		}
	})
}
