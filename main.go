package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		//引数で受け取ったポートを利用するのでここでは指定しない
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	//別ゴルーチンでHTTPサーバを起動する
	eg.Go(func() error {
		// http.ErrServerClosedはhttp.Server.Shutdown()が正常に終了したことを示すので異常ではない
		if err := s.Serve(l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	//チャネルからの通知（終了通知）を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("faield to shutdown: %+v", err)
	}
	return eg.Wait()
}

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}
