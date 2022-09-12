package main

import (
	"fmt"
	"sync"
	"time"
)

func cat(fishCH, catCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("cat")
			catCH <- struct{}{}
			<-fishCH
		}
	}()
}

func dog(catCH, dogCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			<-catCH
			fmt.Println("dog")
			dogCH <- struct{}{}
		}
	}()

}
func fish(dogCH, fishCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			<-dogCH
			fmt.Println("fish")
			time.Sleep(time.Second)
			fishCH <- struct{}{}
		}
	}()
}

func main() {
	catCH := make(chan struct{})
	dogCH := make(chan struct{})
	fishCH := make(chan struct{})

	wg := sync.WaitGroup{}

	cat(fishCH, catCH, &wg)
	dog(catCH, dogCH, &wg)
	fish(dogCH, fishCH, &wg)

	wg.Wait()
}
