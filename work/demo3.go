package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "heloo"
		e
		f = 100
		g
		h = iota
	)
	fmt.Println(a, b, c, d, e, f, g, h)
}
