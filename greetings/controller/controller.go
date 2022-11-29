package controller

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Oi, %v. Seja Bem-vindo!", name)
	return message
}
