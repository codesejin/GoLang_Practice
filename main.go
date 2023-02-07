package main

import (
	"codesejin/golearn/mydict"
	"fmt"
)

func main() {
	dictionary := mydict.Dictionary{}
	basWord := "hello"
	dictionary.Add(basWord, "First")
	dictionary.Search(basWord)
	dictionary.Delete(basWord)
	word, err := dictionary.Search(basWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word, err)
	}
}
