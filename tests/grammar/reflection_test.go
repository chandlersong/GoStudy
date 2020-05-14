package grammar

import (
	"fmt"
	"reflect"
	"testing"
)

type Item struct {
	name string
}

func TestReflectionDemos(t *testing.T) {

	t.Run("how to new", func(t *testing.T) {
		var item Item
		var myType reflect.Type
		myType = reflect.TypeOf(item)

		var myValue reflect.Value
		myValue = reflect.New(myType)

		var b Item
		b = myValue.Elem().Interface().(Item)
		fmt.Printf("b is %v \n", b)
	})
}
