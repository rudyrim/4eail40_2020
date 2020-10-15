package main

import (
	"fmt"
)

type Shape interface {
	fmt.Stringer
}

type Circle struct {
	r int
}

type Rectangle struct {
	Width  int
	Length int
}

func (r Rectangle) Area() int {
	return r.Width * r.Length
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle of width %d and length %d", r.Width, r.Length)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", c.r)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}
}
