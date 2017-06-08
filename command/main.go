package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"simple-template"
)

func main() {
	router := gin.Default()
	router.StaticFile("/play.html", "./index.html")
	router.POST("/run", func(c *gin.Context) {
		code := c.PostForm("code")
		defines := c.PostForm("define")
		if code == "" {
			c.String(http.StatusOK, "code empty")
			return
		}
		definesMap := map[string]interface{}{}
		if defines != "" {
			err := json.Unmarshal([]byte(defines), &definesMap)
			if err != nil {
				c.String(http.StatusOK, fmt.Sprintf("define err : %v", err))
				return
			}
		}
		ec := simple_template.NewExecuteContenxt(definesMap, map[string]simple_template.Func{})
		exprs, errs := simple_template.Parse(code)
		if (errs != nil && len(errs) > 0) || exprs == nil || len(exprs) == 0 {
			c.String(http.StatusOK, fmt.Sprintf("syntax err : %v, or expression empty", errs))
			return
		}
		result, err := ec.EvaluateExpression(exprs[len(exprs)-1])
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("evaluate err : %v", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%v", result))
	})

	router.Run(":8080")
}
