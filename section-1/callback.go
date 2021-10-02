package main

import "fmt"

type profanity func(string) string

func main() {
	sentence := greet("hafidz", checkProfanity)
	profanity := greet("anjing", checkProfanity)
	profanity2 := greet("bangsat", checkProfanity)
	fmt.Println(sentence, "<<<< safe sentence")
	fmt.Println(profanity, "<<<< unsafe sentence")
	fmt.Println(profanity2, "<<<< unsafe sentence")

}

func checkProfanity(name string) string {
	if name == "anjing" || name == "bangsat" {
		return "..."
	}
	return name
}

func greet(name string, callback profanity) string {
	return "hello " + callback(name) + ", nice to meet you"
}
