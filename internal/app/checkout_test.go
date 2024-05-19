package checkout

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
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
	type calls struct {
		readFileCalls  int
		unmarshalCalls int
	}
	tests := []struct {
		name    string
		i       *Handler
		calls   calls
		want    []SKU
		wantErr bool
	}{
		{
			name: "When given a valid priceList.json, no error is returned from the ReadFile function, or unmarshalFunction",
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
			want: []SKU{},
			calls: calls{
				readFileCalls:  1,
				unmarshalCalls: 1,
			},
			wantErr: false,
		}, {
			name: "When given an invalid priceList.json, an error is returned from the ReadFile function",
			i: &Handler{
				ReadFile: func(s string) ([]byte, error) {
					return []byte{}, errors.New("expected error in ReadFile")
				}, Unmarshal: func(b []byte, a any) error { return nil },
			},
			want: []SKU{},
			calls: calls{
				readFileCalls:  1,
				unmarshalCalls: 0,
			},
			wantErr: true,
		}, {
			name: "When given a valid JSON format, the information is unmarshalled correctly",
			i: &Handler{ReadFile: func(s string) ([]byte, error) {
				return []byte(`
				[{
					"identifier": "A",
					"unitPrice": 50,
					"specialPriceQuantity": 3,
					"specialPrice": 130
				}]`), nil
			},
				Unmarshal: json.Unmarshal,
			},
			calls: calls{
				readFileCalls:  1,
				unmarshalCalls: 1,
			},
			want: []SKU{
				{Identifier: "A",
					UnitPrice:            50,
					SpecialPriceQuantity: 3,
					SpecialPrice:         130,
				},
			},
			wantErr: false,
		}, {
			name: "Given an error from JSON unmarshaller, error is handled correctly",
			i: &Handler{ReadFile: func(s string) ([]byte, error) {
				return []byte(`
				[{
					"identifier": "A",
					"unitPrice": 50,
					"specialPriceQuantity": 3,
					"specialPrice": 130
				}]`), nil
			},
				Unmarshal: func(b []byte, a any) error { return errors.New("expected error in unmarshal") },
			},
			calls: calls{
				readFileCalls:  1,
				unmarshalCalls: 1,
			},
			want:    []SKU{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := calls{}
			i := &Handler{
				Unmarshal: func(b []byte, a any) error {
					calls.unmarshalCalls++
					return tt.i.Unmarshal(b, a)
				},
				ReadFile: func(s string) ([]byte, error) {
					calls.readFileCalls++
					return tt.i.ReadFile(s)
				},
			}
			got, err := i.createPrices()
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.createPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.calls, calls)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.createPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_Checkout(t *testing.T) {
	type calls struct {
		scanCalls int
	}
	tests := []struct {
		name    string
		i       *Handler
		calls   calls
		wantErr bool
	}{
		{
			name: "Given no error from createPrices, programme executes without issue",
			i: &Handler{Scanln: func(a ...any) (int, error) { return 0, nil },
				ReadFile:  func(s string) ([]byte, error) { return []byte{}, nil },
				Unmarshal: func(b []byte, a any) error { return nil },
			},
			calls:   calls{scanCalls: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		calls := calls{}
		i := &Handler{
			Scanln: func(a ...any) (int, error) {
				calls.scanCalls++
				return tt.i.Scanln(a)
			},
		}
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.calls, calls)
			if err := i.Checkout(); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Checkout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
