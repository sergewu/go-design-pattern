package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {
	char := NewCharDisplay('H')
	d1 := &AbstractDisplay{char}
	d1.Show()

	str := NewStringDisplay("Hello World.")
	d2 := &AbstractDisplay{str}
	d2.Show()
}
