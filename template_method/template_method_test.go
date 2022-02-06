package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {
	d1 := NewCharDisplay('H')
	d1.Display()

	d2 := NewStringDisplay("Hello World.")
	d2.Display()
}
