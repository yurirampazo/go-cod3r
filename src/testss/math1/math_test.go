package math

import "testing"


const errorDefault = "expected value %v, result %v"

func TestAverage(t *testing.T) {
	expectedValue := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9) //change 7.2 to  7.28 to see error

	if value != expectedValue {
		t.Errorf(errorDefault, expectedValue, value)
	}



}