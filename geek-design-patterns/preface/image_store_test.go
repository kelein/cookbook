package preface

import "testing"

func TestImageStore_Upload(t *testing.T) {
	type args struct {
		image  Image
		bucket string
	}
	tests := []struct {
		name  string
		args  args
		store ImageStore
	}{
		{"A", args{Image{"car", "png", 64}, "001"}, &AliyunImageStore{}},
		{"B", args{Image{"boy", "jpg", 128}, "002"}, &PrivateImageStore{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.store.Upload(tt.args.image, tt.args.bucket)
		})
	}
}

func TestPrivateImageStore_Download(t *testing.T) {
	type args struct {
		url   string
		store ImageStore
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"A", args{"http://localhost/aaa.png", &AliyunImageStore{}}, false},
		{"B", args{"http://localhost/bbb.jpg", &PrivateImageStore{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.store.Download(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateImageStore.Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
