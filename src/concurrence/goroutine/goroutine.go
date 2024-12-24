package main

import (
	"fmt"
	"time"
)

func speak(person, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s %s (iteração) %d\n", person, text, i + 1)
	}
}

func main() {
	// speak("Maria", "Why dont you speak to me?", 3)
	// speak("Gabriel", "I can only speak after you?", 1)

	// go speak("Maria", "Heyy", 5)
	// go speak("Gabriel", "Hoo", 5)

	// time.Sleep(time.Second * 5)
	// fmt.Println("End")

	go speak("Maria", "Hola que tal?", 5)
	speak("John", "Bien, y tu?", 2)
}