package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func work(ctx context.Context, id int) {
	fmt.Println("begin routine", id)

	if id < 5 {
		go work(ctx, id+1)
		<-ctx.Done()
		fmt.Println(id, "routine recieve done")
	} else if id < 10 {
		ctx, cancenlFunc := context.WithCancel(ctx)
		go work(ctx, id+1)
		fmt.Println(id, "routine recieve done")
		if id == 5 {
			cancenlFunc()
		}
	}
}

func TestCancel(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	go work(ctx, 0)
	time.Sleep(3 * time.Second)
	//cancelFunc()
	time.Sleep(3 * time.Second)
}
