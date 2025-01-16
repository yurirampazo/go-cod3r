package math

import "testing"


const errorDefault = "expected value %v, result %v"

func TestAverage(t *testing.T) {
	expectedValue := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9) //change 7.2 to  7.28 to see error

	if value != expectedValue {
		t.Errorf(errorDefault, expectedValue, value)
	}


	/*
	TO run tests with coverage we should use:
	go test -cover -> result: 
	
	PASS
  coverage: 100.0% of statements
  ok      your-module-name/src/testss/math1       0.005s
	


	But too see more details we can do something like:
	go test -coverprofile=result.out

	this will create a file with an outrput that could be read with some options:

	option 1:
	go tool cover -func=result.out

	it'll give the methods percentage

	or you can create a html with the report:
	go tool cover -html=result.out


	*/


}