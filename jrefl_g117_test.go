//go:build !go1.18

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
var obj map[string]interface{}
var array = []interface{}{0.0, 1.0, 2.0}
var objArray = []interface{}{
	map[string]interface{}{
		"id": 0.0,
	},
	map[string]interface{}{
		"id": 1.0,
	},
	map[string]interface{}{
		"id": 2.0,
	},
}

func init() {
	if err := json.Unmarshal(jsonData, &obj); err != nil {
		panic(err)
	}
	log.Println("testing <1.18")
}

func TestFromObject(t *testing.T) {
	var err error

	var topStr string
	err = FromObject(obj, topKey, &topStr)
	tt.Expect(t, err, nil)
	tt.Expect(t, topStr, topVal)

	var arr []interface{}
	err = FromObject(obj, arrKey, &arr)
	tt.Expect(t, err, nil)
	tt.Expect(t, len(arr), 3)
}

func TestFromObjectBadVal(t *testing.T) {
	var err error

	var wrongType int = -1
	err = FromObject(obj, topKey, &wrongType)
	tt.NotNil(t, err)
	tt.Expect(t, wrongType, -1)

	var notPtr string = "oop"
	err = FromObject(obj, topKey, notPtr)
	tt.NotNil(t, err)
	tt.Expect(t, notPtr, "oop")
}

func TestFromObjectBadKey(t *testing.T) {
	var err error

	var topStr string
	err = FromObject(obj, "", &topStr)
	tt.NotNil(t, err)
	tt.Expect(t, topStr, "")

	err = FromObject(obj, "nope", &topStr)
	tt.NotNil(t, err)
	tt.Expect(t, topStr, "")
}

func TestFromNilObject(t *testing.T) {
	var err error

	var topStr string
	var obj map[string]interface{}
	err = FromObject(obj, topKey, &topStr)
	tt.NotNil(t, err)
	tt.Expect(t, obj, nil)
}

func TestEmptyInterfaceFromObject(t *testing.T) {
	var err error

	var topStr interface{}
	err = FromObject(obj, topKey, &topStr)
	tt.IsNil(t, err)
	tt.Expect(t, topStr, topVal)
}

func getArray(t *testing.T) []interface{} {
	var arr []interface{}
	if err := FromObject(obj, arrKey, &arr); err != nil {
		t.Fatal(err)
	}

	tt.Expect(t, arr, array)

	return arr
}

func getObjArray(t *testing.T) []interface{} {
	var arr []interface{}
	if err := FromObject(obj, objArrKey, &arr); err != nil {
		t.Fatal(err)
	}

	tt.Expect(t, arr, objArray)

	return arr
}

func TestFromArray(t *testing.T) {
	var err error
	arr := getArray(t)

	var i float64
	err = FromArray(arr, 1, &i)
	tt.Expect(t, err, nil)
	tt.Expect(t, i, array[1])
}

func TestObjectFromArray(t *testing.T) {
	var err error
	arr := getObjArray(t)

	var obj map[string]interface{}
	err = FromArray(arr, 1, &obj)
	tt.Expect(t, err, nil)
	tt.Expect(t, obj, objArray[1])
}

func TestFromArrayBadVal(t *testing.T) {
	var err error
	arr := getArray(t)

	var wrongType string = "oop"
	err = FromArray(arr, 1, &wrongType)
	tt.NotNil(t, err)
	tt.Expect(t, wrongType, "oop")

	var notPtr float64 = -1
	err = FromArray(arr, 1, notPtr)
	tt.NotNil(t, err)
	tt.Expect(t, notPtr, -1.0)
}

func TestFromArrayBadIndex(t *testing.T) {
	var err error
	arr := getArray(t)

	var ff float64 = -1
	err = FromArray(arr, -1, &ff)
	tt.NotNil(t, err)
	tt.Expect(t, ff, -1.0)

	ff = -1
	err = FromArray(arr, 5, &ff)
	tt.NotNil(t, err)
	tt.Expect(t, ff, -1.0)
}

func TestFromNilArray(t *testing.T) {
	var err error
	var arr []interface{}

	var i float64 = -1
	err = FromArray(arr, 1, &i)
	tt.NotNil(t, err)
	tt.Expect(t, i, -1.0)
}

func TestEmptyInterfaceFromArray(t *testing.T) {
	var err error
	arr := getArray(t)

	var ff interface{}
	err = FromArray(arr, 1, &ff)
	tt.IsNil(t, err)
	tt.Expect(t, ff, array[1])
}
