package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	exit := make(chan bool)

	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println()
			fmt.Println("exit")

			time.Sleep(time.Second * 5)
			exit <- true
		}
	}()

	<-exit
}
