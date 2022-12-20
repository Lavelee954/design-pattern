package main

import "example.com/design-pattern/Observer/ex_code/subject"

func main() {

	shirtItem := subject.NewItem("Nike Shirt")

	observerFirst := subject.Customer{Id: "abc@gmail.com"}
	observerSecond := subject.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(&observerFirst)
	shirtItem.Register(&observerSecond)

	shirtItem.UpdateAvailability()
}
