package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexgo92/GO_level_2/homework_4/task_2/internal/smile"
)

func main() {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancelFunc := context.WithCancel(context.Background())

	defer cancelFunc()
	smile.Run(ctx)

	s := <-signalChanel
	fmt.Println(s)
}
