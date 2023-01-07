//go:build go1.18

package jsonrefl

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/zibnix/tt"
)

const (
	topKey    = "key"
	topVal    = "val"
	arrKey    = "array"
	objArrKey = "obj_array"
	idKey     = "id"
)

var jsonData = []byte(fmt.Sprintf(`{"%s":"%s","%s":[0,1,2],"%s":[{"%s":0}, {"%s":1}, {"%s":2}]}`, topKey, topVal, arrKey, objArrKey, idKey, idKey, idKey))
var obj map[string]any
var array = []any{0.0, 1.0, 2.0}
var objArray = []any{
	map[string]any{
		"id": 0.0,
	},
	map[string]any{
		"id": 1.0,
	},
	map[string]any{
		"id": 2.0,
	},
}

func init() {
	if err := json.Unmarshal(jsonData, &obj); err != nil {
		panic(err)
	}
	log.Println("testing >=go1.18")
}

func TestFromObject(t *testing.T) {
	val, err := FromObject[string](obj, topKey)
	tt.Expect(t, err, nil)
	tt.Expect(t, val, topVal)

	arr, err := FromObject[[]any](obj, arrKey)
	tt.Expect(t, err, nil)
	tt.Expect(t, len(arr), 3)
}

func TestFromObjectBadVal(t *testing.T) {
	wt, err := FromObject[int](obj, topKey)
	tt.NotNil(t, err)
	tt.Expect(t, wt, 0)

	wt2, err := FromObject[[]any](obj, topKey)
	tt.NotNil(t, err)
	tt.Expect(t, len(wt2), 0)
	tt.IsNil(t, wt2)
}

func TestFromObjectBadKey(t *testing.T) {
	str, err := FromObject[string](obj, "")
	tt.NotNil(t, err)
	tt.Expect(t, str, "")

	str, err = FromObject[string](obj, "nope")
	tt.NotNil(t, err)
	tt.Expect(t, str, "")
}

func TestFromNilObject(t *testing.T) {
	var obj map[string]any
	str, err := FromObject[string](obj, topKey)
	tt.NotNil(t, err)
	tt.Expect(t, str, "")
	tt.Expect(t, obj, nil)
}

func TestEmptyInterfaceFromObject(t *testing.T) {
	str, err := FromObject[any](obj, topKey)
	tt.IsNil(t, err)
	tt.Expect(t, str, topVal)
}

func getArray(t *testing.T) []any {
	arr, err := FromObject[[]any](obj, arrKey)
	if err != nil {
		t.Fatal(err)
	}

	tt.Expect(t, arr, array)

	return arr
}

func getObjArray(t *testing.T) []any {
	arr, err := FromObject[[]any](obj, objArrKey)
	if err != nil {
		t.Fatal(err)
	}

	tt.Expect(t, arr, objArray)

	return arr
}

func TestFromArray(t *testing.T) {
	arr := getArray(t)

	f, err := FromArray[float64](arr, 1)
	tt.Expect(t, err, nil)
	tt.Expect(t, f, array[1])
}

func TestObjectFromArray(t *testing.T) {
	arr := getObjArray(t)

	obj, err := FromArray[map[string]any](arr, 1)
	tt.Expect(t, err, nil)
	tt.Expect(t, obj, objArray[1])
}

func TestFromArrayBadVal(t *testing.T) {
	arr := getArray(t)

	wt, err := FromArray[string](arr, 1)
	tt.NotNil(t, err)
	tt.Expect(t, wt, "")
}

func TestFromArrayBadIndex(t *testing.T) {
	arr := getArray(t)

	ff, err := FromArray[float64](arr, -1)
	tt.NotNil(t, err)
	tt.Expect(t, ff, 0.0)

	ff, err = FromArray[float64](arr, 5)
	tt.NotNil(t, err)
	tt.Expect(t, ff, 0.0)
}

func TestFromNilArray(t *testing.T) {
	var arr []any

	ff, err := FromArray[float64](arr, 1)
	tt.NotNil(t, err)
	tt.Expect(t, ff, 0.0)
}

func TestEmptyInterfaceFromArray(t *testing.T) {
	arr := getArray(t)

	ff, err := FromArray[any](arr, 1)
	tt.IsNil(t, err)
	tt.Expect(t, ff, array[1])
}
