package checkout

type Handler struct{}

type SKU struct{}

func (i *Handler) Checkout() error {
	panic("NYI")
}

func createPrices() ([]SKU, error) {
	panic("NYI")
}

func calculateTotal(skuPriceList []SKU, itemList string) int {
	panic("NYI")
}
