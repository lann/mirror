package mirror

import "reflect"

func Convert(i interface{}, protoType interface{}) interface{} {
	return ConvertType(i, reflect.TypeOf(protoType))
}

func ConvertType(i interface{}, typ reflect.Type) interface{} {
	return reflect.ValueOf(i).Convert(typ).Interface()
}

func ForEach(s interface{}, f func(interface{})) {
	ForEachValue(s, func(val reflect.Value) {
		f(val.Interface())
	})
}

func ForEachValue(s interface{}, f func(reflect.Value)) {
	val := reflect.ValueOf(s)
	
	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		panic(&reflect.ValueError{Method: "mirror.ForEach", Kind: val.Kind()})
	}

	l := val.Len()
	for i := 0; i < l; i++ {
		f(val.Index(i))
	}
}
