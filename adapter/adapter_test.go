package adapter

import "testing"

func TestPrintAdapter(t *testing.T) {
	var p = NewBannerPrint("hello world")
	p.PrintWeak()
	p.PrintStrong()
}
