package main

import (
	"fmt"
	"time"
)

var dog = make(chan struct{})
var cat = make(chan struct{})
var fish = make(chan struct{})

func Dog() {
	<-fish
	fmt.Println("Dog")
	dog <- struct{}{}
}
func Cat() {
	<-dog
	fmt.Println("Cat")
	cat <- struct{}{}
}
func Fish() {
	<-cat
	fmt.Println("Fish")
	fish <- struct{}{}
}

func main() {
	for i := 0; i < 100; i++ {
		go Dog()
		go Cat()
		go Fish()
	}
	fish <- struct{}{}
	time.Sleep(10 * time.Second)
}
