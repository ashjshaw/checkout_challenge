package checkout

type Handler struct{}

type SKU struct {
	Identifier           string `json:"identifier"`
	UnitPrice            int    `json:"unitPrice"`
	SpecialPriceQuantity int    `json:"specialPriceQuantity"`
	SpecialPrice         int    `json:"specialPrice"`
}

func (i *Handler) Checkout() error {
	panic("NYI")
}

func createPrices() ([]SKU, error) {
	panic("NYI")
}

func calculateTotal(skuPriceList []SKU, itemList string) int {
	panic("NYI")
}
