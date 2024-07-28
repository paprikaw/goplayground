package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	StatusCode int
	err        error
}

func main() {
	ctx := context.Background()
	err, _ := helloHanlder(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello, World!\n")
}

func helloHanlder(ctx context.Context) (error, int) {
	reqchan := make(chan Response)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*250)
	defer cancel()
	go func() {
		err, code := fetchRequest()
		if err != nil {
			return
		}
		reqchan <- Response{StatusCode: code, err: err}
	}()
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("Request timeout!"), 0
		case res := <-reqchan:
			return res.err, res.StatusCode // 在这里也会调用cancel()
		}
	}
}
func fetchRequest() (error, int) {
	time.Sleep(time.Millisecond * 100)
	return nil, 0
}
