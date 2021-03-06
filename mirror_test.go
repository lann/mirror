package mirror

import (
	"reflect"
	"testing"
)

type X struct {
	x int
}

type Y X

func TestConvert(t *testing.T) {
	x := X{x: 1}
	expected := Y(x)
	y := Convert(x, Y{}).(Y)
	if y != expected {
		t.Errorf("expected %v, got %v", expected, y)
	}
}

func TestConvertType(t *testing.T) {
	x := X{x: 1}
	expected := Y(x)
	y := ConvertType(x, reflect.TypeOf(Y{})).(Y)
	if y != expected {
		t.Errorf("expected %v, got %v", expected, y)
	}
}

func TestForEach(t *testing.T) {
	ints := []int{1, 2, 3}
	out := []int{0, 0, 0}
	ForEach(ints, func(i int, v interface{}) {
		out[i] = v.(int)
	})
	if !reflect.DeepEqual(out, ints) {
		t.Errorf("expected %v, got %v", ints, out)
	}
}

func TestIsArrayOrSlice(t *testing.T) {
	test := func(expect bool, v interface{}) {
		if IsArrayOrSlice(v) != expect {
			t.Errorf("expected IsArrayOrSlice(%T) == %v", v, expect)
		}
	}

	test(true, []interface{}{})
	test(true, []int{})
	test(true, reflect.ValueOf([]int{}))
	test(true, reflect.Array)
	test(false, 1)
	test(false, reflect.ValueOf(1))
	test(false, reflect.Map)
}

func TestForEachValue(t *testing.T) {
	ints := []int{0, 0}
	expected := []int{1, 1}
	oneVal := reflect.ValueOf(1)
	ForEachValue(reflect.ValueOf(ints), func(_ int, val reflect.Value) {
		val.Set(oneVal)
	})
	if !reflect.DeepEqual(ints, expected) {
		t.Errorf("expected %v, got %v", expected, ints)
	}
}

func TestForEachValuePanic(t *testing.T) {
	var panicVal *reflect.ValueError
	func() {
		defer func() { panicVal = recover().(*reflect.ValueError) }()
		ForEachValue(reflect.ValueOf(X{}), func(_ int, _ reflect.Value) {})
	}()
	if panicVal == nil {
		t.Errorf("expected panic, didn't")
	}
}
