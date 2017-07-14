package simple_template

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

/**
only have 3 types in this template language: string/bool/float64
*/

func TryBool(v interface{}) (bool, bool) {
	if v == nil {
		return false, true
	}
	switch tv := v.(type) {
	case bool:
		return bool(tv), true
	case string:
		return tv != "", true
	case float64:
		return tv != 0, true
	}
	return false, false
}

func TryNumber(v interface{}) (float64, bool) {
	if v == nil {
		return 0, true
	}
	switch tv := v.(type) {
	case bool:
		if tv {
			return 1, true
		} else {
			return 0, true
		}
	case float64:
		return tv, true

	case string:
		num, err := strconv.ParseFloat(tv, 0)
		return num, err == nil
	}
	return 0, false
}

func TryString(v interface{}) (string, bool) {
	if v == nil {
		return "", true
	}
	switch tv := v.(type) {
	case bool:
		if tv {
			return "1", true
		} else {
			return "0", true
		}
	case float64:
		return strconv.FormatFloat(tv, 'f', -1, 64), true
	case string:
		return tv, true
	}
	return "", false
}

func ToParserValue(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	switch vt := v.(type) {
	case string, float64, bool:
		return vt, nil
	case json.Number:
		vti, err := vt.Int64()
		if err == nil {
			return float64(vti), nil
		}
		vtf, err := vt.Float64()
		if err == nil {
			return vtf, nil
		}
		return nil, err
	case []byte:
		return string(vt), nil
	case int:
		return float64(vt), nil
	case int8:
		return float64(vt), nil
	case int16:
		return float64(vt), nil
	case int32:
		return float64(vt), nil
	case int64:
		return float64(vt), nil
	case uint:
		return float64(vt), nil
	case uint8:
		return float64(vt), nil
	case uint16:
		return float64(vt), nil
	case uint32:
		return float64(vt), nil
	case uint64:
		return float64(vt), nil
	case float32:
		return float64(vt), nil
	}

	val := reflect.ValueOf(v)
	if val.Type().Kind() == reflect.Array || val.Type().Kind() == reflect.Slice {
		// loop all val
		res := make([]interface{}, 0, val.Len())
		for i := 0; i < val.Len(); i++ {
			finalVal, err := ToParserValue(val.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			res = append(res, finalVal)
		}
		return res, nil
	}

	if val.Type().Kind() == reflect.Map {
		if val.Type().Key().Kind() != reflect.String {
			return nil, fmt.Errorf("we only allow string key map")
		}
		res := make(map[string]interface{})
		for _, k := range val.MapKeys() {
			keyStr, ok := k.Interface().(string)
			if !ok {
				return nil, fmt.Errorf("you should never reach here")
			}
			finalVal, err := ToParserValue(val.MapIndex(k).Interface())
			if err != nil {
				return nil, err
			}
			res[keyStr] = finalVal
		}
		return res, nil

	}

	return nil, fmt.Errorf("cannot parse value %+v", v)
}
