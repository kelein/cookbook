package singleton

import (
	"log/slog"
	"sync"
)

var once sync.Once

var intraCartManager *ShoppingCartManager

// ShoppingCartManager for shopping
type ShoppingCartManager struct {
	cart  map[string]int
	keys  []string
	mutex sync.Mutex
}

// GetShoppingCartManager return a global shopping cart
func GetShoppingCartManager() *ShoppingCartManager {
	once.Do(func() {
		intraCartManager = &ShoppingCartManager{
			cart: make(map[string]int),
		}
	})
	return intraCartManager
}

// AddProduct adds a product into the shopping cart
func (m *ShoppingCartManager) AddProduct(name string, quantity int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.cart[name] += quantity
	if _, ok := m.cart[name]; !ok {
		m.keys = append(m.keys, name)
	}
}

// ShowProduct show all products in the shopping cart
func (m *ShoppingCartManager) ShowProduct() {
	for _, item := range m.keys {
		slog.Info("product", "name", item, "quantity", m.cart[item])
	}
}
