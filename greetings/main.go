package main

import (
	"fmt"

	"github.com/antoniogeroncio/golang/greetings/controller"
)

func main() {
	message := controller.Hello("Antonio Geroncio")
	fmt.Println(message)
}
