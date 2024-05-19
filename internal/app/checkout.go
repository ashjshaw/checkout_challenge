package checkout

import (
	"fmt"
	"strings"
)

type Handler struct {
	ReadFile  func(string) ([]byte, error)
	Scanln    func(...any) (int, error)
	Unmarshal func([]byte, any) error
}

type SKU struct {
	Identifier           string `json:"identifier"`
	UnitPrice            int    `json:"unitPrice"`
	SpecialPriceQuantity int    `json:"specialPriceQuantity"`
	SpecialPrice         int    `json:"specialPrice"`
}

func (i *Handler) Checkout() error {
	skuPriceList, _ := i.createPrices()
	var userInput string
	fmt.Println("Please input scanned items from the Price list, eg:'BCAABDA' | invalid inputs will be ignored.")
	i.Scanln(&userInput)
	fmt.Println(calculateTotal(skuPriceList, userInput))
	return nil
}

func (i *Handler) createPrices() ([]SKU, error) {
	skuPriceList := []SKU{}
	pricesJson, err := i.ReadFile("../priceList.json")
	if err != nil {
		return []SKU{}, fmt.Errorf("error occured reading priceList.json: %w", err)
	}
	err = i.Unmarshal(pricesJson, &skuPriceList)
	if err != nil {
		return skuPriceList, fmt.Errorf("error occured unmarshalling json: %w", err)
	}
	return skuPriceList, nil
}

func calculateTotal(skuPriceList []SKU, itemList string) int {
	totalPrice := 0
	for _, sku := range skuPriceList {
		skuQuantity := strings.Count(itemList, sku.Identifier)
		if sku.SpecialPriceQuantity > 0 {
			totalPrice += (skuQuantity / sku.SpecialPriceQuantity) * sku.SpecialPrice
			skuQuantity = skuQuantity % sku.SpecialPriceQuantity
		}
		totalPrice += skuQuantity * sku.UnitPrice
	}
	return totalPrice
}
