package preface

import "log"

func init() {
	log.SetPrefix("‣ ")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
}

// Image defination
type Image struct {
	Name string
	Kind string
	Size int64
}

// ImageStore of abstract
type ImageStore interface {
	Upload(image Image, bucket string) error
	Download(url string) error
}

// AliyunImageStore concrete image store of Aliyun
type AliyunImageStore struct{}

// Upload of AliyunImageStore
func (s *AliyunImageStore) Upload(image Image, bucket string) error {
	log.Printf("[AliyunImageStore] uploading image <%s.%s>", image.Name, image.Kind)
	return nil
}

// Download of AliyunImageStore
func (s *AliyunImageStore) Download(url string) error {
	log.Printf("[AliyunImageStore] downloading image from %q", url)
	return nil
}

// PrivateImageStore concrete image store of Personal
type PrivateImageStore struct{}

// Upload of PrivateImageStore
func (p *PrivateImageStore) Upload(image Image, bucket string) error {
	log.Printf("[PrivateImageStore] uploading image <%s.%s>", image.Name, image.Kind)
	return nil
}

// Download of PrivateImageStore
func (p *PrivateImageStore) Download(url string) error {
	log.Printf("[PrivateImageStore] downloading image from %q", url)
	return nil
}
