package main

import (
	"errors"
	"fmt"
	"log"
)

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

type VoucherAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}

type CoffeeMorningVoucherFactory struct{}

func (CoffeeMorningVoucherFactory) GetDrink() Drink { return CoffeeDrink{} }
func (CoffeeMorningVoucherFactory) GetFood() Food   { return Cake{} }

func GetVoucherAbstractFactory(campaignName string) (VoucherAbstractFactory, error) {
	if campaignName == "morning" {
		return CoffeeMorningVoucherFactory{}, nil
	}

	return nil, errors.New("coffee is not morning voucher")
}

func GetVoucher(factory VoucherAbstractFactory) Voucher {
	return Voucher{
		Food:  factory.GetFood(),
		Drink: factory.GetDrink(),
	}
}

func main() {
	voucherFactory, err := GetVoucherAbstractFactory("morning")

	if err != nil {
		fmt.Println(err)
	}

	voucher := GetVoucher(voucherFactory)

	log.Println(voucher)

}
