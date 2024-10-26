// Task 1. Написать код функции, которая делает merge N каналов.
// Весь входной поток перенаправляется в один канал

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://eëeë",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	urlch := make(chan string)

	gourl := func(url string) {
		defer wg.Done()
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			urlch <- fmt.Sprintf("адресс %s -not ok", url)
		} else {
			urlch <- fmt.Sprintf("адресс %s - ok", url)
		}
		if resp != nil {
			resp.Body.Close() // Avoid a potential panic if resp is nil.
		}

	}

	for _, url := range urls {
		go gourl(url)
	}
	go func() {
		wg.Wait()
		close(urlch)
	}()
	for n := range urlch {
		fmt.Println(n)
	}
}
