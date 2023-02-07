package main

import (
	"fmt"
	"time"
)

// Channel 은 goroutine이랑 메인함수 사이에 정보를 전달하기 위한 방법
func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	result := <-c
	fmt.Println(result)
	fmt.Println(<-c)

}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
