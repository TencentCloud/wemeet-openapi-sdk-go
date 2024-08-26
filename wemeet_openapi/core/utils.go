package wemeetcore

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// PtrBool is a helper routine that returns a pointer to given boolean value.
func PtrBool(v bool) *bool { return &v }

// PtrInt is a helper routine that returns a pointer to given integer value.
func PtrInt(v int) *int { return &v }

// PtrInt32 is a helper routine that returns a pointer to given integer value.
func PtrInt32(v int32) *int32 { return &v }

// PtrInt64 is a helper routine that returns a pointer to given integer value.
func PtrInt64(v int64) *int64 { return &v }

// PtrFloat32 is a helper routine that returns a pointer to given float value.
func PtrFloat32(v float32) *float32 { return &v }

// PtrFloat64 is a helper routine that returns a pointer to given float value.
func PtrFloat64(v float64) *float64 { return &v }

// PtrString is a helper routine that returns a pointer to given string value.
func PtrString(v string) *string { return &v }

// PtrTime is helper routine that returns a pointer to given Time value.
func PtrTime(v time.Time) *time.Time { return &v }

// PathValue is helper routine that returns a string to given path value.
func PathValue(v interface{}) string { return fmt.Sprint(v) }

// QueryValue is helper routine that returns a string to given query value.
func QueryValue(v interface{}) string {
	if vt := reflect.ValueOf(v); vt.Kind() == reflect.Ptr {
		if vt.IsNil() {
			return ""
		}
		v = vt.Elem().Interface()
	}
	return fmt.Sprint(v)
}

// ToString 使用反射获取值的类型
func ToString(value interface{}) string {
	// 使用反射获取值的类型
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(value).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(value).Uint(), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value.(bool))
	case string:
		return value.(string)
	default:
		return fmt.Sprintf("%v", value)
	}
}
