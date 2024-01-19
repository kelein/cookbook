package singleton

import (
	"log/slog"
	"sync"
)

var (
	once      sync.Once
	intraCart *ShoppingCart
)

// ShoppingCart for shopping
type ShoppingCart struct {
	products map[string]int
}

// GetShoppingCart return a global shopping cart
func GetShoppingCart() *ShoppingCart {
	once.Do(func() {
		intraCart = &ShoppingCart{products: make(map[string]int)}
	})
	return intraCart
}

// AddProduct adds a product into the shopping cart
func (cart *ShoppingCart) AddProduct(name string, quantity int) {
	cart.products[name] += quantity
}

// ShowProduct show all products in the shopping cart
func (cart *ShoppingCart) ShowProduct() {
	for name, quantity := range cart.products {
		slog.Info("product", "name", name, "quantity", quantity)
	}
}
