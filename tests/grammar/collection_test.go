package grammar

import "testing"

func TestDefineArray(t *testing.T) {
	array1 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		t.Log(array1[i])
	}

	var array2 [5]int
	for i := 0; i < 5; i++ {
		array2[i] = i
		t.Log(array2[i])
	}

	array3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		t.Log(array3[i])
	}

	array4 := [5]int{1: 10, 2: 20}
	for i := 0; i < 5; i++ {
		t.Log(array4[i])
	}
}
