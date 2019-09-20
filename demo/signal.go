package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var start = struct {
	Stop bool
	WG   sync.WaitGroup
}{
	Stop: false,
	WG:   sync.WaitGroup{},
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println()
			fmt.Println("exit")
			start.Stop = true
			start.WG.Wait()
			//time.Sleep(time.Second * 5)
			os.Exit(0)
		}
	}()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if start.Stop != false {
			_, _ = writer.Write([]byte("404"))
			return
		}
		start.WG.Add(1)
		defer start.WG.Done()
		time.Sleep(time.Second * 3)
		_, _ = writer.Write([]byte("success"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
