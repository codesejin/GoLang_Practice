package main

import (
	"fmt"
	"time"
)

// Channel 은 goroutine이랑 메인함수 사이에 정보를 전달하기 위한 방법
func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "sejin", "mommy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
