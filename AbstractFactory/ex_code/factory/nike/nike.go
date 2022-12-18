package nike

import (
	"example.com/design-pattern/AbstractFactory/ex_code/factory/interface"
)

type Nike struct {
}

func (n *Nike) MakeShoe() _interface.IShoe {
	return &NikeShoe{
		Shoe: _interface.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() _interface.IShirt {
	return &NikeShirt{
		Shirt: _interface.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
