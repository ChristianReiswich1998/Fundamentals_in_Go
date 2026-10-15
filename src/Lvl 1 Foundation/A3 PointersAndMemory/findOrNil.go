package main

import "fmt"

type Item struct {
	Name  string
	Stock int
}

// LowestStock returns a pointer to the item with the least stock, so the
// caller can restock it by adding to Stock. The returned pointer must point
// into the slice, so the caller's write reaches the real item.
func LowestStock(items []Item) *Item {
	// Your code here
	lowestStock := items[0].Stock

	for i := 1; i < len(items); {
		if items[i].Stock < lowestStock {
			lowestStock = items[i].Stock
			return &items[i]
		}

		return &items[0]
	}

	return nil
}

func main() {
	items := []Item{
		{Name: "USB cable", Stock: 40},
		{Name: "HDMI cable", Stock: 6},
		{Name: "Power adapter", Stock: 22},
	}

	if it := LowestStock(items); it != nil {
		it.Stock += 50
	}

	fmt.Println(items[1].Stock) // want 56, prints 6 until LowestStock is fixed
}
