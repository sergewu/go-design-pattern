package template_method

import "fmt"

type Display interface {
	Open()
	Print()
	Close()
}

type AbstractDisplay struct {
	Display
}

func (s *AbstractDisplay) Show() {
	s.Open()
	for i := 0; i < 5; i++ {
		s.Print()
	}
	s.Close()
}

type CharDisplay struct {
	char rune
}

func NewCharDisplay(char rune) Display {
	return &CharDisplay{char: char}
}

func (s *CharDisplay) Open() {
	fmt.Print("<<")
}

func (s *CharDisplay) Print() {
	fmt.Print(string(s.char))
}

func (s *CharDisplay) Close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	text string
	with int
}

func NewStringDisplay(text string) Display {
	return &StringDisplay{text: text, with: len(text)}
}

func (s *StringDisplay) Open() {
	s.printLine()
}

func (s *StringDisplay) Print() {
	fmt.Println("|" + s.text + "|")
}

func (s *StringDisplay) Close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.with; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
