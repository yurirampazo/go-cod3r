package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	if i := randomNumber(); i > 5 {
		fmt.Println("Won")
	} else {
		fmt.Println("Lost")
	}

}


func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}