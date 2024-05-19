package checkout

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
	panic("NYI")
}

func (i *Handler) createPrices() ([]SKU, error) {
	panic("NYI")
}

func (i *Handler) calculateTotal(skuPriceList []SKU, itemList string) int {
	panic("NYI")
}
