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
		}
		for _, testCase := range testTable {
			ress := Parse(testCase.Input)
			So(len(ress), ShouldEqual, len(testCase.Output))

			for index, v := range ress {
				So(v.String(), ShouldEqual, testCase.Output[index])
			}
		}
	})
}
