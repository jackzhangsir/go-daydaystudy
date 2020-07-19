package factory

import (
	"fmt"
	"testing"
)

func TestProductFactory_Create(t *testing.T) {

	product1 := &Product1{}
	product1.SetName("p1")
	fmt.Println(product1.GetName())

}

func TestProductFactory_Create2(t *testing.T) {
	factory := ProductFactory{}
	p1 := factory.Create(p1)
	p1.SetName("p1")
	fmt.Println(p1.GetName())

}
