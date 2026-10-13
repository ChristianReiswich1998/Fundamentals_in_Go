package main

import "fmt"

func CheckStock(inventory map[string]int, sku string) string {
	// Your code here
	result := ""

	value, ok := inventory[sku]
	if !ok {
		result = "unknown SKU"
	} else if value == 0 {
		result = "out of stock"
	} else if value >= 0 {
		result = fmt.Sprintf("%d in stock", value)
	}

	return result
}

func Discontinue(inventory map[string]int, sku string) bool {
	// Your code here
	var result bool

	if value, b := inventory[sku]; b {
		if !b == false {
			_ = fmt.Sprintf("Value: %s not found!", value)
			result = b

		} else if b == true {
			delete(inventory, sku)
			result = b
		}
	}

	return result
}

func main() {
	stock := map[string]int{
		"GO-SHIRT":   42,
		"GO-MUG":     0,
		"GO-STICKER": 100,
	}

	fmt.Println(CheckStock(stock, "GO-SHIRT"))
	fmt.Println(CheckStock(stock, "GO-MUG"))
	fmt.Println(CheckStock(stock, "GO-PLUSHIE"))

	fmt.Println(Discontinue(stock, "GO-MUG"))
	fmt.Println(Discontinue(stock, "GO-MUG"))
}
