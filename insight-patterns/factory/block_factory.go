package factory

import (
	"log/slog"
	"sync"
)

var once sync.Once

// Block of abstract
type Block interface {
	Show()
}

// CircleBlock concrete circle block
type CircleBlock struct{ Block }

// Show display block info
func (c *CircleBlock) Show() {
	slog.Info("Circle Block")
}

// SquareBlock concrete square block
type SquareBlock struct{ Block }

// Show display block info
func (s *SquareBlock) Show() {
	slog.Info("Square Block")
}

// Factory of block
type Factory interface {
	Product() Block
}

var _ Factory = (*CircleBlockFactory)(nil)

var intraCircleBlockFactory *CircleBlockFactory

// CircleBlockFactory product circle blocks
type CircleBlockFactory struct{ Factory }

// GetCircleBlockFactory return a singleton instance of CircleBlockFactory
func GetCircleBlockFactory() Factory {
	once.Do(func() {
		intraCircleBlockFactory = new(CircleBlockFactory)
	})
	return intraCircleBlockFactory
}

// Product create circle blocks
func (c *CircleBlockFactory) Product() Block {
	return new(CircleBlock)
}

var _ Factory = (*SquareBlockFactory)(nil)

var intraSquareBlockFactory *SquareBlockFactory

// SquareBlockFactory product square blocks
type SquareBlockFactory struct{ Factory }

// GetSquareBlockFactory return a singleton instance of SquareBlockFactory
func GetSquareBlockFactory() Factory {
	once.Do(func() {
		intraSquareBlockFactory = new(SquareBlockFactory)
	})
	return intraSquareBlockFactory
}

// Product create square blocks
func (s *SquareBlockFactory) Product() Block {
	return new(SquareBlock)
}

// BlockFactory product blocks
type BlockFactory struct{}

// GetFactory return
func (f *BlockFactory) GetFactory(kind string) Factory {
	if kind == "circle" {
		return GetCircleBlockFactory()
	}
	if kind == "square" {
		return GetSquareBlockFactory()
	}
	return nil
}

// ProductBlock product block of specified kind
func (f *BlockFactory) ProductBlock(kind string) Block {
	return f.GetFactory(kind).Product()
}
