package checkout

import "testing"

func Test_calculateTotal(t *testing.T) {
	type args struct {
		skuPriceList []SKU
		itemList     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Given item list without special prices, total price calculates correctly",
			args: args{
				skuPriceList: []SKU{
					{
						Identifier:           "D",
						UnitPrice:            15,
						SpecialPriceQuantity: 0,
						SpecialPrice:         0,
					},
					{
						Identifier:           "C",
						UnitPrice:            20,
						SpecialPriceQuantity: 0,
						SpecialPrice:         0,
					},
				},
				itemList: "DDDDDCCC",
			},
			want: 135,
		}, {
			name: "Given item list with special prices, total price calculates correctly",
			args: args{
				skuPriceList: []SKU{
					{
						Identifier:           "A",
						UnitPrice:            50,
						SpecialPriceQuantity: 3,
						SpecialPrice:         130,
					},
					{
						Identifier:           "C",
						UnitPrice:            20,
						SpecialPriceQuantity: 0,
						SpecialPrice:         0,
					},
				},
				itemList: "AAAAACCC",
			},
			want: 290,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTotal(tt.args.skuPriceList, tt.args.itemList); got != tt.want {
				t.Errorf("calculateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
