package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://caat"
}

func TestCheckWebsites(t *testing.T) {
	t.Run("basic check", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://bing.com",
			"waat://caat",
		}
		want := map[string]bool{
			"http://google.com": true,
			"http://bing.com":   true,
			"waat://caat":       false,
		}

		got := CheckWebsites(mockWebsiteChecker, websites)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}

	})
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
