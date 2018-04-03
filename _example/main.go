package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastiankoetsier/schedule"
)

func main() {

	// set up signal handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := &schedule.Schedule{}
	scheduleShutDown := s.Start(ctx)
	// spin up a go routine to check for a signla to shut down the schedule
	go func() {
		<-sigs
		cancel()
		<-scheduleShutDown
	}()

	// You can set up your commands very easily, using a fluent API
	s.Command(schedule.RunFunc(func(ctx context.Context) {
		fmt.Println("I run every minute", time.Now())
	})).EveryMinute()

	// You can even chain these together
	s.Command(schedule.RunFunc(func(ctx context.Context) {
		fmt.Println("I run only mondays, but every fifteen minutes of that weekday", time.Now())
	})).Mondays().EveryFifteenMinutes()

	<-scheduleShutDown
}
