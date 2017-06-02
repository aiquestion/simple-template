package simple_template

import (
	"encoding/json"
	"strconv"
)

func TryString(v interface{}) (string, bool){
	switch tv := v.(type) {
	case string:
		return tv, true
	case []byte:
		return string(tv), true
	case int64:
		return strconv.FormatInt(int64(tv), 10), true
	case uint64:
		return strconv.FormatUint(uint64(tv), 10), true
	case int32:
		return strconv.FormatInt(int64(tv), 10), true
	case uint32:
		return strconv.FormatUint(uint64(tv), 10), true
	case int:
		return strconv.Itoa(int(tv)), true
	case json.Number:
		return tv.String(), true
	}
	return "", false
}

func TryBool(v interface{}) (bool, bool) {
	switch tv := v.(type) {
	case bool:
		return bool(tv), true
	case string:
		return tv != "", true
	case []byte:
		return string(tv) != "", true
	case int64:
		return tv != 0, true
	case uint64:
		return tv != 0, true
	case int32:
		return tv != 0, true
	case uint32:
		return tv != 0, true
	case int:
		return tv != 0, true
	}
	return false, false
}

// TODO TryInt

