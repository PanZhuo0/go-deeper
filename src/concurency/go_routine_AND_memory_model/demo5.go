package main

import (
	"context"
	"fmt"
	"time"
)

// 埋点Tracer
type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func main() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()
	tr.Shutdown(ctx)
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}
