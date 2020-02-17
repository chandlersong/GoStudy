package grammar

import (
	"testing"
)

func TestPointer(t *testing.T) {

	v := Object{value: 10, property: &Property{value: 11}}
	t.Logf("value is %d\n", v.value)
	p := &v
	t.Logf("pointer value is %d\n", p.value)

	var pointer *Object
	pointer = &v
	t.Logf("pointer value is %v\n", pointer.value)

	t.Logf("property value is %d\n", v.property.value)
	v.property.say(t)

	var o Object
	o = v
	t.Logf("property value is %d\n", o.value)

}

type Property struct {
	value int
}

func (p Property) say(t *testing.T) {
	t.Logf("hello")
}

type Object struct {
	value    int
	property *Property
}
