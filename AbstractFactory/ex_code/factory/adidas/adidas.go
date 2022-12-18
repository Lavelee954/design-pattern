package adidas

import (
	"example.com/design-pattern/AbstractFactory/ex_code/factory/interface"
)

type Adidas struct {
}

func (a *Adidas) MakeShoe() _interface.IShoe {
	return &AdidasShoe{
		Shoe: _interface.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() _interface.IShirt {
	return &AdidasShirt{
		Shirt: _interface.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
