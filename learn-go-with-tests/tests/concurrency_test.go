package tests

import (
	"reflect"
	"strings"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return true
	}
	return false
}

func TestCheckWebsites(t *testing.T) {
	type args struct {
		checker WebsiteChecker
		urls    []string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{"A", args{mockWebsiteChecker, []string{
			"https://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}}, map[string]bool{
			"https://google.com":         true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckWebsites(tt.args.checker, tt.args.urls)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckWebsites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		if i%2 == 0 {
			urls[i] = "http://" + strings.Repeat("-", 10)
		} else {
			urls[i] = "https://" + strings.Repeat("-", 10)
		}
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(mockWebsiteChecker, urls)
	}
}

func BenchmarkCheckWebsitesWithChannel(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		if i%2 == 0 {
			urls[i] = "http://" + strings.Repeat("-", 10)
		} else {
			urls[i] = "https://" + strings.Repeat("-", 10)
		}
	}

	for i := 0; i < b.N; i++ {
		CheckWebsitesWithChannel(mockWebsiteChecker, urls)
	}
}
