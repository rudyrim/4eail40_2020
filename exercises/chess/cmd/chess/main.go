package main

import (
	"fmt"
)

var (
	a = [8]string{"♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"}
	b = [8]string{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"}
	c = [8]string{" ", " ", " ", " ", " ", " ", " ", " "}
	d = [8]string{" ", " ", " ", " ", " ", " ", " ", " "}
	e = [8]string{" ", " ", " ", " ", " ", " ", " ", " "}
	f = [8]string{" ", " ", " ", " ", " ", " ", " ", " "}
	g = [8]string{"♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"}
	h = [8]string{"♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"}
)

func display() {
	fmt.Println(" ", " 0", "1", "2", "3", "4", "5", "6", "7")
	fmt.Println("A", a, "A")
	fmt.Println("B", b, "B")
	fmt.Println("C", c, "C")
	fmt.Println("D", d, "D")
	fmt.Println("E", e, "E")
	fmt.Println("F", f, "F")
	fmt.Println("G", g, "G")
	fmt.Println("H", h, "H")
	fmt.Println(" ", " 0", "1", "2", "3", "4", "5", "6", "7")
}
func main() {

	e[1] = "♙"
	b[1] = " "
	display()

}
