package adapter

import "fmt"

type Banner struct {
	Title string
}

func (b *Banner) ShowWithParen() string {
	return "(" + b.Title + ")"
}
func (b *Banner) ShowWithAster() string {
	return "*" + b.Title + "*"
}

type Printer interface {
	PrintWeak()
	PrintStrong()
}

type BannerPrint struct {
	banner *Banner
}

func NewBannerPrint(title string) Printer {
	return &BannerPrint{&Banner{title}}
}

func (b *BannerPrint) PrintWeak() {
	fmt.Println(b.banner.ShowWithParen())
}

func (b *BannerPrint) PrintStrong() {
	fmt.Println(b.banner.ShowWithAster())
}
