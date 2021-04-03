package main

// Category enum
type Category int

const (
	pop Category = iota
	jazz
	rock
)

// Song type
type Song struct {
	title    string
	category Category
	likes    int
}
