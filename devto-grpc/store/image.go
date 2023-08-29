package store

import (
	"bytes"
	"sync"
)

// ImageStore store for image file
type ImageStore interface {
	Save(laptopID string, imageType string, imageData bytes.Buffer)
}

// DiskImageStore store image on disk
type DiskImageStore struct {
	lock   sync.RWMutex
	folder string
	images map[string]*ImageInfo
}

// ImageInfo stands for image meta info
type ImageInfo struct {
	LaptopID string
	Kind     string
	Path     string
}

// NewDiskImageStore create a new DiskImageStore instance
func NewDiskImageStore(folder string) *DiskImageStore {
	return &DiskImageStore{
		folder: folder,
		images: make(map[string]*ImageInfo),
	}
}

// Save store image on disk
func (d *DiskImageStore) Save() error { return nil }
