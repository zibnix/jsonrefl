//go:build go1.18

package jsonrefl

import (
	"reflect"

	"github.com/pkg/errors"
)

// FromObject is a helper function for unpacking values out of
// arbitrary JSON objects. Inner objects are of type `map[string]any`.
// Inner arrays are of type `[]any`. You can use `any` as the
// type argument, but then you are really better off not using these
// helpers, as they use reflection to try and match types.
func FromObject[T any](obj map[string]any, key string) (T, error) {
	var t T

	if obj == nil {
		return t, errors.New("attempt to get value from a nil JSON object")
	}

	if len(key) == 0 {
		return t, errors.New("attempt to pull value for empty key from JSON object")
	}

	val, gotVal := obj[key]
	if !gotVal {
		return t, errors.Errorf("no value found for key: %s", key)
	}

	return getVal[T](val)
}

// FromArray is a helper function for unpacking values out of
// arbitrary JSON arrays. Inner objects are of type `map[string]any`.
// Inner arrays are of type `[]any`. You can use `any` as the
// type argument, but then you are really better off not using these
// helpers, as they use reflection to try and match types.
func FromArray[T any](arr []any, index int) (T, error) {
	var t T

	if arr == nil {
		return t, errors.New("attempt to get value from a nil JSON aray")
	}

	if index < 0 || index >= len(arr) {
		return t, errors.Errorf("provided index %d was out of range", index)
	}

	val := arr[index]

	return getVal[T](val)
}

func getVal[T any](val any) (T, error) {
	var t T

	ty := reflect.TypeOf(t)

	if ty != nil && !reflect.TypeOf(val).AssignableTo(ty) {
		return t, errors.Errorf(
			"provided reference is to a value of type %T that cannot be assigned to type found in JSON: %T.",
			t, val)
	}

	t = val.(T)
	return t, nil
}
