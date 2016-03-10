package main

import (
	"fmt"
	"time"
	//"math/rand"	
    "strings"

)

func say(s string) {
	for i := 0; i < 4; i++ {
		time.Sleep(77 * time.Millisecond)
		
		//rand.NewSource(time.Now() )
		fmt.Println(s)
	}
}

func e1(s string){
	//time.Sleep(150 * time.Millisecond)
    fmt.Printf("Number of e's in a secret string: ")
    fmt.Printf("%d\n", strings.Count(s, "e"))
}

func e2(s string){
	//time.Sleep(33 * time.Millisecond)
    fmt.Printf("Number of p's in a secret string 2: ")
    fmt.Printf("%d\n", strings.Count(s, "p"))

}

func main() {
	words := "Lets count the number of e's in this string"
	wordmore :="And here we have anotehr secret string. yeap."
	
	//rand.Seed(time.Now().Unix())
	
	go say("Let's run this in thread 1")
	say("And here comes thread 2")
	
	for l := 0; l<3; l++ {
		go e1(words)
		e2(wordmore)
	}
}
