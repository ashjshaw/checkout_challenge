package checkout

import (
	"reflect"
	"testing"
)

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
		}, {
			name: "Input with invalid characters are disregarded",
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
				itemList: "AAAAACCC42bob",
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

func TestHandler_createPrices(t *testing.T) {
	tests := []struct {
		name    string
		i       *Handler
		want    []SKU
		wantErr bool
	}{
		{
			name: "When given a valid priceList.json, no error is returned from the ReadFile function",
			i: &Handler{ReadFile: func(s string) ([]byte, error) {
				return []byte(`
				[{
					"identifier": "A",
					"unitPrice": 50,
					"specialPriceQuantity": 3,
					"specialPrice": 130
				}]`), nil
			},
				Unmarshal: func(b []byte, a any) error { return nil },
			},
			want:    []SKU{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.createPrices()
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.createPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.createPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
