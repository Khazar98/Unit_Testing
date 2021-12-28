package math

import "testing"

type TestData struct {
	a, b   int
	result int
}

func TestAdd(t *testing.T) {
	testdata := []TestData{
		{1, 2, 3},
		{2, 6, 8},
		{23, 6, 29},
		{99, 7, 106},
	}
	for _, data := range testdata {
		result := Add(data.a, data.b)
		if result != data.result {
			t.Errorf("Add(%d,%d) Failed. Expected %d got %d\n", data.a, data.b, data.result, result)
		}
	}

	// result := Add(2, 3)
	// if result != 5 {
	// 	t.Errorf("Add(2,3) Failed. Expected %d, got %d \n", 5, result)
	// } else {
	// 	t.Logf("Add(2,3) Passed. Expectd %d, got %d\n", 5, result)
	// }
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	if result != 5 {
		t.Errorf("Divide(10,2) FAILED. Expected %f, got %f\n", 5.0, result)
	} else {
		t.Logf("Divide(10,2) PASSED.Expected %f", 5.0)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	if result != 20 {
		t.Errorf("Failed. Result is not equal to %d", result)
	} else {
		t.Logf("PASSED ")

	}

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestSubtract(t *testing.T) {
	result := Subtract(12, 3)
	if result != 9 {
		t.Errorf("Failed. Result is not equal to %d ", result)
	} else {
		t.Logf("Passed")
	}
}
