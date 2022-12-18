package main

//	Abstract Factory
//	관련 객체들의 구상 클래스들을 지정하지 않고도 관련 객체들의 모음을 생성할 수 있도록 하는 생성패턴

import (
	"example.com/design-pattern/AbstractFactory/ex_code/factory/adidas"
	"example.com/design-pattern/AbstractFactory/ex_code/factory/interface"
	"example.com/design-pattern/AbstractFactory/ex_code/factory/nike"
	"fmt"
)

func GetSportsFactory(brand string) (_interface.ISportsFactory, error) {
	if brand == "adidas" {
		return &adidas.Adidas{}, nil
	}

	if brand == "nike" {
		return &nike.Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s _interface.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s _interface.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()

}
