package main

import (
	"encoding/json"
	"fmt"
	"os"

	checkout "github.com/ashjshaw/checkout_challenge/internal/app"
)

func main() {
	i := checkout.Handler{ReadFile: os.ReadFile, Scanln: fmt.Scanln, Unmarshal: json.Unmarshal}
	err := i.Checkout()
	if err != nil {
		fmt.Println(err)
	}
}
