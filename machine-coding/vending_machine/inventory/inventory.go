package inventory

import "fmt"

type Inventory map[string]ProductDetails

var inventoryInfo Inventory

func init() {
	inventoryInfo = map[string]ProductDetails{
		"beverage_001": {
			Cost:  10,
			Count: 5,
		},
		"beverage_002": {
			Cost:  8,
			Count: 8,
		},
		"chips_001": {
			Cost:  12,
			Count: 10,
		},
	}
}

type ProductDetails struct {
	Cost  int
	Count int
}

func View() string {
	return fmt.Sprintf("%+v", inventoryInfo)
}

func GetProductDetails(productId string) (ProductDetails, bool) {
	val, ok := inventoryInfo[productId]
	return val, ok
}
