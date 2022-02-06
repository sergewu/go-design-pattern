package factory_method

import "testing"

func Test_FactoryMethod(t *testing.T) {
	var factory = NewIDCardFactory()
	var p1 = factory.Create("小明")
	var p2 = factory.Create("小红")
	var p3 = factory.Create("小刚")
	p1.Use()
	p2.Use()
	p3.Use()
}
