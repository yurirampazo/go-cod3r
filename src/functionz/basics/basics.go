package main


import "fmt"

func fun() {
	fmt.Println("Function 1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
} 


func f4(p1, p2 string) string {
	return "f4: " + p1 + p2 
}


func f5() (string, string, int) {
	return "this is", "my return", 1
}

func main() {
	fun()
	f2("hi, my name is", "slim shady")
	fmt.Println(f3())
	stringz := f4("Hey", "Yo")
	fmt.Println(stringz)
	test, _ , number:= f5()
	fmt.Println(test, number)
}

