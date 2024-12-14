package main

import "fmt"

func shop(work1, work2 bool) (bool, bool, bool) {
	buyTv50 := work1 && work2
	buyTv30 := work1 != work2 //xor in go
	buyIceCream := work1 || work2

	return buyTv50, buyTv30, buyIceCream
}

func main() {
	tv50, tv32, iceCream := shop(true, true)
	fmt.Printf("Tv 50: %t, Tv 32: %t, Icecream: %t , Healthy: %t",
		tv50, tv32, iceCream, !iceCream)
}
