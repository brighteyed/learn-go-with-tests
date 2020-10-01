package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://qwe.rty" {
		return false
	}

	return true
}

func slowStubWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"https://www.google.com",
		"https://yandex.ru",
		"waat://qwe.rty",
	}

	got := CheckWebsites(mockWebsiteChecker, urls)

	want := map[string]bool{
		"https://www.google.com": true,
		"https://yandex.ru":      true,
		"waat://qwe.rty":         false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
