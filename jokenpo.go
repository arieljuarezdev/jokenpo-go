package main

import (
	"fmt"
	"math/rand"
)

var computer int
var player int

func combat(c int, p int) {

	var comb_val int = (c + p) % 2

	fmt.Println("-----------")
	fmt.Println("p =>", p)
	fmt.Println("c =>", c)
	fmt.Println("comb_val =>", comb_val)

	if p != c {
		if comb_val == 1 {
			if p > c {
				fmt.Println("-------------")
				fmt.Println("player wins!!")
			} else {
				fmt.Println("-------------")
				fmt.Println("computer wins!!")
			}
		} else {
			if p < c {
				fmt.Println("-------------")
				fmt.Println("player wins!!")
			} else {
				fmt.Println("-------------")
				fmt.Println("computer wins!!")
			}
		}
	} else {
		fmt.Println("------")
		fmt.Println("Empate")
	}
}
func main() {
	min := 0
	max := 3

	computer = rand.Intn(max-min) + min

	fmt.Println("Selecione uma das opções:")
	fmt.Println("0 - Pedra")
	fmt.Println("1 - Papel")
	fmt.Println("2 - Tesoura")
	fmt.Print("Digite aqui: ")
	fmt.Scan(&player)

	combat(computer, player)

	// Test Area
	// fmt.Println("mod: Pe + Pa:", (0+1)%2)
	// fmt.Println("mod: Pe + Te:", (0+2)%2)
	// fmt.Println("mod: Pa + Te:", (1+2)%2)
}
