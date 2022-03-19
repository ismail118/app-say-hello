package main

import (
	"fmt"
	go_say_hello "github.com/ismail118/go-say-hello"

	go_say_hellov2 "github.com/ismail118/go-say-hello/v2"
)

/*
	for more
	see documentation https://docs.gofiber.io/
*/

func main() {
	fmt.Println(go_say_hello.SayHello())
	fmt.Println(go_say_hellov2.SayHello("Ismail"))
}
