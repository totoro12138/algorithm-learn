package main

import (
	"context"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var (
	listenerGin net.Listener = nil
	serverGin   *http.Server = nil
	gracefulGin              = flag.Bool("graceful", false, "listen on fd open 3(internal use only)")
)

func main() {
	var err error
	flag.Parse()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "gin_hello_world_3!")
	})

	serverGin = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if *gracefulGin {
		log.Println("listening on the existing file descriptor 3")
		f := os.NewFile(3, "")
		listenerGin, err = net.FileListener(f)
	} else {
		log.Println("listening on a new file descriptor")
		listenerGin, err = net.Listen("tcp", serverGin.Addr)
	}
	if err != nil {
		log.Fatalf("listener error: %v\n", err)
	}

	go func() {
		if err := serverGin.Serve(listenerGin); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error:%v\n", err)
		}
	}()

	handleSig()
	log.Println("signal end")
}

func handleSig() {
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-sign
		log.Printf("signal receive: %v\n", sig)
		//ctx, cancel := context.WithDeadline(context.Background(), time.Now())
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		//defer cancel()
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Println("shutdown")
			signal.Stop(sign)
			if err := serverGin.Shutdown(ctx); err != nil {
				log.Fatalf("service shutdown error:%v\n", err)
			}
			return
		case syscall.SIGUSR2:
			log.Println("reload")
			if err := reloadGin(); err != nil {
				log.Printf("service reload error: %v\n", err)
				continue
			}
			if err := serverGin.Shutdown(ctx); err != nil {
				log.Fatalf("service shutdown error:%v\n", err)
			}
			log.Println("service reload success")
			return
		}
	}
}

func reloadGin() error {
	tl, ok := listenerGin.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"-graceful"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{f}

	return cmd.Start()
}
