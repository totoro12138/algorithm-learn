package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 2)
		wg.Add(1)
		go Client(&wg)
	}
	wg.Wait()
}

func Client(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
		wg.Done()
	}()
	var i int
	for {
		resp, err := http.Get("http://127.0.0.1:8080")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(i, ":", string(bytes))
		i++
		time.Sleep(time.Microsecond * 100)
	}

}
