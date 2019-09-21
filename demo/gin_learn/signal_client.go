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
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * 2)
		wg.Add(1)
		go Client(i, &wg)
	}
	wg.Wait()
}

func Client(idx int, wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
		wg.Done()
	}()
	var i int
	for {
		resp, err := http.Get("http://192.168.1.192:8080")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("[", idx, "]", i, ":", string(bytes))
		i++
		time.Sleep(time.Millisecond * 100)
	}

}
