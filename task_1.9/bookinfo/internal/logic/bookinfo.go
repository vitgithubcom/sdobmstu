package logic

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("\"%s\" by %s, %d", b.Title, b.Author, b.Year)
}