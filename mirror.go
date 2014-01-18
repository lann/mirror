package mirror

import "reflect"

func Convert(i interface{}, protoType interface{}) interface{} {
	return ConvertType(i, reflect.TypeOf(protoType))
}

func ConvertType(i interface{}, typ reflect.Type) interface{} {
	return reflect.ValueOf(i).Convert(typ).Interface()
}

type kindly interface {
	Kind() reflect.Kind
}

func IsArrayOrSlice(i interface{}) bool {
	var kind reflect.Kind
	switch v := i.(type) {
	case []interface{}:
		return true
	case reflect.Kind:
		kind = v
	case kindly:
		kind = v.Kind()
	default:
		kind = reflect.TypeOf(i).Kind()
	}
	return kind == reflect.Array || kind == reflect.Slice
}

func ForEach(s interface{}, f func(int, interface{})) {
	ForEachValue(reflect.ValueOf(s), func(i int, val reflect.Value) {
		f(i, val.Interface())
	})
}

func ForEachValue(val reflect.Value, f func(int, reflect.Value)) {
	if !IsArrayOrSlice(val) {
		panic(&reflect.ValueError{Method: "mirror.ForEach", Kind: val.Kind()})
	}

	l := val.Len()
	for i := 0; i < l; i++ {
		f(i, val.Index(i))
	}
}
