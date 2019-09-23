package keyword

import "testing"

func doubleDefer(t *testing.T) {
	defer t.Log("deffer 1")
	defer t.Log("deffer 2")
	defer t.Log("deffer 3")

	t.Log("run func")
}

func TestInterface(t *testing.T) {
	doubleDefer(t)
}
