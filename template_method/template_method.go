package template_method

import "fmt"

type Display interface {
	Open()
	Print()
	Close()
}

type AbstractDisplay struct {
	impl Display
}

func (s *AbstractDisplay) Display() {
	s.impl.Open()
	for i := 0; i < 5; i++ {
		s.impl.Print()
	}
	s.impl.Close()
}

type CharDisplay struct {
	char rune
	AbstractDisplay
}

func NewCharDisplay(char rune) *CharDisplay {
	s := &CharDisplay{char: char}
	s.impl = s
	return s
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
	AbstractDisplay
}

func NewStringDisplay(text string) *StringDisplay {
	s := &StringDisplay{text: text, with: len(text)}
	s.impl = s
	return s
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
