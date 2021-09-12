package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
        return false
    }
    return true
}
func slowWebsiteChecker(url string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
        "http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
    }

	actualRes := CheckWebsites(mockWebsiteChecker, websites)
	want := len(websites)
    got := len(actualRes)
    if want != got {
        t.Fatalf("Wanted %v, got %v", want, got)
    }
	expectedResults := map[string]bool{
        "http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
    }
	if !reflect.DeepEqual(expectedResults, actualRes) {
        t.Fatalf("Wanted %v, got %v", expectedResults, actualRes)
    }
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
        urls[i] = "a url"
    }
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		CheckWebsites(slowWebsiteChecker, urls)
	}
}