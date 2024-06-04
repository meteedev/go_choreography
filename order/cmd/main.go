package main

import (
	"context"
	"time"

	"github.com/meteedev/go_choreography/queue"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	queue.SetUpQueue(ctx)
}
