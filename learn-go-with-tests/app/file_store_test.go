package app

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_tape_Write(t *testing.T) {
	type args struct {
		name string
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"A", args{"a.txt", "ABC"}, "ABC"},
	}
	for _, tt := range tests {
		temp, _ := createTempFile(t, tt.args.name)
		// defer clean()

		t.Run(tt.name, func(t *testing.T) {
			tr := &tape{file: temp}
			_, err := tr.Write([]byte(tt.args.text))
			if err != nil {
				t.Errorf("tape.Write() error = %v", err)
				return
			}

			temp.Seek(0, 0)
			content, _ := ioutil.ReadAll(temp)
			got := string(content)
			if got != tt.want {
				t.Errorf("tape.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	temp, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("create temp file failed: %v", err)
	}

	temp.Write([]byte(initialData))
	clean := func() { os.Remove(temp.Name()) }
	return temp, clean
}
