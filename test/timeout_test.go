package test

import (
	"context"
	"log"
	"testing"
	"time"
)

func routine(ctx context.Context, id int) {
	log.Println(id, "routine begin")
	select {
	case <-time.After(10 * time.Second):
		{
			log.Println(id, "timeout")
		}
	case <-ctx.Done():
		{
			log.Println(id, "ctx done")
		}

	}
}

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	go routine(ctx, 1)
	go routine(ctx, 2)
	time.Sleep(1 * time.Second)
}
