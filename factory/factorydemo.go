package factory

type Product interface {
	SetName(name string)
	GetName() string
}

type Product1 struct {
	name string
}

func (p1 *Product1)SetName(name string)  {
	p1.name = name
}
func (p1 *Product1)GetName() string {
	return p1.name
}

type Product2 struct {
	name string
}

func (p2 *Product2)SetName(name string)  {
	p2.name = name
}
func (p2 *Product2)GetName() string {
	return p2.name
}


type productType int
const(
	p1 productType = 1
	p2 productType = 2
)

type ProductFactory struct {
}

func (pf ProductFactory)Create(type1 productType) Product {
	switch type1 {
	case p1:
		return &Product1{}
	case p2:
		return &Product2{}
	}
	return nil

}