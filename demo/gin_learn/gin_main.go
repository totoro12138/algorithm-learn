package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	listenerGin *net.Listener = nil
	serverGin *http.Server = nil
	graceful
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "gin hello world!")
	})

	f := os.NewFile(3,"")

	serverGin = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := serverGin.Serve(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error:%v\n", err)
		}
	}()

	handleSig()
}

func handleSig() {
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-sign
		log.Printf("signal receive: %v\n",sig)
		//ctx, cancel := context.WithDeadline(context.Background(), time.Now())
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		switch sig {
		case syscall.SIGINT,syscall.SIGTERM:
			log.Println("shutdown")
			signal.Stop(sign)
			if err := serverGin.Shutdown(ctx);err != nil {
				log.Fatalf("service shutdown error:%v\n",err)
			}
			return
		case syscall.SIGUSR2:
			log.Println("reload")
			if err := reloadGin();err != nil {
				log.Printf("service reload error: %v\n",err)
				continue
			}
			if err := serverGin.Shutdown(ctx);err != nil {
				log.Fatalf("service shutdown error:%v\n",err)
			}
			log.Println("service reload success")
			return
		}
	}
}

func reloadGin() error {
	tl,ok := listenerGin.(*net.TCPListener)

}
