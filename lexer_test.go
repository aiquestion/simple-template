package simple_template

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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
				[]string{`<<<<identxpr,did>[<numexpr,1>]>[<numexpr,2>]>[<strexpr,test>]>`},
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
