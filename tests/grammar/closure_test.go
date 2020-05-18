package grammar

import "testing"

func TestClosure(t *testing.T) {

	t.Run("hello world", func(t *testing.T) {
		helloWorldOuter(t, helloWorld)
	})

	t.Run("return value", func(t *testing.T) {
		t.Log(helloWorldOuterWithReturn(hello))
	})
}

func helloWorldOuter(t *testing.T, fn func(t *testing.T)) {
	fn(t)
}

func helloWorldOuterWithReturn(fn func() string) string {
	return fn()
}

func helloWorld(t *testing.T) {
	t.Log("hello World \n")
}

func hello() string {
	return "hello World with return value"
}
