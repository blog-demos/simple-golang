package main

import (
	"encoding/json"
	"fmt"
)

// Product 商品信息
type Product struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

func main() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	var p1 Product
	json.Unmarshal([]byte(data), &p1)
	fmt.Println(p1)

	s1 := `{"Name":"Xiao mi 6", "ProductID":1, "Price":2499, "Number":10000, "IsOnSale":true}`;
	var p2 Product
	json.Unmarshal([]byte(s1), &p2)
	fmt.Println(p2)
}
