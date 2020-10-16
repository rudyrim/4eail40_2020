package main

import (
	"fmt"
)

func main() {
	const(
	ligne int = 2
	colonne int =6
	)
	var tab = [ligne][colonne] string {{"♙","♘","♗","♖","♕","♔"},{"♟","♞","♝","♜","♛","♚"}}
	fmt.Println(tab[1][5])
}