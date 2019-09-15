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

	array5 := [5]*int{new(int), new(int), new(int), new(int), new(int)}
	for i := 0; i < 5; i++ {
		*array5[i] = i
		t.Log(*array5[i])
	}
}

func TestCreateSlice(t *testing.T) {

	slice1 := make([]int, 5)
	for i := 0; i < 5; i++ {
		t.Log(slice1[i])
	}

	slice2 := []int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		t.Log(slice2[i])
	}

	slice3 := []string{5: "hello world"}
	for i := 0; i < 6; i++ {
		t.Log(slice3[i])
	}

	newSlice := slice2[2:3] // include 2. exclude 3
	t.Log(cap(newSlice))
	t.Log(newSlice)

	/**
	two share slice share the value
	*/
	slice2[2] = 8
	t.Log(newSlice[0])
}

func TestSliceAppend(t *testing.T) {

	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[1:3]

	newSlice = append(newSlice, 60)

	t.Logf("old slice is %d, the value of 3 has changed", slice)
	t.Logf("new slice is %d", newSlice)

	biggerSlice := slice[1:3]
	biggerSlice = append(biggerSlice, 160)
	t.Logf("cap of bigger Slice: %d", cap(biggerSlice))
	biggerSlice = append(biggerSlice, 170)
	t.Logf("cap of bigger Slice: %d", cap(biggerSlice))
	biggerSlice = append(biggerSlice, 180, 190)
	t.Logf("cap of bigger Slice: %d", cap(biggerSlice))
	biggerSlice = append(biggerSlice, 100)
	t.Logf("cap of bigger Slice: %d", cap(biggerSlice))

	biggerSlice[0] = 110

	t.Logf("old slice is %d, the value of 3 has changed", slice)
	t.Logf("new slice is %d", newSlice)
	t.Logf("bigger slice is %d, the slice should not share the same array with others when the cap is bigger than others", biggerSlice)

	for key, value := range biggerSlice {
		t.Logf("index is %d,value is %d", key, value)
	}
}
