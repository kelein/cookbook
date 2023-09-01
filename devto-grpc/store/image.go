package store

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

// ImageStore store for image file
type ImageStore interface {
	Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error)
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
func (d *DiskImageStore) Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error) {
	imageID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("gen image uuid id failed: %w", err)
	}

	path := fmt.Sprintf("%s/%s.%s", d.folder, imageID, imageType)
	file, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("create image file error: %w", err)
	}
	if _, err := imageData.WriteTo(file); err != nil {
		return "", fmt.Errorf("write image file error: %w", err)
	}

	d.lock.Lock()
	defer d.lock.Unlock()
	d.images[imageID.String()] = &ImageInfo{
		LaptopID: laptopID,
		Kind:     imageType,
		Path:     path,
	}
	return imageID.String(), nil
}
