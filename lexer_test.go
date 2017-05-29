package simple_template

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Lexer(t *testing.T) {
	Convey("test lexer", t, func() {
		yyDebug = 0
		yyErrorVerbose = true
		code := `if(event["did"] == "123456"){123}else{789}`
		res := Parse(code)
		fmt.Printf("%+v\r\n", res)
	})
}
