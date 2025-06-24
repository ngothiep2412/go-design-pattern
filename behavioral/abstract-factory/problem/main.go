package main

import "fmt"

type Drink interface {
	Drink()
}

type Food interface {
	Food()
}

type CoffeeDrink struct{}

func (CoffeeDrink) Drink() {
	fmt.Println("Coffee is drinking")
}

type Cake struct{}

func (Cake) Food() {
	fmt.Println("Cake is drinking")
}

type Voucher struct {
	Drink
	Food
}

func main() {
	fmt.Println([]Voucher{
		Voucher{
			Drink: CoffeeDrink{},
			Food:  Cake{},
		},
		Voucher{
			Drink: CoffeeDrink{},
		},
	})
}
