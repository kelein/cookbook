package convert

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/kelein/cookbook/devto-grpc/repo"
	"github.com/kelein/cookbook/devto-grpc/service"
)

func TestWriteBinaryFile(t *testing.T) {
	type args struct {
		message  proto.Message
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"A", args{service.NewLaptop(), "../tests/laptop.bin"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteBinaryFile(tt.args.message, tt.args.filename)
			t.Logf("WriteBinaryFile() output = %v", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteBinaryFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadBinaryFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{"A", "../tests/laptop.1", false},
		{"B", "../tests/laptop.2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origin := service.NewLaptop()
			if err := WriteBinaryFile(origin, tt.filename); err != nil {
				t.Logf("WriteBinaryFile() error = %v", err)
			}

			laptop := &repo.Laptop{}
			err := ReadBinaryFile(tt.filename, laptop)
			t.Logf("ReadBinaryFile() output = %v", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBinaryFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !proto.Equal(origin, laptop) {
				t.Errorf("ReadBinaryFile() got different proto messages")
			}
		})
	}
}

func TestWriteJSONFile(t *testing.T) {
	type args struct {
		message  proto.Message
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		{"A", args{service.NewLaptop(), "../tests/laptop.json"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteJSONFile(tt.args.message, tt.args.filename); err != nil {
				t.Errorf("WriteJSONFile() error = %v", err)
			}
		})
	}
}
