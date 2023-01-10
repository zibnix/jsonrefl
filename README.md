# A tiny pakcage for reflecting values from arbitrary JSON

You probably don't need this, and shouldn't use it. Ideally, your JSON has a static
structure that can be reliably marshaled into tagged structs. Use that!

If however, your JSON changes structure and you can't statically assume what the types
will be for certain keys, this may be helpful. It uses reflection though.

In addition, this project provides one version for pre go1.18, that doesn't use `any` or
generics, and another version for go1.18 and newer, that does use these features. Build
constraints control which version gets compiled, and tests exist for both.

In general, this is another bad idea. Go projects should just move to using the new
features and not bother with maintaining multiple versions, [see](https://github.com/golang/go/issues/52880).
This is why this module is marked as relying on go1.18, because there is no way to tell
the module system that it supports earlier versions of Go while containing source files
that use `any` and generics. I mostly did this as an experiment and to keep a visible
record of the older style.

```go
import "github.com/zibnix/jsonrefl"
```


# Go 1.18 and above:

## func FromArray
```go
func FromArray[T any](arr []any, index int) (T, error)
```
FromArray is a helper function for unpacking values out of
arbitrary JSON arrays. Inner objects are of type `map[string]any`.
Inner arrays are of type `[]any`. You can use `any` as the
type argument, but then you are really better off not using these
helpers, as they use reflection to try and match types.

## func FromObject
```go
func FromObject[T any](obj map[string]any, key string) (T, error)
```
FromObject is a helper function for unpacking values out of
arbitrary JSON objects. Inner objects are of type `map[string]any`.
Inner arrays are of type `[]any`. You can use `any` as the
type argument, but then you are really better off not using these
helpers, as they use reflection to try and match types.


# Go 1.17 and below:

## func FromArray
``` go
func FromArray(arr []interface{}, index int, value interface{}) error
```
FromArray is a helper function for unpacking values out of
arbitrary JSON arrays. `value` should be a pointer to a value of the type you
expect to retrieve. Inner objects are of type `map[string]interface{}`. Inner
arrays are of type `[]interface{}`. You can pass down a pointer to an
`interface{}`, but then you are really better off not using these helpers, as
they use reflection to try and match types.


## func FromObject
``` go
func FromObject(obj map[string]interface{}, key string, value interface{}) error
```
FromObject is a helper function for unpacking values out of
arbitrary JSON objects. `value` should be a pointer to a value of the type
you expect to retrieve. Inner objects are of type `map[string]interface{}`.
Inner arrays are of type `[]interface{}`. You can pass down a pointer to an
`interface{}`, but then you are really better off not using these helpers, as
they use reflection to try and match types.
