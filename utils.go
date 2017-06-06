package simple_template

import "strconv"

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
		return num, err != nil
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
		return strconv.FormatFloat(tv, 'f', -1, 0), true
	case string:
		return tv, true
	}
	return "", false
}
