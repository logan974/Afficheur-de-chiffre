package main

// Import

import (
	"fmt"
	"time"
)

// Variable de longueur de l'affichage des chiffre

var c1, c2, c3, c4, c5, c6, c7, c8, c9 []int

func initialisation() {
	c1 = append(c1, 1, 1, 1, 1, 1, 1, 1)
	c2 = append(c2, 0, 1, 1, 1, 0, 1, 1)
}

func affichage(c []int) {
	//l := len(c)
	fmt.Print("\x0c")
	for i := 0; i < 5; i++ {
		if i == 0 || i == 2 || i == 4 {
			for j := 0; j < 3; j++ {
				switch {
				case c[j] == 1 && j == 1:
					fmt.Print(" ████ ")
				case c[j] == 1 && (j == 0 || j == 2):
					fmt.Print(" ")
				case c[j] == 0 && (j == 0 || j == 2):
					fmt.Print(" ")
				default:
					fmt.Print("      ")
				}
			}
			fmt.Print("\n")
		} else {
			for i := 0; i < 2; i++ {
				for j := 0; j < 3; j++ {
					switch {
					case c[j] == 1 && (j == 0 || j == 2):
						fmt.Print("█")
					case c[j] == 0 && (j == 0 || j == 2):
						fmt.Print(" ")
					default:
						fmt.Print("      ")
					}

				}
				fmt.Print("\n")
			}
		}
	}
	fmt.Print("\n")

}

func main() {
	initialisation()
	affichage(c1)
	time.Sleep(time.Second)
	affichage(c2)

}
