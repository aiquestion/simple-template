package simple_template

import (
	"testing"

	"reflect"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToParserValue(t *testing.T) {
	Convey("test", t, func() {
		testTable := []*struct {
			Input  interface{}
			Output interface{}
		}{
			{
				int(1),
				float64(1),
			},
			{
				float32(1),
				float64(1),
			},
			{
				[]byte("test"),
				"test",
			},
			{
				true,
				true,
			},
			{
				[]int{1, 2, 3},
				[]interface{}{float64(1), float64(2), float64(3)},
			},
			{
				[]string{"1", "2", "3"},
				[]interface{}{"1", "2", "3"},
			},
			{
				map[string]string{
					"key1": "value1",
					"key2": "2",
				},
				map[string]interface{}{
					"key1": "value1",
					"key2": "2",
				},
			},
			{
				map[string]interface{}{
					"key1": "value1",
					"key2": 2,
				},
				map[string]interface{}{
					"key1": "value1",
					"key2": float64(2),
				},
			},
		}

		for _, c := range testTable {
			res, err := ToParserValue(c.Input)
			So(err, ShouldBeNil)
			if reflect.ValueOf(res).Type().Kind() == reflect.Array || reflect.ValueOf(res).Type().Kind() == reflect.Slice {
				ra, ok := c.Output.([]interface{})
				So(ok, ShouldBeTrue)
				for i := 0; i < reflect.ValueOf(res).Len(); i++ {
					ar := reflect.ValueOf(res).Index(i).Interface()
					So(ar, ShouldEqual, ra[i])
					So(ar, ShouldHaveSameTypeAs, ra[i])
				}
				continue
			}
			if reflect.ValueOf(res).Type().Kind() == reflect.Map {
				So(reflect.DeepEqual(c.Output, res), ShouldBeTrue)
				continue
			}
			So(res, ShouldEqual, c.Output)
			So(res, ShouldHaveSameTypeAs, c.Output)
		}
	})
}
