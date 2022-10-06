package main

import (
	"fmt"
	"sync"
)

var dog = make(chan struct{})
var cat = make(chan struct{})
var fish = make(chan struct{})

func Dog(fish, dog chan struct{}, wg *sync.WaitGroup) {
	go func() {
		for {
			defer wg.Done()
			dog <- struct{}{}
			fmt.Println("Dog")
			<-fish
			// dog <- struct{}{}
		}

	}()
}
func Cat(dog, cat chan struct{}, wg *sync.WaitGroup) {
	go func() {
		for {
			defer wg.Done()
			<-dog
			fmt.Println("Cat")
			cat <- struct{}{}
		}

	}()
}
func Fish(cat, fish chan struct{}, wg *sync.WaitGroup) {
	go func() {
		for {
			defer wg.Done()
			<-cat
			fmt.Println("Fish")
			fish <- struct{}{}
		}

	}()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)

	Dog(fish, dog, wg)
	Cat(dog, cat, wg)
	Fish(cat, fish, wg)

	// fish <- struct{}{}
	wg.Wait()
}
