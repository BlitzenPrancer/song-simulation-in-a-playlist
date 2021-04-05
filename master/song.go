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

func (c Category) toString() string {
	switch c {
	case jazz:
		return "jazz";
	case pop:
		return "pop";
	case rock:
		return "rock";
	default:
		return "";
	}
}
