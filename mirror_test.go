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
	out := []int{}
	ForEach(ints, func(i interface{}) {
		out = append(out, i.(int))
	})
	if !reflect.DeepEqual(out, ints) {
		t.Errorf("expected %v, got %v", ints, out)
	}
}

func TestForEachValue(t *testing.T) {
	ints := []int{0, 0}
	expected := []int{1, 1}
	oneVal := reflect.ValueOf(1)
	ForEachValue(ints, func(val reflect.Value) {
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
		ForEachValue(X{}, func(_ reflect.Value) {})
	}()
	if panicVal == nil {
		t.Errorf("expected panic, didn't")
	}
}
