package factory_method

import "fmt"

type Factory interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

type FactoryHelper struct {
	Factory
}

func (s *FactoryHelper) Create(owner string) Product {
	p := s.CreateProduct(owner)
	s.RegisterProduct(p)
	return p
}

type Product interface {
	Use()
}

type IDCardFactory struct {
	owners []string
}

func NewIDCardFactory() *IDCardFactory {
	return &IDCardFactory{owners: make([]string, 0)}
}

func (s *IDCardFactory) CreateProduct(owner string) Product {
	return NewIDCard(owner)
}

func (s *IDCardFactory) RegisterProduct(product Product) {
	s.owners = append(s.owners, product.(*IDCard).owner)
}

type IDCard struct {
	owner string
}

func NewIDCard(owner string) Product {
	fmt.Println("制作 " + owner + "的 ID 卡")
	return &IDCard{owner: owner}
}

func (s *IDCard) Use() {
	fmt.Println("使用 " + s.owner + "的 ID 卡")
}
