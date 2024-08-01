package chantest

import (
	"context"
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	fatherChan()
}

func fatherChan() {
	ch := make(chan int)
	defer close(ch)
	ctx, cancel := context.WithCancel(context.Background())

	go sonChan(ctx, ch)
	fmt.Println(<-ch)
	cancel()
}

func sonChan(ctx context.Context, ch chan<- int) {
	ch <- 1
	defer func() {
		fmt.Println("sonEnd")
	}()
	<-ctx.Done()
}
