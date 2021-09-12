package concurrency

type WebChecker func(string) bool
type result struct {
    string
    bool
}
func CheckWebsites(wc WebChecker,urls []string) map[string]bool {
	len := len(urls)
	res := make(map[string]bool, len)
	resChan := make(chan *result)
	for _, url := range urls {
		// res[url] = wc(url)
		go func(u string) {
			resChan <- &result{u, wc(u)}
		}(url)
	}
	for i:= 0; i < len; i++ {
		result := <- resChan
		res[result.string] = result.bool
	}
	return res
}
